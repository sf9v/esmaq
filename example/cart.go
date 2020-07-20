package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/stevenferrer/esmaq"
	"github.com/stevenferrer/esmaq/example/internal/cart"
	"github.com/stevenferrer/esmaq/gen"
)

func cartExample() {
	myCart := cart.NewCart(&cart.Actions{
		Shopping: esmaq.Actions{
			OnExit: func(ctx context.Context) error {
				fmt.Println("shopping: exit")
				return nil
			},
		},
		Finalizing: esmaq.Actions{
			OnEnter: func(ctx context.Context) error {
				fmt.Println("finalizing: enter")
				return nil
			},
			OnExit: func(ctx context.Context) error {
				fmt.Println("finalizing: exit")
				return nil
			},
		},
		Paid: esmaq.Actions{
			OnEnter: func(ctx context.Context) error {
				fmt.Println("paid: enter")
				return nil
			},
			OnExit: func(ctx context.Context) error {
				fmt.Println("paid: exit")
				return nil
			},
		},
		Submitted: esmaq.Actions{
			OnEnter: func(ctx context.Context) error {
				fmt.Println("submitted: enter")
				return nil
			},
		},
	}, &cart.Callbacks{
		Checkout: func(ctx context.Context, cartID int64) (err error) {
			fmt.Printf("checkout cart %d\n", cartID)
			return nil
		},
		Pay: func(ctx context.Context, cartID, paymentID int64) (err error) {
			fmt.Printf("cart %d payment id %d\n", cartID, paymentID)
			return nil
		},
		Submit: func(ctx context.Context, cartID int64) (orderId int64, err error) {
			fmt.Printf("order submitted for cart %d\n", cartID)
			return 999, nil
		},
	})

	ctx := context.Background()

	cartID := int64(123)
	ctx = cart.CtxWtFrom(ctx, cart.StateShopping)
	err := myCart.Checkout(ctx, cartID)
	checkErr(err)

	paymentID := int64(55555)
	ctx = cart.CtxWtFrom(ctx, cart.StateFinalizing)
	err = myCart.Pay(ctx, cartID, paymentID)
	checkErr(err)

	ctx = cart.CtxWtFrom(ctx, cart.StatePaid)
	orderID, err := myCart.Submit(ctx, cartID)
	checkErr(err)

	fmt.Printf("order id is %d\n", orderID)
}

func generateCart() {
	myCart := []esmaq.StateConfig{
		{
			From:    "shopping",
			Actions: esmaq.Actions{},
			Transitions: []esmaq.TransitionConfig{
				{
					Event: "checkout",
					To:    "finalizing",
					Callback: esmaq.Callback{
						Ins: esmaq.Ins{
							"cartID": int64(0),
						},
					},
				},
			},
		},
		{
			From: "finalizing",
			Transitions: []esmaq.TransitionConfig{
				{
					Event: "pay",
					To:    "paid",
					Callback: esmaq.Callback{
						Ins: esmaq.Ins{
							"cartID":    int64(0),
							"paymentId": int64(0),
						},
					},
				},
			},
		},
		{
			From: "paid",
			Transitions: []esmaq.TransitionConfig{
				{
					Event: "submit",
					To:    "submitted",
					Callback: esmaq.Callback{
						Ins: esmaq.Ins{
							"cartID": int64(0),
						},
						Outs: esmaq.Outs{
							"orderId": int64(0),
						},
					},
				},
			},
		},
		{
			From: "submitted",
		},
	}

	filePath := "internal/cart/cart.go"
	err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	checkErr(err)

	out, err := os.Create(filePath)
	checkErr(err)
	err = gen.Gen(gen.Schema{
		Name:   "Cart",
		Pkg:    "cart",
		States: myCart,
	}, out)
	checkErr(err)
}
