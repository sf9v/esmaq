package matter

import (
	"context"
	"errors"
	esmaq "github.com/stevenferrer/esmaq"
)

type ctxKey int

const (
	fromKey ctxKey = iota
	toKey
)

func CtxWtFrom(ctx context.Context, from esmaq.State) context.Context {
	return context.WithValue(ctx, fromKey, from)
}

func ctxWtTo(ctx context.Context, to esmaq.State) context.Context {
	return context.WithValue(ctx, toKey, to)
}

func fromCtx(ctx context.Context) (esmaq.State, bool) {
	from, ok := ctx.Value(fromKey).(esmaq.State)
	return from, ok
}

func ToCtx(ctx context.Context) (esmaq.State, bool) {
	to, ok := ctx.Value(toKey).(esmaq.State)
	return to, ok
}

type Matter struct {
	core *esmaq.Core
	cbs  *Callbacks
}

type Callbacks struct {
	Melt     func(ctx context.Context) (err error)
	Freeze   func(ctx context.Context) (err error)
	Vaporize func(ctx context.Context) (err error)
	Condense func(ctx context.Context) (err error)
}

func (sm *Matter) Melt(ctx context.Context) (err error) {
	from, ok := fromCtx(ctx)
	if !ok {
		return errors.New("\"from\" state not set in context")
	}

	err = sm.core.Fire(EventMelt, from)
	if err != nil {
		return err
	}

	ctx = ctxWtTo(ctx, StateLiquid)

	err = sm.cbs.Melt(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (sm *Matter) Freeze(ctx context.Context) (err error) {
	from, ok := fromCtx(ctx)
	if !ok {
		return errors.New("\"from\" state not set in context")
	}

	err = sm.core.Fire(EventFreeze, from)
	if err != nil {
		return err
	}

	ctx = ctxWtTo(ctx, StateSolid)

	err = sm.cbs.Freeze(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (sm *Matter) Vaporize(ctx context.Context) (err error) {
	from, ok := fromCtx(ctx)
	if !ok {
		return errors.New("\"from\" state not set in context")
	}

	err = sm.core.Fire(EventVaporize, from)
	if err != nil {
		return err
	}

	ctx = ctxWtTo(ctx, StateGas)

	err = sm.cbs.Vaporize(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (sm *Matter) Condense(ctx context.Context) (err error) {
	from, ok := fromCtx(ctx)
	if !ok {
		return errors.New("\"from\" state not set in context")
	}

	err = sm.core.Fire(EventCondense, from)
	if err != nil {
		return err
	}

	ctx = ctxWtTo(ctx, StateLiquid)

	err = sm.cbs.Condense(ctx)
	if err != nil {
		return err
	}

	return nil
}

const (
	StateSolid  esmaq.State = "solid"
	StateLiquid esmaq.State = "liquid"
	StateGas    esmaq.State = "gas"
)

const (
	EventMelt     esmaq.Event = "melt"
	EventFreeze   esmaq.Event = "freeze"
	EventVaporize esmaq.Event = "vaporize"
	EventCondense esmaq.Event = "condense"
)

func NewMatter(cbs *Callbacks) *Matter {
	stateConfigs := []esmaq.StateConfig{
		{
			From: StateSolid,
			Transitions: []esmaq.Transitions{
				{
					Event: EventMelt,
					To:    StateLiquid,
				},
			},
		},
		{
			From: StateLiquid,
			Transitions: []esmaq.Transitions{
				{
					Event: EventFreeze,
					To:    StateSolid,
				},
				{
					Event: EventVaporize,
					To:    StateGas,
				},
			},
		},
		{
			From: StateGas,
			Transitions: []esmaq.Transitions{
				{
					Event: EventCondense,
					To:    StateLiquid,
				},
			},
		},
	}

	matter := &Matter{
		cbs:  cbs,
		core: esmaq.NewCore(stateConfigs),
	}

	return matter
}
