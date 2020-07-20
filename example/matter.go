package main

import (
	"os"
	"path/filepath"

	"github.com/stevenferrer/esmaq"
	"github.com/stevenferrer/esmaq/example/internal/matter"
	"github.com/stevenferrer/esmaq/gen"
)

func matterExample() {
	_ = matter.NewMatter(&matter.Actions{}, &matter.Callbacks{})
}

func generateMatter() {
	statesOfMatter := []esmaq.StateConfig{
		{
			From: "solid",
			Transitions: []esmaq.TransitionConfig{
				{
					Event: "melt",
					To:    "liquid",
				},
			},
		},
		{
			From: "liquid",
			Transitions: []esmaq.TransitionConfig{
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
			Transitions: []esmaq.TransitionConfig{
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
