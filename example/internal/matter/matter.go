// Code generated by esmaq, DO NOT EDIT.
package matter

import (
	"context"
	"errors"
	esmaq "github.com/stevenferrer/esmaq"
)

type State esmaq.StateType

const (
	StateSolid  State = "solid"
	StateLiquid State = "liquid"
	StateGas    State = "gas"
)

type Event esmaq.EventType

const (
	EventMelt     Event = "melt"
	EventFreeze   Event = "freeze"
	EventVaporize Event = "vaporize"
	EventCondense Event = "condense"
)

type ctxKey int

const (
	fromKey ctxKey = iota
	toKey
)

type Matter struct {
	core      *esmaq.Core
	callbacks *Callbacks
}

type Callbacks struct {
	Melt     func(ctx context.Context) (err error)
	Freeze   func(ctx context.Context) (err error)
	Vaporize func(ctx context.Context) (err error)
	Condense func(ctx context.Context) (err error)
}

type Actions struct {
	Solid  *esmaq.Actions
	Liquid *esmaq.Actions
	Gas    *esmaq.Actions
}

func (sm *Matter) Melt(ctx context.Context) (err error) {
	from, ok := fromCtx(ctx)
	if !ok {
		return errors.New("\"from\" state not set in context")
	}

	fromState, err := sm.core.GetState(esmaq.StateType(from))
	if err != nil {
		return err
	}

	toState, err := sm.core.Transition(esmaq.EventType(EventMelt), esmaq.StateType(from))
	if err != nil {
		return err
	}

	// inject "to" in context
	ctx = ctxWtTo(ctx, StateLiquid)

	fromState.Actions.OnExit()

	err = sm.callbacks.Melt(ctx)
	if err != nil {
		return err
	}

	toState.Actions.OnEnter()

	return nil
}

func (sm *Matter) Freeze(ctx context.Context) (err error) {
	from, ok := fromCtx(ctx)
	if !ok {
		return errors.New("\"from\" state not set in context")
	}

	fromState, err := sm.core.GetState(esmaq.StateType(from))
	if err != nil {
		return err
	}

	toState, err := sm.core.Transition(esmaq.EventType(EventFreeze), esmaq.StateType(from))
	if err != nil {
		return err
	}

	// inject "to" in context
	ctx = ctxWtTo(ctx, StateSolid)

	fromState.Actions.OnExit()

	err = sm.callbacks.Freeze(ctx)
	if err != nil {
		return err
	}

	toState.Actions.OnEnter()

	return nil
}

func (sm *Matter) Vaporize(ctx context.Context) (err error) {
	from, ok := fromCtx(ctx)
	if !ok {
		return errors.New("\"from\" state not set in context")
	}

	fromState, err := sm.core.GetState(esmaq.StateType(from))
	if err != nil {
		return err
	}

	toState, err := sm.core.Transition(esmaq.EventType(EventVaporize), esmaq.StateType(from))
	if err != nil {
		return err
	}

	// inject "to" in context
	ctx = ctxWtTo(ctx, StateGas)

	fromState.Actions.OnExit()

	err = sm.callbacks.Vaporize(ctx)
	if err != nil {
		return err
	}

	toState.Actions.OnEnter()

	return nil
}

func (sm *Matter) Condense(ctx context.Context) (err error) {
	from, ok := fromCtx(ctx)
	if !ok {
		return errors.New("\"from\" state not set in context")
	}

	fromState, err := sm.core.GetState(esmaq.StateType(from))
	if err != nil {
		return err
	}

	toState, err := sm.core.Transition(esmaq.EventType(EventCondense), esmaq.StateType(from))
	if err != nil {
		return err
	}

	// inject "to" in context
	ctx = ctxWtTo(ctx, StateLiquid)

	fromState.Actions.OnExit()

	err = sm.callbacks.Condense(ctx)
	if err != nil {
		return err
	}

	toState.Actions.OnEnter()

	return nil
}

func CtxWtFrom(ctx context.Context, from State) context.Context {
	return context.WithValue(ctx, fromKey, from)
}

func ctxWtTo(ctx context.Context, to State) context.Context {
	return context.WithValue(ctx, toKey, to)
}

func fromCtx(ctx context.Context) (State, bool) {
	from, ok := ctx.Value(fromKey).(State)
	return from, ok
}

func ToCtx(ctx context.Context) (State, bool) {
	to, ok := ctx.Value(toKey).(State)
	return to, ok
}

func NewMatter(actions *Actions, callbacks *Callbacks) *Matter {
	stateConfigs := []esmaq.StateConfig{
		{
			From:    esmaq.StateType(StateSolid),
			Actions: actions.Solid,
			Transitions: []esmaq.TransitionConfig{
				{
					Event: esmaq.EventType(EventMelt),
					To:    esmaq.StateType(StateLiquid),
				},
			},
		},
		{
			From:    esmaq.StateType(StateLiquid),
			Actions: actions.Liquid,
			Transitions: []esmaq.TransitionConfig{
				{
					Event: esmaq.EventType(EventFreeze),
					To:    esmaq.StateType(StateSolid),
				},
				{
					Event: esmaq.EventType(EventVaporize),
					To:    esmaq.StateType(StateGas),
				},
			},
		},
		{
			From:    esmaq.StateType(StateGas),
			Actions: actions.Gas,
			Transitions: []esmaq.TransitionConfig{
				{
					Event: esmaq.EventType(EventCondense),
					To:    esmaq.StateType(StateLiquid),
				},
			},
		},
	}

	matter := &Matter{
		core:      esmaq.NewCore(stateConfigs),
		callbacks: callbacks,
	}

	return matter
}