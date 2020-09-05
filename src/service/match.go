package service

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"gtihub.com/gin-websocket/src/fsm"
	"gtihub.com/gin-websocket/src/protocol"
	"math/rand"
	"time"
)

type position struct {
	x, y uint32
}
type object struct {
	objectId uint32
	pos      position
}

type match struct {
	a, b *client
	fsm  fsm.FSM

	data map[string]interface{}

	objects         map[uint32]*object
	playerBObjectId uint32
	gameWin         bool
}

func (m *match) acceptMatch(c *client) {
	if c == m.a {
		m.data["a_accepted_match"] = struct{}{}
	}
	if c == m.b {
		m.data["a_accepted_match"] = struct{}{}
	}

	if _, ok := m.data["a_accepted_match"]; ok {
		if _, ok := m.data["b_accepted_match"]; ok {
			m.bothAcceptedMatch()
		}
	}
}

func (m *match) bothAcceptedMatch() {
	NoError(m.fsm.FireEvent("accept_match", nil))
}

func (m *match) notAcceptMatch() {
	m.shutdown()
	if m.a.availableForMatch {
		m.a.srv.availableForMatch <- m.a
	}
	if m.b.availableForMatch {
		m.b.srv.availableForMatch <- m.b
	}
}

func (m *match) shutdown() {
	NoError(m.fsm.FireEvent("end_match", nil))
	m.a.match = nil
	m.b.match = nil
}

func (m *match) sendPendingMatchEvent() {
	pending := &protocol.MatchMakingResponse{
		Type: protocol.MatchMakingResponse_Available,
	}
	m.a.sendMessage(protocol.Wrapper_MatchMakingResponse, pending)
	m.b.sendMessage(protocol.Wrapper_MatchMakingResponse, pending)
}

func (m *match) sendInitEvent() {
	userAInitEvent := &protocol.UserAInitEvent{
		Objects: m.getCurrentObjectsMessage(),
	}
	userBInitEvent := &protocol.UserBInitEvent{
		Objects:        m.getCurrentObjectsMessage(),
		ThatObjectIsMe: m.playerBObjectId,
	}
	m.a.sendMessage(protocol.Wrapper_UserAInitEvent, userAInitEvent)
	m.b.sendMessage(protocol.Wrapper_UserBInitEvent, userBInitEvent)
}

func (m *match) getCurrentObjectsMessage() (res []*protocol.Object) {
	for _, object := range m.objects {
		res = append(res, &protocol.Object{
			Position: &protocol.Position{
				X: object.pos.x,
				Y: object.pos.y,
			},
			ObjectId: object.objectId,
		})
	}
	return
}

const (
	objectNum = 20

	canvasHeight = 1024
	canvasWeight = 1024

	gameDuration = 3 * time.Minute
	chatDuration = time.Minute
)

func (m *match) genObjects() {
	m.objects = make(map[uint32]*object)
	for i := 0; i < objectNum; i++ {
		obj := &object{
			objectId: uint32(rand.Int()),
			pos: position{
				x: uint32(rand.Int()) % canvasWeight,
				y: uint32(rand.Int()) % canvasHeight,
			},
		}
		m.objects[obj.objectId] = obj
		m.playerBObjectId = obj.objectId
	}
}

func newMatch(a, b *client) *match {
	match := &match{
		a:    a,
		b:    b,
		data: make(map[string]interface{}),
	}
	machine := fsm.New(
		"match_pending",
		[]fsm.EventDesc{
			{"accept_match", []string{"match_pending"}, "match_accepted"},
			{"not_accept_match", []string{"match_pending"}, "match_exit"},
			{"clients_ready", []string{"match_accepted"}, "game_ready"},
			{"begin_game", []string{"game_ready"}, "game_begin"},
			{"end_game", []string{"game_begin"}, "game_end"},
			{"game_success", []string{"game_end"}, "chat_pending"},
			{"game_failure", []string{"game_end"}, "match_exit"},
			{"clients_ready", []string{"chat_pending"}, "chat_ready"},
			{"begin_chat", []string{"chat_ready"}, "chat_begin"},
			{"end_chat", []string{"chat_begin"}, "chat_end"},
			{"end_match", []string{"chat_end"}, "match_exit"},
		},
		map[string]fsm.Callback{
			"enter_match_accepted": func(e *fsm.Event) {
				a.availableForMatch = false
				b.availableForMatch = false
				match.sendMatchStateChangeEventToAB(&protocol.MatchStateChangeEvent{Type: protocol.MatchStateChangeEvent_MatchAccepted})
				match.sendInitEvent()
			},
			"enter_game_ready": func(e *fsm.Event) {
				match.sendMatchStateChangeEventToAB(&protocol.MatchStateChangeEvent{Type: protocol.MatchStateChangeEvent_GameReady})
			},
			"enter_game_begin": func(e *fsm.Event) {
				match.gameWin = false
				go match.gameWorker()
				match.sendMatchStateChangeEventToAB(&protocol.MatchStateChangeEvent{Type: protocol.MatchStateChangeEvent_GameBegin})
			},
			"enter_game_end": func(e *fsm.Event) {
				match.sendMatchStateChangeEventToAB(&protocol.MatchStateChangeEvent{Type: protocol.MatchStateChangeEvent_GameEnd})
				match.sendGameResultEvent()
			},
			"enter_chat_pending": func(e *fsm.Event) {
				match.sendMatchStateChangeEventToAB(&protocol.MatchStateChangeEvent{Type: protocol.MatchStateChangeEvent_ChatPending})
			},
			"enter_chat_ready": func(e *fsm.Event) {
				match.sendMatchStateChangeEventToAB(&protocol.MatchStateChangeEvent{Type: protocol.MatchStateChangeEvent_ChatReady})
			},
			"enter_chat_begin": func(e *fsm.Event) {
				go match.chatTimerWorker()
				match.data["chat_end_time"] = time.Now().Add(chatDuration)
				match.sendMatchStateChangeEventToAB(&protocol.MatchStateChangeEvent{Type: protocol.MatchStateChangeEvent_ChatBegin})
			},
			"enter_chat_end": func(e *fsm.Event) {
				match.sendMatchStateChangeEventToAB(&protocol.MatchStateChangeEvent{Type: protocol.MatchStateChangeEvent_ChatEnd})
			},
			"enter_match_exit": func(e *fsm.Event) {
				match.sendMatchStateChangeEventToAB(&protocol.MatchStateChangeEvent{Type: protocol.MatchStateChangeEvent_MatchExit})
			},
		},
	)
	match.fsm = machine
	match.genObjects()
	match.sendMatchStateChangeEventToAB(&protocol.MatchStateChangeEvent{Type: protocol.MatchStateChangeEvent_MatchPending})
	match.sendPendingMatchEvent()
	return match
}

func (m *match) sendGameResultEvent() {
	msg := protocol.GameResultEvent{}

	if m.gameWin {
		msg.Type = protocol.GameResultEvent_Success
	} else {
		msg.Type = protocol.GameResultEvent_Failure
	}
	m.a.sendMessage(protocol.Wrapper_GameResultEvent, &msg)
	m.b.sendMessage(protocol.Wrapper_GameResultEvent, &msg)
}

func (m *match) sendMatchStateChangeEventToAB(msg proto.Message) {
	m.a.sendMessage(protocol.Wrapper_MatchStateChangeEvent, msg)
	m.b.sendMessage(protocol.Wrapper_MatchStateChangeEvent, msg)
}

func (m *match) handleMatchEvents(a *client, t protocol.Wrapper_MessageType, any *any.Any) {
	switch t {
	case protocol.Wrapper_ClientReadyEvent:
		m.handleClientReadyEvent(a)
	case protocol.Wrapper_UserAKillEvent:
		var e protocol.UserAKillEvent
		MustUnmarshalAnyProto(any, &e)
		m.handleUserAKillEvent(&e)

	}
}

func (m *match) handleClientReadyEvent(c *client) {
	if c == m.a {
		m.data["a_ready"] = struct{}{}
	}
	if c == m.b {
		m.data["a_ready"] = struct{}{}
	}

	if _, ok := m.data["a_ready"]; ok {
		if _, ok := m.data["b_ready"]; ok {
			m.bothReady()
			delete(m.data, "a_ready")
			delete(m.data, "b_ready")
		}
	}
}

func (m *match) bothReady() {
	NoError(m.fsm.FireEvent("clients_ready", nil))
	if m.fsm.Is("game_ready") {
		NoError(m.fsm.FireEvent("begin_game", nil))
	}
	if m.fsm.Is("chat_pending") {
		NoError(m.fsm.FireEvent("begin_chat", nil))
	}
}

func (m *match) sendObjectsLocationUpdateEvent() {
	loc := m.getCurrentObjectsMessage()
	msg := protocol.ObjectsLocationUpdateEvent{Objects: loc}
	m.a.sendMessage(protocol.Wrapper_ObjectsLocationUpdateEvent, &msg)
	m.b.sendMessage(protocol.Wrapper_ObjectsLocationUpdateEvent, &msg)
}

func (m *match) gameWorker() {
	t := time.NewTicker(gameDuration)
	defer func() {
		t.Stop()
		NoError(m.fsm.FireEvent("end_game", nil))
		if m.gameWin {
			NoError(m.fsm.FireEvent("game_success", nil))
		} else {
			NoError(m.fsm.FireEvent("game_failure", nil))
		}
	}()
	for {
		select {
		case <-t.C:
			return
		default:
			if _, ok := m.objects[m.playerBObjectId]; !ok {
				return
			}
			if len(m.objects) == 1 {
				m.gameWin = true
				return
			}
			m.sendObjectsLocationUpdateEvent()
		}
	}
}

func (m *match) handleUserAKillEvent(e *protocol.UserAKillEvent) {
	if m.fsm.Is("game_begin") {
		delete(m.objects, e.ObjectId)
		return
	}
	panic("invalid event")
}

func (m *match) handleUserBMoveEvent(e *protocol.UserBMoveEvent) {
	if m.fsm.Is("game_begin") {
		if object, ok := m.objects[m.playerBObjectId]; ok {
			object.pos = position{
				x: e.To.X,
				y: e.To.Y,
			}
		}
		return
	}
	panic("invalid event")
}

func (m *match) handleRawChatMessageEvent(c *client, r *protocol.RawChatMessageEvent) {
	if m.fsm.Is("chat_begin") {
		var msg protocol.NewMessageEvent

		timestamp, err := ptypes.TimestampProto(time.Now())
		NoError(err)
		msg.Timestamp = timestamp
		msg.TimeRemaining = int64(time.Now().Sub(m.data["chat_end_time"].(time.Time)) / time.Second)
		msg.Raw = r
		if c == m.a {
			msg.Sender = protocol.NewMessageEvent_A
		} else {
			msg.Sender = protocol.NewMessageEvent_B
		}

		m.a.sendMessage(protocol.Wrapper_NewMessageEvent, &msg)
		m.b.sendMessage(protocol.Wrapper_NewMessageEvent, &msg)
	}
	panic("invalid event")
}

func (m *match) chatTimerWorker() {
	ticker := time.NewTicker(chatDuration + time.Second*5)
	defer ticker.Stop()
	select {
	case <-ticker.C:
		NoError(m.fsm.FireEvent("end_chat", nil))
		NoError(m.fsm.FireEvent("end_match", nil))
	}
}
