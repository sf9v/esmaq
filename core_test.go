package esmaq_test

import (
	"testing"

	"github.com/stevenferrer/esmaq"
	"github.com/stretchr/testify/assert"
)

func TestStateMachine(t *testing.T) {
	core := esmaq.NewCore([]esmaq.StateConfig{
		{
			From: "off",
			Transitions: []esmaq.Transitions{
				{
					Event: "switchOn",
					To:    "on",
				},
			},
		},
		{
			From: "on",
			Transitions: []esmaq.Transitions{
				{
					Event: "switchOff",
					To:    "off",
				},
			},
		},
	})

	err := core.CanTransition("switchOn", "on")
	assert.NoError(t, err)
}
