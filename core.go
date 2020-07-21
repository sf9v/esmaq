package esmaq

// State is a state
type State struct {
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

// Transition will return an error if transition is not allowed
func (c *Core) Transition(event EventType, from StateType) (*State, error) {
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

	for _, stateConfig := range stateConfigs {
		if _, ok := stateMap[stateConfig.From]; ok {
			continue
		}

		trsns := map[EventType]*Transition{}
		for _, trsn := range stateConfig.Transitions {
			event := trsn.Event
			if _, ok := trsns[event]; ok {
				continue
			}

			trsns[event] = &Transition{To: trsn.To}

			if _, ok := states[trsn.To]; !ok {
				states[trsn.To] = true
			}
		}

		stateMap[stateConfig.From] = &State{
			Transitions: trsns,
			Actions:     stateConfig.Actions,
		}
	}

	// make sure all states are defined
	for state := range states {
		if _, ok := stateMap[state]; !ok {
			stateMap[state] = &State{}
		}
	}

	return &Core{stateMap: stateMap}
}
