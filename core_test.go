package esmaq_test

import (
	"testing"

	"github.com/stevenferrer/esmaq"
	"github.com/stretchr/testify/assert"
)

func TestCore(t *testing.T) {
	core := esmaq.NewCore([]esmaq.StateConfig{
		{
			From: "a",
			Transitions: []esmaq.TransitionConfig{
				{
					Event: "to_b",
					To:    "b",
				},
				{
					// transition to self
					Event: "to_a",
					To:    "a",
				},
			},
		},
		{
			From: "b",
			Transitions: []esmaq.TransitionConfig{
				{
					To:    "c",
					Event: "to_c",
				},
				{
					To:    "a",
					Event: "to_a",
				},
				// duplicates
				{
					To:    "c",
					Event: "to_c",
				},
			},
		},
		{
			// duplicates
			From: "b",
		},
	})

	// valid transition
	next, err := core.Transition("a", "to_b")
	assert.NoError(t, err)
	assert.Equal(t, "b", string(next.State))

	// transition to self
	next, err = core.Transition("a", "to_a")
	assert.NoError(t, err)
	assert.Equal(t, "a", string(next.State))

	next, err = core.Transition("b", "to_c")
	assert.NoError(t, err)
	assert.Equal(t, "c", string(next.State))

	next, err = core.Transition("b", "to_a")
	assert.NoError(t, err)
	assert.Equal(t, "a", string(next.State))

	// invalid transitions
	_, err = core.Transition("a", "to_c")
	assert.Error(t, err)

	_, err = core.Transition("c", "to_a")
	assert.Error(t, err)

	// invalid states
	_, err = core.Transition("a", "to_d")
	assert.Error(t, err)

	_, err = core.Transition("d", "to_e")
	assert.Error(t, err)
}
