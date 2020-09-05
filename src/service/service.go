package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"gtihub.com/hariolate/that-is-me/src/protocol"
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

func FromConfig(conf *Config, ctx context.Context) *Service {
	srv := &Service{
		id: uuid.New(),

		c: ctx,

		Engine: gin.Default(),
		uid:    0,

		clients:           make(map[uint64]*client),
		availableForMatch: make(chan *client, 10000),
	}
	go srv.matchMakingWorker()
	return srv
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

func (s *Service) handleMatchMakingRequest(a *client, any []byte) {
	var matchMakingRequest protocol.MatchMakingRequest
	MustUnmarshalProto(any, &matchMakingRequest)
	requestType := matchMakingRequest.Type
	log.Printf("client %d: match making request: %s", a.id, matchMakingRequest.Type)
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
			log.Printf("client %d: avalable for match making", cli.id)
			if a == nil {
				a = cli
				continue
			}
			if b == nil {
				b = cli
			}

			s.newMatch(a, b)
			a = nil
			b = nil
		}
	}
}

func (s *Service) newMatch(a, b *client) {
	log.Printf("creating new match for client %d and %d", a.id, b.id)
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
