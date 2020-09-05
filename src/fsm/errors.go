package fsm

import "fmt"

type ErrorInTransition struct {
	Event string
}

func (i ErrorInTransition) Error() string {
	return fmt.Sprintf("event %s inappropriate because previous transition did not complete", i.Event)
}

type ErrorInvalidEvent struct {
	Event, State string
}

func (i ErrorInvalidEvent) Error() string {
	return fmt.Sprintf("event %s inappropriate in current state %s", i.Event, i.State)
}

type ErrorUnknownEvent struct {
	Event string
}

func (u ErrorUnknownEvent) Error() string {
	return fmt.Sprintf("event %s does not exist", u.Event)
}

type ErrorNotInTransition struct{}

func (n ErrorNotInTransition) Error() string {
	return "transition inappropriate because no state change in progress"
}

type ErrorNoTransition struct {
	Err error
}

func (n ErrorNoTransition) Error() string {
	if n.Err != nil {
		return fmt.Sprintf("no transition with error: %s", n.Err)
	}
	return "no transition"
}

type ErrorCanceled struct {
	Err error
}

func (c ErrorCanceled) Error() string {
	if c.Err != nil {
		return fmt.Sprintf("transition canceled with error: %s", c.Err)
	}
	return "transition canceled"
}

type ErrorAsync struct {
	Err error
}

func (a ErrorAsync) Error() string {
	if a.Err != nil {
		return fmt.Sprintf("async started with error: %s", a.Err)
	}
	return "async started"
}

type ErrorInternal struct{}

func (i ErrorInternal) Error() string {
	return "internal error on state transition"
}

type ErrorUnknownVirtualizeType struct{}

func (u ErrorUnknownVirtualizeType) Error() string {
	return "unknown visualize type"
}
