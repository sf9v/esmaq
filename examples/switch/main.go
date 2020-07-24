package main

import (
	"log"
	"os"

	"github.com/sf9v/esmaq/gen"
)

func main() {
	schema := gen.Schema{
		Name: "Switch",
		Pkg:  "main",
		States: []gen.State{
			{
				From: "on",
				Transitions: []gen.Transition{
					{
						To:    "off",
						Event: "switchOff",
					},
				},
			},
			{
				From: "off",
				Transitions: []gen.Transition{
					{
						To:    "on",
						Event: "switchOn",
					},
				},
			},
		},
	}

	f, err := os.Create("switch.go")
	if err != nil {
		log.Fatal(err)
	}

	err = gen.Generate(schema, f)
	if err != nil {
		log.Fatal(err)
	}
}
