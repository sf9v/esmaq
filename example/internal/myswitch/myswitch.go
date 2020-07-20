package myswitch

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

type MySwitch struct {
	core *esmaq.Core
	cbs  *Callbacks
}

type Callbacks struct {
	SwitchOn  func(ctx context.Context) (err error)
	SwitchOff func(ctx context.Context) (err error)
}

func (sm *MySwitch) SwitchOn(ctx context.Context) (err error) {
	from, ok := fromCtx(ctx)
	if !ok {
		return errors.New("\"from\" state not set in context")
	}

	err = sm.core.Fire(EventSwitchOn, from)
	if err != nil {
		return err
	}

	ctx = ctxWtTo(ctx, StateOn)

	err = sm.cbs.SwitchOn(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (sm *MySwitch) SwitchOff(ctx context.Context) (err error) {
	from, ok := fromCtx(ctx)
	if !ok {
		return errors.New("\"from\" state not set in context")
	}

	err = sm.core.Fire(EventSwitchOff, from)
	if err != nil {
		return err
	}

	ctx = ctxWtTo(ctx, StateOff)

	err = sm.cbs.SwitchOff(ctx)
	if err != nil {
		return err
	}

	return nil
}

const (
	StateOff esmaq.State = "off"
	StateOn  esmaq.State = "on"
)

const (
	EventSwitchOn  esmaq.Event = "switchOn"
	EventSwitchOff esmaq.Event = "switchOff"
)

func NewMySwitch(cbs *Callbacks) *MySwitch {
	stateConfigs := []esmaq.StateConfig{
		{
			From: StateOff,
			Transitions: []esmaq.Transitions{
				{
					Event: EventSwitchOn,
					To:    StateOn,
				},
			},
		},
		{
			From: StateOn,
			Transitions: []esmaq.Transitions{
				{
					Event: EventSwitchOff,
					To:    StateOff,
				},
			},
		},
	}

	mySwitch := &MySwitch{
		cbs:  cbs,
		core: esmaq.NewCore(stateConfigs),
	}

	return mySwitch
}
