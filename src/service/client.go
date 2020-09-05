package service

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/gorilla/websocket"
	"gtihub.com/gin-websocket/src/protocol"
	"log"
	"time"
)

type client struct {
	conn *websocket.Conn
	srv  *Service

	id        uint32
	firstPong bool

	toSendChan chan *protocol.Message

	c          context.Context
	cancelFunc context.CancelFunc
}

func (c *client) redisTimeoutKey() string {
	return fmt.Sprintf("client:%d:timeout", c.id)
}

//const clientTimeout = time.Minute * 2

func (c *client) setupWorkers() {
	//go c.timeoutWorker()
	//go c.pingWorker()
	go c.receiveWorker()
	go c.sendWorker()
}

//func (c *client) timeoutWorker() {
//	NoError(c.srv.r.Set(c.srv.c, c.redisTimeoutKey(), 1, clientTimeout).Err())
//
//	for {
//		if c.srv.r.Get(c.srv.c, c.redisTimeoutKey()).Err() != nil {
//			_ = c.conn.Close()
//			break
//		}
//	}
//}

func (c *client) receiveWorker() {
	type readMessageStruct struct {
		messageType int
		p           []byte
		err         error
	}
	readMessageChan := make(chan readMessageStruct)

	defer func() {
		if r := recover(); r != nil {
			log.Printf("client %d crashed: %s\n", c.id, r)
		}
		close(readMessageChan)
	}()

	c.conn.SetReadLimit(maxMessageSize)
	NoError(c.conn.SetReadDeadline(time.Now().Add(pongWait)))

	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("client %d: receiveWorker crashed: %s\n", c.id, r)
			}
		}()
		t, msg, err := c.conn.ReadMessage()
		readMessageChan <- readMessageStruct{t, msg, err}
	}()

	for {
		select {
		case rm := <-readMessageChan:
			NoError(rm.err)
			go c.srv.onNewMessage(c.handleNewMessage(rm.messageType, rm.p))
		case <-c.c.Done():
			return
		}
	}
}

func (c *client) handleNewMessage(messageType int, data []byte) *protocol.Message {
	//if messageType == websocket.PingMessage {
	//	c.handlePing()
	//	return nil
	//}

	//var raw RawMessage
	//NoError(json.Unmarshal(data, &raw))

	var raw protocol.RawMessage
	NoError(proto.Unmarshal(data, &raw))
	//if messageType != websocket.TextMessage {
	//	raw.Message = "--unsupported message--"
	//}

	//return &Message{
	//	UID:       c.id,
	//	Raw:       raw,
	//	Timestamp: time.Now(),
	//}

	now, err := ptypes.TimestampProto(time.Now())
	NoError(err)

	return &protocol.Message{
		Uid:       c.id,
		Raw:       &raw,
		Timestamp: now,
	}
}

//func (c *client) handlePing() {
//	writer, err := c.conn.NextWriter(websocket.PongMessage)
//	NoError(err)
//	_, err = writer.Write([]byte("pong"))
//	NoError(err)
//
//	NoError(c.srv.r.Set(c.srv.c, c.redisTimeoutKey(), 1, clientTimeout).Err())
//
//	if c.firstPong {
//		go c.sendMessages(c.srv.getAllHistoryMessages())
//		c.firstPong = false
//	}
//}

//func (c *client) pingWorker() {
//	ticker := time.NewTicker(pingPeriod)
//	defer func() {
//		ticker.Stop()
//		c.srv.removeClient(c)
//		_ = c.conn.Close()
//	}()
//
//
//
//	for {
//		<-ticker.C
//		NoError(c.conn.SetWriteDeadline(time.Now().Add(writeWait)))
//		NoError(c.conn.WriteMessage(websocket.PingMessage, nil))
//	}
//}

func (c *client) sendWorker() {
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		if r := recover(); r != nil {
			log.Printf("client %d crashed: %s\n", c.id, r)
			c.srv.removeClient(c)
		}
	}()

	c.conn.SetPongHandler(func(string) error {
		if c.firstPong {
			go c.sendAllHistoryMessages()
			c.firstPong = false
		}
		NoError(c.conn.SetReadDeadline(time.Now().Add(pongWait)))
		return nil
	})

	NoError(c.conn.SetWriteDeadline(time.Now().Add(writeWait)))
	NoError(c.conn.WriteMessage(websocket.PingMessage, nil))

	for {
		select {
		case <-ticker.C:
			NoError(c.conn.SetWriteDeadline(time.Now().Add(writeWait)))
			NoError(c.conn.WriteMessage(websocket.PingMessage, nil))
		case toSend := <-c.toSendChan:
			data, err := proto.Marshal(toSend)
			NoError(err)
			NoError(c.conn.SetWriteDeadline(time.Now().Add(writeWait)))
			NoError(c.conn.WriteMessage(websocket.BinaryMessage, data))
		case <-c.c.Done():
			return
		}
	}
}

func (c *client) sendMessage(m *protocol.Message) {
	//defer func() {
	//	if r := recover(); r != nil {
	//		log.Printf("client %d crashed: %s\n", c.id, r)
	//		c.srv.removeClient(c)
	//		_ = c.conn.Close()
	//	}
	//}()
	//data, err := proto.Marshal(m)
	//NoError(err)
	//NoError(c.conn.SetWriteDeadline(time.Now().Add(writeWait)))
	//NoError(c.conn.WriteMessage(websocket.BinaryMessage, data))
	c.toSendChan <- m
}

func (c *client) sendMessages(ms []*protocol.Message) {
	for _, m := range ms {
		go c.sendMessage(m)
	}
}

func (c *client) sendAllHistoryMessages() {
	c.sendMessages(c.srv.getAllHistoryMessages())
}
