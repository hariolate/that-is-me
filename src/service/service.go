package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"gtihub.com/gin-websocket/src/protocol"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

type Service struct {
	id uuid.UUID

	r *redis.Client
	c context.Context

	*gin.Engine

	uid uint32

	clients map[uint32]*client
}

func FromConfig(c *Config, ctx context.Context) *Service {
	return &Service{
		id: uuid.New(),

		r: redis.NewClient(MustParseRedisURL(c.Storage.RedisURL)),
		c: ctx,

		Engine: gin.Default(),
		uid:    0,

		clients: make(map[uint32]*client),
	}
}

func (s *Service) redisChatHistoryKey() string {
	return fmt.Sprintf("chat:%s:history", s.id)
}

func (s *Service) storeChatMessage(m *protocol.Message) {
	data, err := proto.Marshal(m)
	NoError(err)
	NoError(s.r.LPush(s.c, s.redisChatHistoryKey(), string(data)).Err())
}

func (s *Service) getNextUid() uint32 {
	return atomic.AddUint32(&s.uid, 1)
}

func (s *Service) getAllHistoryMessages() []*protocol.Message {
	protoMessages, err := s.r.LRange(s.c, s.redisChatHistoryKey(), 0, -1).Result()
	NoError(err)

	var results = make([]*protocol.Message, len(protoMessages))

	for i, protoMessage := range protoMessages {
		var message protocol.Message
		NoError(proto.Unmarshal([]byte(protoMessage), &message))
		//NoError(json.Unmarshal([]byte(jsonMessage), &message))
		results[i] = &message
	}

	return results
}

func (s *Service) onNewMessage(m *protocol.Message) {
	if m == nil {
		return
	}
	s.broadcastMessage(m)
	s.storeChatMessage(m)
}

func (s *Service) broadcastMessage(m *protocol.Message) {
	clients := s.clients

	for _, client := range clients {
		go client.sendMessage(m)
	}
}

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

func (s *Service) newClient(w http.ResponseWriter, r *http.Request) {
	wsupgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	wsupgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	conn, err := wsupgrader.Upgrade(w, r, nil)
	NoError(err)

	c, f := context.WithCancel(s.c)

	cli := &client{
		conn:       conn,
		srv:        s,
		id:         s.getNextUid(),
		firstPong:  true,
		toSendChan: make(chan *protocol.Message, 100),
		c:          c,
		cancelFunc: f,
	}

	cli.setupWorkers()

	s.clients[cli.id] = cli
}

func (s *Service) removeClient(c *client) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("removing client %d: %s\n", c.id, r)
		}
	}()
	_, ok := s.clients[c.id]
	if !ok {
		return
	}
	delete(s.clients, c.id)

	c.cancelFunc()
	close(c.toSendChan)
	_ = c.conn.Close()
}

func (s *Service) Handler(c *gin.Context) {
	s.newClient(c.Writer, c.Request)
	c.Status(http.StatusOK)
}

func (s *Service) Shutdown() {
	for _, c := range s.clients {
		s.removeClient(c)
	}
}

// TODO close all connections when service is closed
