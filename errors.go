package esmaq

import "fmt"

// UndefinedStateError is an error when a state is not defined
type UndefinedStateError struct {
	state StateType
}

func (e *UndefinedStateError) Error() string {
	return fmt.Sprintf("state %q is not defined", e.state)
}

func newUndefinedStateError(state StateType) *UndefinedStateError {
	return &UndefinedStateError{state: state}
}

// TransitionNotAllowedError is an error when a transition is not allowed
type TransitionNotAllowedError struct {
	event EventType
	from  StateType
}

func (e *TransitionNotAllowedError) Error() string {
	return fmt.Sprintf("transition %q is not allowed in %q state",
		e.event, e.from)
}

func newUndefinedEventError(event EventType, from StateType) *TransitionNotAllowedError {
	return &TransitionNotAllowedError{event: event, from: from}
}
