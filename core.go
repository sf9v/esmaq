package esmaq

type State struct {
	Transitions map[EventType]*Transition

	Actions *Actions
}

type Transition struct {
	To StateType
}

// Core is a state machine core
type Core struct {
	states map[StateType]*State
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
	st, ok := c.states[s]
	if !ok {
		return nil, newUndefinedStateError(s)
	}

	return st, nil
}

// NewCore is a factory for state machine core
func NewCore(stateConfigs []StateConfig) *Core {
	states := map[StateType]*State{}
	for _, stateConfig := range stateConfigs {
		_, ok := states[stateConfig.From]
		if ok {
			continue
		}

		trsns := map[EventType]*Transition{}
		for _, trsn := range stateConfig.Transitions {
			event := trsn.Event
			_, ok = trsns[event]
			if ok {
				continue
			}

			trsns[event] = &Transition{To: trsn.To}
		}

		states[stateConfig.From] = &State{
			Transitions: trsns,
			Actions:     stateConfig.Actions,
		}
	}

	return &Core{states: states}
}
