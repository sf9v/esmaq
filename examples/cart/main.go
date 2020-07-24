package main

import (
	"log"
	"os"

	"github.com/sf9v/esmaq/gen"
)

func main() {
	schema := gen.Schema{
		Name: "Cart",
		Pkg:  "main",
		States: []gen.State{
			{
				From: "new",
				Transitions: []gen.Transition{
					{
						To:    "finalizing",
						Event: "checkout",
						Callback: gen.Callback{
							Ins: []gen.Param{{ID: "cartID", V: int64(0)}},
						},
					},
				},
			},
			{
				From: "finalizing",
				Transitions: []gen.Transition{
					{
						To:    "submitted",
						Event: "submit",
						Callback: gen.Callback{
							Ins:  []gen.Param{{ID: "cartID", V: int64(0)}},
							Outs: []gen.Param{{ID: "orderID", V: int64(0)}},
						},
					},
					{
						To:    "new",
						Event: "modify",
						Callback: gen.Callback{
							Ins: []gen.Param{{ID: "cartID", V: int64(0)}},
						},
					},
					{
						To:    "cancelled",
						Event: "cancel",
						Callback: gen.Callback{
							Ins: []gen.Param{{ID: "cartID", V: int64(0)}},
						},
					},
				},
			},
		},
	}

	f, err := os.Create("cart.go")
	if err != nil {
		log.Fatal(err)
	}

	err = gen.Generate(schema, f)
	if err != nil {
		log.Fatal(err)
	}
}
