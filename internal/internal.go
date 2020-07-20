package internal

// State is a state
type State struct {
	Transitions map[string]*Transition
}

// Transition is a transition
type Transition struct {
	To string
}
