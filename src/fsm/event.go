package fsm

type Event struct {
	Machine FSM
	Event   string
	Src     string
	Dst     string
	Err     error
	Args    []interface{}

	canceled, async bool
}

func NewEvent(machine FSM) *Event {
	return &Event{
		Machine:  machine,
		canceled: false,
		async:    false,
	}
}

func (e *Event) Cancel(errors []error) {
	if e.canceled {
		return
	}
	e.canceled = true
	if len(errors) > 0 {
		e.Err = errors[0]
	}
}

func (e *Event) Async() {
	e.async = true
}
