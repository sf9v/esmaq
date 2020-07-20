package esmaq

type (
	// State is a state type
	State string
	// Event is an event type
	Event string
)

// StateConfig is a state config
type StateConfig struct {
	From        State
	Transitions []Transitions
}

// Transitions is a transition config
type Transitions struct {
	To       State
	Event    Event
	Callback Callback
}

// Callback is the callback function config
type Callback struct {
	Ins  Ins
	Outs Outs
}

// Ins are input parameters
type Ins map[string]interface{}

// Outs are output parameters
type Outs map[string]interface{}
