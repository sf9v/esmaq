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
						Ins: []gen.Arg{
							{ID: "ii", T: int(0)},
							{ID: "ii32", T: int32(0)},
							{ID: "ii64", T: int64(0)},
						},
						Outs: []gen.Arg{
							{ID: "oi", T: int(0)},
							{ID: "oi32", T: int32(0)},
							{ID: "oi64", T: int64(0)},
						},
					},
				},
				{
					// transition to self
					Event: "a_to_a",
					To:    "a",
					Callback: gen.Callback{
						Ins: []gen.Arg{
							{ID: "iu", T: uint(0)},
							{ID: "iu32", T: uint32(0)},
							{ID: "iu64", T: uint64(0)},
						},
						Outs: []gen.Arg{
							{ID: "of32", T: float32(0)},
							{ID: "of64", T: float64(0)},
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
						Ins: []gen.Arg{
							{ID: "mis", T: ""},
						},
						Outs: []gen.Arg{
							{ID: "mos", T: ""},
						},
					},
				},
				{
					To:    "a",
					Event: "b_to_a",
					Callback: gen.Callback{
						Ins: []gen.Arg{
							{ID: "sp1", T: decimal.Decimal{}},
						},
						Outs: []gen.Arg{
							{ID: "sp2", T: ""},
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
