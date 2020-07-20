package main

import (
	"os"
	"path/filepath"

	"github.com/stevenferrer/esmaq"
)

func main() {
	// switchExample()
	generateSwitch()
	generateMatter()
}

func switchExample() {
	// eventHandlers := myswitch.EventHandlers{
	// 	SwitchOff: &myswitch.SwitchOffEventHandlers{
	// 		OnBefore: func(ctx context.Context) error {
	// 			fmt.Println("")
	// 			return nil
	// 		},
	// 	},
	// }

	// mySwitch := myswitch.NewMySwitch()
	// ctx := context.Background()

	// err := mySwitch.SwitchOn(myswitch.CtxWtFrom(ctx, myswitch.StateOff))
	// checkErr(err)
}

func generateSwitch() {
	mySwitch := []esmaq.StateConfig{
		{
			From: "off",
			Transitions: []esmaq.Transitions{
				{
					Event: "switchOn",
					To:    "on",
					Callback: esmaq.Callback{
						Ins: esmaq.Ins{
							"a": 1,
						},
						Outs: esmaq.Outs{
							"b": 1.0,
						},
					},
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
		States: mySwitch,
	}, out)
	checkErr(err)
}

func matterExample() {
	// myMatter := matter.NewMatter(&matter.Callbacks{
	// 	Condense: func(ctx context.Context) (err error) {
	// 		fmt.Print("condensed")
	// 		return
	// 	},
	// 	Freeze: func(ctx context.Context) (err error) {
	// 		fmt.Println("freezed")
	// 		return
	// 	},
	// 	Melt: func(ctx context.Context) (err error) {
	// 		fmt.Println("melted")
	// 		return
	// 	},
	// 	Vaporize: func(ctx context.Context) (err error) {
	// 		fmt.Println("liquid vaporized")
	// 		next, ok := matter.ToCtx(ctx)
	// 		if !ok {
	// 			panic("to state not set")
	// 		}

	// 		fmt.Printf("next state is %q\n", next)
	// 		return
	// 	},
	// })

	// ctx := context.Background()

	// err := myMatter.Vaporize(matter.CtxWtFrom(ctx, matter.StateLiquid))
	// checkErr(err)
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
