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

// UndefinedEventError is an error when an event is not defined
type UndefinedEventError struct {
	event EventType
	from  StateType
}

func (e *UndefinedEventError) Error() string {
	return fmt.Sprintf("transition event %q is not allowed in %q state", e.event, e.from)
}

func newUndefinedEventError(event EventType, from StateType) *UndefinedEventError {
	return &UndefinedEventError{event: event, from: from}
}
