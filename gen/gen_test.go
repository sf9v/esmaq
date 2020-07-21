package gen_test

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stevenferrer/esmaq/gen"
	"github.com/stretchr/testify/assert"
)

func TestGen(t *testing.T) {
	path := "simple/simple.go"
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	output, err := os.Create(path)
	assert.NoError(t, err)
	// defer os.RemoveAll(path)

	err = genSimple(output)
	assert.NoError(t, err)
}

func genSimple(output io.Writer) error {

	states := []gen.State{
		{
			From: "a",
			Transitions: []gen.Transition{
				{
					Event: "a_to_b",
					To:    "b",
					Callback: gen.Callback{
						Ins: gen.Ins{
							"ii":   int(0),
							"ii32": int32(0),
							"ii64": int64(0),
						},
						Outs: gen.Outs{
							"oi":   int(0),
							"oi32": int32(0),
							// FIXME: there is an issue with the last param (err)
							// when we define more than 2 output parameters
							// "oi64": int64(0),
						},
					},
				},
				{
					// transition to self
					Event: "a_to_a",
					To:    "a",
					Callback: gen.Callback{
						Ins: gen.Ins{
							"iu":   uint(0),
							"iu32": uint32(0),
							"iu64": uint64(0),
						},
						Outs: gen.Outs{
							"of32": float32(0),
							"of64": float64(0),
						},
					},
				},
			},
		},
		{
			From: "b",
			Transitions: []gen.Transition{
				{
					To:    "c",
					Event: "b_to_c",
					Callback: gen.Callback{
						Ins: gen.Ins{
							"mis": "",
						},
						Outs: gen.Outs{
							"mos": "",
						},
					},
				},
				{
					To:    "a",
					Event: "b_to_a",
					Callback: gen.Callback{
						Ins: gen.Ins{
							"sp1": decimal.Decimal{},
						},
						Outs: gen.Outs{
							"sp2": "",
						},
					},
				},
			},
		},
	}

	schema := gen.Schema{
		Name:   "Simple",
		Pkg:    "simple",
		States: states,
	}

	return gen.Gen(schema, output)
}
