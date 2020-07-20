package esmaq

import "fmt"

// UndefinedStateError is an error when a state is not defined
type UndefinedStateError struct {
	s State
}

func (e *UndefinedStateError) Error() string {
	return fmt.Sprintf("state %q is not defined", e.s)
}

func newUndefinedStateError(s State) *UndefinedStateError {
	return &UndefinedStateError{s: s}
}

// UndefinedEventError is an error when an event is not defined
type UndefinedEventError struct {
	e Event
	s State
}

func (e *UndefinedEventError) Error() string {
	return fmt.Sprintf("event %q not defined for %q state", e.e, e.s)
}

func newUndefinedEventError(e Event, s State) *UndefinedEventError {
	return &UndefinedEventError{e: e, s: s}
}
