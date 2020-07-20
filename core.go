package esmaq

import (
	"github.com/stevenferrer/esmaq/internal"
)

// Core is a state machin core
type Core struct {
	stateMap map[State]*internal.State
}

// CanTransition will return an error if transition is not allowed
func (c *Core) CanTransition(e Event, from State) error {
	// get from state
	fs, err := c.getState(from)
	if err != nil {
		return err
	}

	// verify event is allowed
	tr, ok := fs.Transitions[string(e)]
	if !ok {
		return newUndefinedEventError(e, from)
	}

	// get to state
	_, err = c.getState(State(tr.To))
	if err != nil {
		return err
	}

	return nil
}

func (c *Core) getState(s State) (*internal.State, error) {
	st, ok := c.stateMap[s]
	if !ok {
		return nil, newUndefinedStateError(s)
	}

	return st, nil
}

// NewCore is a factory for state machine core
func NewCore(stCfgs []StateConfig) *Core {
	stMap := map[State]*internal.State{}

	for _, stCfg := range stCfgs {
		_, ok := stMap[stCfg.From]
		if ok {
			continue
		}

		trs := map[string]*internal.Transition{}
		for _, tr := range stCfg.Transitions {
			event := string(tr.Event)
			_, ok = trs[event]
			if ok {
				continue
			}

			trs[event] = &internal.Transition{To: string(tr.To)}
		}

		stMap[stCfg.From] = &internal.State{Transitions: trs}
	}

	return &Core{stateMap: stMap}
}
