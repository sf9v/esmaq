package main

import (
	"os"
	"path/filepath"

	"github.com/stevenferrer/esmaq/example/internal/matter"
	"github.com/stevenferrer/esmaq/gen"
)

func matterExample() {
	_ = matter.NewMatter(&matter.Actions{}, &matter.Callbacks{})
}

func generateMatter() {
	statesOfMatter := []gen.State{
		{
			From: "solid",
			Transitions: []gen.Transition{
				{
					Event: "melt",
					To:    "liquid",
				},
			},
		},
		{
			From: "liquid",
			Transitions: []gen.Transition{
				{
					Event: "freeze",
					To:    "solid",
				},
				{
					Event: "vaporize",
					To:    "gas",
				},
			},
		},
		{
			From: "gas",
			Transitions: []gen.Transition{
				{
					Event: "condense",
					To:    "liquid",
				},
			},
		},
	}

	filePath := "internal/matter/matter.go"
	err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	checkErr(err)

	out, err := os.Create(filePath)
	checkErr(err)

	err = gen.Gen(gen.Schema{
		Name:   "Matter",
		Pkg:    "matter",
		States: statesOfMatter,
	}, out)
	checkErr(err)
}
