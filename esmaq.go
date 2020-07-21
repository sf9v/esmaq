package esmaq

import "context"

type (
	// StateType is a state type
	StateType string
	// Event is an event type
	EventType string
)

// StateConfig is a state config
type StateConfig struct {
	From        StateType
	Actions     Actions
	Transitions []TransitionConfig
}

type Actions struct {
	OnEnter func(context.Context) error
	OnExit  func(context.Context) error
}

// Transitions is a transition config
type TransitionConfig struct {
	To    StateType
	Event EventType
}
