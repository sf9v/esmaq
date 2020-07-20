package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/stevenferrer/esmaq"
	"github.com/stevenferrer/esmaq/gen"

	"github.com/stevenferrer/esmaq/example/internal/myswitch"
)

func main() {
	generateSwitch()
	generateMatter()
}

func switchExample() {
	mySwitch := myswitch.NewMySwitch(&myswitch.Actions{
		Off: &esmaq.Actions{
			OnEnter: func() {
				fmt.Println("off: on enter")
			},
			OnExit: func() {
				fmt.Println("off: on exit")
			},
		},
		On: &esmaq.Actions{
			OnEnter: func() {
				fmt.Println("on: on enter")
			},
			OnExit: func() {
				fmt.Println("on: on exit")
			},
		},
	}, &myswitch.Callbacks{
		SwitchOn: func(ctx context.Context, a int) (b string, err error) {
			fmt.Println("switch on callback")
			return
		},
		SwitchOff: func(ctx context.Context) (err error) {
			fmt.Println("switch off callback")
			return
		},
	})

	ctx := context.Background()
	_, err := mySwitch.SwitchOn(myswitch.CtxWtFrom(ctx, myswitch.StateOff), 0)
	checkErr(err)
}

func generateSwitch() {
	mySwitch := []esmaq.StateConfig{
		{
			From: "off",
			Transitions: []esmaq.TransitionConfig{
				{
					Event: "switchOn",
					To:    "on",
					Callback: esmaq.Callback{
						Ins: esmaq.Ins{
							"a": 0,
						},
						Outs: esmaq.Outs{"b": ""},
					},
				},
			},
		},
		{
			From: "on",
			Transitions: []esmaq.TransitionConfig{
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
	err = gen.Gen(gen.Schema{
		Name:   "MySwitch",
		Pkg:    "myswitch",
		States: mySwitch,
	}, out)
	checkErr(err)
}

func matterExample() {
	// myMatter := matter.NewMatter(&matter.Actions{
	// 	Gas:    &esmaq.Actions{
	// 		OnEnter: func(){
	// 			fmt.Println("gas: on enter")
	// 		},
	// 	},
	// 	Liquid: &esmaq.Actions{},
	// 	Solid:  &esmaq.Actions{},
	// }, &matter.Callbacks{
	// 	Condense: func(ctx context.Context) (err error) {
	// 		return
	// 	},
	// 	Freeze: func(ctx context.Context) (err error) {
	// 		return
	// 	},
	// 	Melt: func(ctx context.Context) (err error) {
	// 		return
	// 	},
	// 	Vaporize: func(ctx context.Context) (err error) {
	// 		return
	// 	},
	// })
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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
