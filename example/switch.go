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

	ctx := context.Background()
	_, err := mySwitch.SwitchOn(myswitch.CtxWtFrom(ctx, myswitch.StateOff), 0)
	checkErr(err)

	err = mySwitch.SwitchOff(myswitch.CtxWtFrom(ctx, myswitch.StateOn))
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
