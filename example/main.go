package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/stevenferrer/esmaq"
	"github.com/stevenferrer/esmaq/example/internal/matter"
	"github.com/stevenferrer/esmaq/example/internal/myswitch"
)

func main() {
	switchExample()
}

func switchExample() {
	mySwitch := myswitch.NewMySwitch(&myswitch.Callbacks{
		SwitchOff: func(ctx context.Context) (err error) {
			fmt.Println("switched off!")
			return
		},
		SwitchOn: func(ctx context.Context) (err error) {
			fmt.Println("switched on")
			return
		},
	})
	ctx := context.Background()

	err := mySwitch.SwitchOn(myswitch.CtxWtFrom(ctx, myswitch.StateOff))
	checkErr(err)

}

func generateSwitch() {
	lightBulb := []esmaq.StateConfig{
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
	}

	filePath := "internal/myswitch/myswitch.go"
	err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	checkErr(err)

	out, err := os.Create(filePath)
	checkErr(err)
	err = esmaq.Gen(esmaq.Schema{
		Name:   "MySwitch",
		Pkg:    "myswitch",
		States: lightBulb,
	}, out)
	checkErr(err)
}

func matterExample() {
	myMatter := matter.NewMatter(&matter.Callbacks{
		Condense: func(ctx context.Context) (err error) {
			fmt.Print("condensed")
			return
		},
		Freeze: func(ctx context.Context) (err error) {
			fmt.Println("freezed")
			return
		},
		Melt: func(ctx context.Context) (err error) {
			fmt.Println("melted")
			return
		},
		Vaporize: func(ctx context.Context) (err error) {
			fmt.Println("liquid vaporized")
			next, ok := matter.ToCtx(ctx)
			if !ok {
				panic("to state not set")
			}

			fmt.Printf("next state is %q\n", next)
			return
		},
	})

	ctx := context.Background()

	err := myMatter.Vaporize(matter.CtxWtFrom(ctx, matter.StateLiquid))
	checkErr(err)
}

func generateMatter() {
	statesOfMatter := []esmaq.StateConfig{
		{
			From: "solid",
			Transitions: []esmaq.Transitions{
				{
					Event: "melt",
					To:    "liquid",
				},
			},
		},
		{
			From: "liquid",
			Transitions: []esmaq.Transitions{
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
			Transitions: []esmaq.Transitions{
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

	err = esmaq.Gen(esmaq.Schema{
		Name:   "Matter",
		Pkg:    "matter",
		States: statesOfMatter,
	}, out)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
