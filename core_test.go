package esmaq_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stevenferrer/esmaq"
	"github.com/stretchr/testify/assert"
)

func TestStateMachine(t *testing.T) {
	core := esmaq.NewCore([]esmaq.StateConfig{
		{
			From: "off",
			Actions: esmaq.Actions{
				OnEnter: func(_ context.Context) error {
					fmt.Println("enter: off")
					return nil
				},
				OnExit: func(_ context.Context) error {
					fmt.Println("exit: off")
					return nil
				},
			},
			Transitions: []esmaq.TransitionConfig{
				{
					Event: "switch",
					To:    "on",
				},
			},
		},
		{
			From: "on",
			Actions: esmaq.Actions{
				OnEnter: func(_ context.Context) error {
					fmt.Println("enter: on")
					return nil
				},
				OnExit: func(_ context.Context) error {
					fmt.Println("exit: on")
					return nil
				},
			},
			Transitions: []esmaq.TransitionConfig{
				{
					Event: "switch",
					To:    "off",
				},
			},
		},
	})

	ts, err := core.Transition("switch", "off")
	assert.NoError(t, err)
	spew.Dump(ts)
}
