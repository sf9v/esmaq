package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/stevenferrer/esmaq"
	"github.com/stevenferrer/esmaq/example/internal/myswitch"
	"github.com/stevenferrer/esmaq/gen"
)

func switchExample() {
	mySwitch := myswitch.NewMySwitch(&myswitch.Actions{
		Off: esmaq.Actions{
			OnEnter: func(_ context.Context) error {
				return errors.New("off: on enter: something bad happened")
			},
		},
		On: esmaq.Actions{},
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

	ctx := myswitch.CtxWtFrom(context.Background(), myswitch.StateOff)
	_, err := mySwitch.SwitchOn(ctx, 0)
	checkErr(err)

	ctx = myswitch.CtxWtFrom(ctx, myswitch.StateOn)
	err = mySwitch.SwitchOff(ctx)
	checkErr(err)
}

func generateSwitch() {
	mySwitch := []gen.State{
		{
			From: "off",
			Transitions: []gen.Transition{
				{
					Event: "switchOn",
					To:    "on",
					Callback: gen.Callback{
						Ins: gen.Ins{
							"a": 0,
						},
						Outs: gen.Outs{"b": ""},
					},
				},
			},
		},
		{
			From: "on",
			Transitions: []gen.Transition{
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
