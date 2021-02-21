package fsm

type Transition struct {
	From, To string
	Event    string
}

type Callback func(t *Transition, args ...interface{})

func OnEntry(to string, f Callback) Callback {
	return func(t *Transition, args ...interface{}) {
		if t.To == to {
			f(t, args...)
		}
	}
}

func OnExit(from string, f Callback) Callback {
	return func(t *Transition, args ...interface{}) {
		if t.From == from {
			f(t, args...)
		}
	}
}

func OnXXXEvent(event string, f Callback) Callback {
	return func(t *Transition, args ...interface{}) {
		if t.Event == event {
			f(t, args...)
		}
	}
}

func OnTransEvent(from, to string, f Callback) Callback {
	return func(t *Transition, args ...interface{}) {
		if t.From == from && t.To == to {
			f(t, args...)
		}
	}
}

type FSM struct {
	Transitions []Transition
	Callbacks   []Callback
}

func (f FSM) Trigger(currentState string, event string) string {
	var (
		t     Transition
		exist bool
	)
	for _, tran := range f.Transitions {
		if currentState == tran.From && tran.Event == event {
			t, exist = tran, true
			break
		}
	}
	if !exist {
		t = Transition{Event: event}
	}

	for _, callback := range f.Callbacks {
		callback(&t)
	}
	return t.To
}
