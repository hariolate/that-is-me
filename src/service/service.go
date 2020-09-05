package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"gtihub.com/gin-websocket/src/protocol"
	"log"
	"net/http"
	"sync/atomic"
)

type Service struct {
	id uuid.UUID
	c  context.Context
	*gin.Engine
	uid               uint64
	clients           map[uint64]*client
	availableForMatch chan *client
}

func FromConfig(ctx context.Context) *Service {
	return &Service{
		id: uuid.New(),

		c: ctx,

		Engine: gin.Default(),
		uid:    0,

		clients:           make(map[uint64]*client),
		availableForMatch: make(chan *client, 10000),
	}
}

func (s *Service) getNextUid() uint64 {
	return atomic.AddUint64(&s.uid, 1)
}

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
		conn:              conn,
		srv:               s,
		id:                s.getNextUid(),
		firstPong:         true,
		toSendChan:        make(chan *protocol.Wrapper, 100),
		c:                 c,
		cancelFunc:        f,
		availableForMatch: false,
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

	c.availableForMatch = false
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

func (s *Service) onReceiveNewMessage(c *client, m *protocol.Wrapper) {
	anyData := m.Message
	messageType := m.Type

	switch messageType {
	case protocol.Wrapper_MatchMakingRequest:
		s.handleMatchMakingRequest(c, anyData)
	default:
		c.match.handleMatchEvents(c, messageType, anyData)
	}
}

func (s *Service) handleMatchMakingRequest(a *client, any *any.Any) {
	var matchMakingRequest protocol.MatchMakingRequest
	MustUnmarshalAnyProto(any, &matchMakingRequest)
	requestType := matchMakingRequest.Type
	switch requestType {
	case protocol.MatchMakingRequest_Begin:
		s.availableForMatch <- a
		a.availableForMatch = true
	case protocol.MatchMakingRequest_Cancel:
		a.availableForMatch = false
	case protocol.MatchMakingRequest_Accept:
		a.match.acceptMatch(a)
	case protocol.MatchMakingRequest_NotAccept:
		a.match.notAcceptMatch()
	}
}

func (s *Service) matchMakingWorker() {
	var a, b *client
	for {
		select {
		case <-s.c.Done():
			return
		case cli := <-s.availableForMatch:
			if a == nil {
				a = cli
				continue
			}
			if b == nil {
				b = cli
				continue
			}

			s.newMatch(a, b)
		}
	}
}

func (s *Service) newMatch(a, b *client) {
	if a.availableForMatch && b.availableForMatch {
		newMatch(a, b)
	} else {
		if a.availableForMatch {
			s.availableForMatch <- a
		}
		if b.availableForMatch {
			s.availableForMatch <- b
		}
	}
}
