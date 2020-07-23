package gen

import (
	"bytes"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	states := []State{
		{
			From: "a",
			Transitions: []Transition{
				{
					Event: "a_to_b",
					To:    "b",
					Callback: Callback{
						Ins: []Param{
							{ID: "x", V: big.NewInt(0)},
						},
						Outs: []Param{
							{ID: "y", V: big.NewFloat(0)},
						},
					},
				},
				{
					// transition to self
					Event: "a_to_a",
					To:    "a",
					Callback: Callback{
						Ins: []Param{
							{ID: "a", V: big.Int{}},
						},
						Outs: []Param{
							{ID: "b", V: big.Float{}},
						},
					},
				},
			},
		},
		{
			From: "b",
			Transitions: []Transition{
				{
					To:    "c",
					Event: "b_to_c",
				},
				{
					To:    "a",
					Event: "b_to_a",
				},
			},
		},
	}

	schema := Schema{
		Name:   "Simple",
		Pkg:    "simple",
		States: states,
	}

	err := Generate(schema, &bytes.Buffer{})
	assert.NoError(t, err)
}
