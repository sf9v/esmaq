package esmaq_test

import (
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
			Actions: &esmaq.Actions{
				OnEnter: func() {
					fmt.Println("enter: off")
				},
				OnExit: func() {
					fmt.Println("exit: off")
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
			Actions: &esmaq.Actions{
				OnEnter: func() {
					fmt.Println("enter: on")
				},
				OnExit: func() {
					fmt.Println("exit: on")
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
