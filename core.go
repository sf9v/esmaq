package esmaq

// State is a state
type State struct {
	State       StateType
	Transitions map[EventType]*Transition
	Actions     Actions
}

// Transition is a transition
type Transition struct {
	To StateType
}

// Core is a state machine core
type Core struct {
	stateMap map[StateType]*State
}

// Transition returns an error if transition is not allowed
func (c *Core) Transition(from StateType, event EventType) (*State, error) {
	// get "from" state
	fromState, err := c.GetState(from)
	if err != nil {
		return nil, err
	}

	// verify event is allowed
	toTrsn, ok := fromState.Transitions[event]
	if !ok {
		return nil, newUndefinedEventError(event, from)
	}

	// get "to" state
	toState, err := c.GetState(toTrsn.To)
	if err != nil {
		return nil, err
	}

	return toState, nil
}

// GetState returns the state
func (c *Core) GetState(s StateType) (*State, error) {
	st, ok := c.stateMap[s]
	if !ok {
		return nil, newUndefinedStateError(s)
	}

	return st, nil
}

// NewCore is a factory for state machine core
func NewCore(stateConfigs []StateConfig) *Core {
	stateMap := map[StateType]*State{}

	// all possible states
	states := map[StateType]bool{}

	for _, sc := range stateConfigs {
		if _, ok := stateMap[sc.From]; ok {
			continue
		}

		trs := map[EventType]*Transition{}
		for _, tr := range sc.Transitions {
			event := tr.Event
			if _, ok := trs[event]; ok {
				continue
			}

			trs[event] = &Transition{To: tr.To}

			if _, ok := states[tr.To]; !ok {
				states[tr.To] = true
			}
		}

		stateMap[sc.From] = &State{
			State:       sc.From,
			Transitions: trs,
			Actions:     sc.Actions,
		}
	}

	// make sure all states are defined
	for state := range states {
		if _, ok := stateMap[state]; !ok {
			stateMap[state] = &State{State: state}
		}
	}

	return &Core{stateMap: stateMap}
}
