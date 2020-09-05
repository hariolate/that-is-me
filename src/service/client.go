package service

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"gtihub.com/gin-websocket/src/protocol"
	"log"
	"time"
)

type client struct {
	conn *websocket.Conn
	srv  *Service

	id        uint64
	firstPong bool

	toSendChan chan *protocol.Wrapper

	c          context.Context
	cancelFunc context.CancelFunc

	availableForMatch bool
	match             *match
}

func (c *client) sendMessage(t protocol.Wrapper_MessageType, m proto.Message) {
	wrappedMessage := &protocol.Wrapper{
		Type:    t,
		Message: MustMarshalAnyProto(m),
	}
	c.toSendChan <- wrappedMessage
}

func (c *client) setupWorkers() {
	go c.receiveWorker()
	go c.sendWorker()
}

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
			go c.srv.onReceiveNewMessage(c, c.handleReceivedMessage(rm.p))
		case <-c.c.Done():
			return
		}
	}
}

func (c *client) handleReceivedMessage(data []byte) *protocol.Wrapper {
	var wrappedMessage protocol.Wrapper
	NoError(proto.Unmarshal(data, &wrappedMessage))
	return &wrappedMessage
}

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

func (c *client) sendWrappedMessage(m *protocol.Wrapper) {
	c.toSendChan <- m
}
