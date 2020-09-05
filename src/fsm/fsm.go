// ported from github.com/chfanghr/fsm

package fsm

import (
	"strings"
	"sync"
)

type callbackType int

const (
	callbackNone callbackType = iota
	callbackBeforeEvent
	callbackLeaveState
	callbackEnterState
	callbackAfterEvent
)

type FSM interface {
	Current() string
	Is(state string) bool
	SetState(state string)
	Can(event string) bool
	Cannot(event string) bool
	AvailableTransitions() []string
	Transition() error
	FireEvent(event string, args []interface{}) error
	SetCallback(name string, callback Callback)
}

func New(initial string, events []EventDesc, callbacks map[string]Callback) FSM {
	machine := &fsm{
		current:      initial,
		transitioner: new(transitionerImpl),
		transitions:  make(map[eKey]string),
		callbacks:    make(map[cKey]Callback),
	}

	allEvents, allStates := make(map[string]struct{}), make(map[string]struct{})

	for _, event := range events {
		for _, src := range event.Src {
			machine.transitions[eKey{event.Name, src}] = event.Dst
			allStates[src] = struct{}{}
			allStates[event.Dst] = struct{}{}
		}
		allEvents[event.Name] = struct{}{}
	}

	for k, v := range callbacks {
		machine.SetCallback(k, v)
	}

	return machine
}

type transitioner interface {
	transition(fsm FSM) error
}

type transitionerImpl struct {
}

type cKey struct {
	target       string
	callbackType callbackType
}

type eKey struct {
	event string
	src   string
}

type Callback func(*Event)

type Callbacks map[string]Callback

type EventDesc struct {
	Name string
	Src  []string
	Dst  string
}

type fsm struct {
	current      string
	transitions  map[eKey]string
	callbacks    map[cKey]Callback
	transition   func()
	transitioner transitioner
	eventMu      sync.Mutex
	stateMu      sync.RWMutex
}

func (f *fsm) Current() string {
	f.stateMu.RLock()
	defer f.stateMu.RUnlock()
	return f.current
}

func (f *fsm) Is(state string) bool {
	f.stateMu.RLock()
	defer f.stateMu.RUnlock()
	return state == f.current
}

func (f *fsm) SetState(state string) {
	f.stateMu.Lock()
	defer f.stateMu.Unlock()
	f.current = state
}

func (f *fsm) Can(event string) bool {
	f.stateMu.RLock()
	defer f.stateMu.RUnlock()
	_, ok := f.transitions[eKey{event, f.current}]
	return ok && f.transition == nil
}

func (f *fsm) Cannot(event string) bool {
	return !f.Can(event)
}

func (f *fsm) AvailableTransitions() []string {
	f.stateMu.RLock()
	defer f.stateMu.RUnlock()
	var res []string
	for key := range f.transitions {
		if key.src == f.current {
			res = append(res, key.event)
		}
	}
	return res
}

func (f *fsm) Transition() error {
	f.stateMu.Lock()
	defer f.stateMu.Unlock()
	return f.doTransition()
}

func (f *fsm) FireEvent(event string, args []interface{}) error {
	f.eventMu.Lock()
	f.stateMu.RLock()

	defer f.eventMu.Unlock()

	shouldRUnlockStateMu := true

	defer func() {
		if shouldRUnlockStateMu {
			f.stateMu.RUnlock()
		}
	}()

	if f.transition != nil {
		return ErrorInTransition{}
	}

	dst, ok := f.transitions[eKey{event, f.current}]
	if !ok {
		for key := range f.transitions {
			if key.event == event {
				return ErrorInvalidEvent{event, f.current}
			}
			return ErrorUnknownEvent{event}
		}
	}

	e := &Event{
		Machine:  f,
		Event:    event,
		Src:      f.current,
		Dst:      dst,
		Err:      nil,
		Args:     args,
		canceled: false,
		async:    false,
	}

	if err := f.beforeEventCallbacks(e); err != nil {
		return err
	}

	if f.current == dst {
		f.afterEventCallbacks(e)
		return ErrorNoTransition{e.Err}
	}

	f.transition = func() {
		f.stateMu.Lock()
		f.current = dst
		f.stateMu.Unlock()

		f.enterStateCallbacks(e)
		f.afterEventCallbacks(e)
	}

	if err := f.leaveStateCallbacks(e); err != nil {
		if _, ok := err.(ErrorCanceled); ok {
			f.transition = nil
		}
		return err
	}

	shouldRUnlockStateMu = false
	f.stateMu.RUnlock()

	if err := f.doTransition(); err != nil {
		return ErrorInternal{}
	}

	return e.Err
}

func (f *fsm) SetCallback(name string, callback Callback) {
	f.stateMu.Lock()
	f.eventMu.Lock()

	defer f.stateMu.Unlock()
	defer f.eventMu.Unlock()

	type t int

	const (
		tNone t = iota
		tEvent
		tState
	)

	find := func(target string) t {
		for k, v := range f.transitions {
			if k.src == target || v == target {
				return tState
			}
			if k.event == target {
				return tEvent
			}
		}
		return tNone
	}

	isEvent := func(target string) bool {
		findRes := find(target)
		return findRes == tEvent
	}

	isState := func(target string) bool {
		findRes := find(target)
		return findRes == tState
	}

	var target string
	var cbType = callbackNone

	targetPtr := &target

	helper := func(prefix string) bool {
		target = *targetPtr
		if strings.HasPrefix(name, prefix) {
			target = name[len(prefix):]
			return true
		}
		return false
	}

	if helper("before_") {
		if target == "event" {
			target = ""
			cbType = callbackBeforeEvent
		} else if isEvent(target) {
			cbType = callbackBeforeEvent
		}
	} else if helper("leave_") {
		if target == "state" {
			target = ""
			cbType = callbackLeaveState
		} else if isState(target) {
			cbType = callbackLeaveState
		}
	} else if helper("enter_") {
		if target == "state" {
			target = ""
			cbType = callbackEnterState
		} else if isState(target) {
			cbType = callbackEnterState
		}
	} else if helper("after_") {
		if target == "event" {
			target = ""
			cbType = callbackAfterEvent
		} else if isEvent(target) {
			cbType = callbackAfterEvent
		}
	} else {
		target = name
		if isState(target) {
			cbType = callbackEnterState
		} else if isEvent(target) {
			cbType = callbackAfterEvent
		}
	}

	if cbType != callbackNone {
		f.callbacks[cKey{target, cbType}] = callback
	}
}

func (f *fsm) doTransition() error {
	return f.transitioner.transition(f)
}

func (f *fsm) afterEventCallbacks(e *Event) {
	callback, ok := f.callbacks[cKey{e.Event, callbackAfterEvent}]
	if ok {
		callback(e)
	}
	callback, ok = f.callbacks[cKey{"", callbackAfterEvent}]
	if ok {
		callback(e)
	}
}

func (f *fsm) enterStateCallbacks(e *Event) {
	callback, ok := f.callbacks[cKey{e.Event, callbackEnterState}]
	if ok {
		callback(e)
	}
	callback, ok = f.callbacks[cKey{"", callbackEnterState}]
	if ok {
		callback(e)
	}
}

func (f *fsm) leaveStateCallbacks(e *Event) error {
	callback, ok := f.callbacks[cKey{e.Event, callbackLeaveState}]
	if ok {
		callback(e)
		if e.canceled {
			return ErrorCanceled{e.Err}
		} else if e.async {
			return ErrorAsync{e.Err}
		}
	}

	callback, ok = f.callbacks[cKey{"", callbackLeaveState}]
	if ok {
		callback(e)
		if e.canceled {
			return ErrorCanceled{e.Err}
		} else if e.async {
			return ErrorAsync{e.Err}
		}
	}

	return nil
}

func (f *fsm) beforeEventCallbacks(e *Event) error {
	callback, ok := f.callbacks[cKey{e.Event, callbackBeforeEvent}]
	if ok {
		callback(e)
		if e.canceled {
			return ErrorCanceled{e.Err}
		}
	}

	callback, ok = f.callbacks[cKey{"", callbackLeaveState}]
	if ok {
		callback(e)
		if e.canceled {
			return ErrorCanceled{e.Err}
		}
	}

	return nil
}

func (t *transitionerImpl) transition(machine FSM) error {
	fsm := machine.(*fsm)
	if fsm.transition == nil {
		return ErrorNotInTransition{}
	}
	fsm.transition()
	fsm.stateMu.Lock()
	defer fsm.stateMu.Unlock()
	fsm.transition = nil
	return nil
}
