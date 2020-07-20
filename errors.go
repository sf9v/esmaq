package esmaq

import "fmt"

// UndefinedStateError is an error when a state is not defined
type UndefinedStateError struct {
	s StateType
}

func (e *UndefinedStateError) Error() string {
	return fmt.Sprintf("state %q is not defined", e.s)
}

func newUndefinedStateError(s StateType) *UndefinedStateError {
	return &UndefinedStateError{s: s}
}

// UndefinedEventError is an error when an event is not defined
type UndefinedEventError struct {
	e EventType
	s StateType
}

func (e *UndefinedEventError) Error() string {
	return fmt.Sprintf("event %q not defined for %q state", e.e, e.s)
}

func newUndefinedEventError(e EventType, s StateType) *UndefinedEventError {
	return &UndefinedEventError{e: e, s: s}
}
