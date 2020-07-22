package simple_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stevenferrer/esmaq"
	"github.com/stevenferrer/esmaq/gen/simple"
	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {

	t.Run("stringer", func(t *testing.T) {
		assert.Equal(t, "a", simple.StateA.String())
		assert.Equal(t, "a_to_a", simple.EventAToA.String())
	})

	t.Run("ok", func(t *testing.T) {
		callCounts := map[string]int{}

		key := func(s simple.State, a string) string {
			return fmt.Sprintf("%s_%s", s, a)
		}

		sm := simple.NewSimple(&simple.Callbacks{
			AToB: func(ctx context.Context, ii int, ii32 int32, ii64 int64) (oi int, oi32 int32, oi64 int64, err error) {
				to, ok := simple.ToCtx(ctx)
				if !ok {
					return 0, 0, 0, errors.New("to not injected in ctx")
				}

				assert.Equal(t, simple.StateB, to)
				return 0, 0, 0, nil
			},
			AToA: func(ctx context.Context, iu uint, iu32 uint32, iu64 uint64) (of32 float32, of64 float32, err error) {
				return 0, 0, nil
			},
			BToA: func(ctx context.Context, sp1 decimal.Decimal) (sp2 string, err error) {
				return "", nil
			},
			BToC: func(ctx context.Context, mis string) (mos string, err error) {
				return "", nil
			},
		}, &simple.Actions{
			A: esmaq.Actions{
				OnEnter: func(_ context.Context) error {
					k := key(simple.StateA, "enter")
					if _, ok := callCounts[k]; !ok {
						callCounts[k] = 0
					}
					callCounts[k]++
					return nil
				},
				OnExit: func(_ context.Context) error {
					k := key(simple.StateA, "exit")
					if _, ok := callCounts[k]; !ok {
						callCounts[k] = 0
					}
					callCounts[k]++
					return nil
				},
			},
			B: esmaq.Actions{
				OnEnter: func(_ context.Context) error {
					k := key(simple.StateB, "enter")
					if _, ok := callCounts[k]; !ok {
						callCounts[k] = 0
					}
					callCounts[k]++
					return nil
				},
				OnExit: func(_ context.Context) error {
					k := key(simple.StateB, "exit")
					if _, ok := callCounts[k]; !ok {
						callCounts[k] = 0
					}
					callCounts[k]++
					return nil
				},
			},
			C: esmaq.Actions{
				OnEnter: func(_ context.Context) error {
					k := key(simple.StateC, "enter")
					if _, ok := callCounts[k]; !ok {
						callCounts[k] = 0
					}
					callCounts[k]++
					return nil
				},
				OnExit: func(_ context.Context) error {
					k := key(simple.StateC, "exit")
					if _, ok := callCounts[k]; !ok {
						callCounts[k] = 0
					}
					callCounts[k]++
					return nil
				},
			},
		})

		ctx := context.Background()
		ctx = simple.CtxWtFrom(ctx, simple.StateA)

		_, _, err := sm.AToA(ctx, 0, 0, 0)
		assert.NoError(t, err)

		_, _, _, err = sm.AToB(ctx, 0, 0, 0)
		assert.NoError(t, err)

		ctx = simple.CtxWtFrom(ctx, simple.StateB)
		_, err = sm.BToA(ctx, decimal.NewFromFloat(0))
		assert.NoError(t, err)

		_, err = sm.BToC(ctx, "")
		assert.NoError(t, err)

		// TODO: assert call counts
	})

	t.Run("errors", func(t *testing.T) {
		t.Run("from state not set", func(t *testing.T) {
			sm := simple.NewSimple(&simple.Callbacks{
				AToA: func(ctx context.Context, iu uint, iu32 uint32, iu64 uint64) (of32 float32, of64 float32, err error) {
					return 0, 0, nil
				},
				AToB: func(ctx context.Context, ii int, ii32 int32, ii64 int64) (oi int, oi32 int32, oi64 int64, err error) {
					return 0, 0, 0, nil
				},
				BToA: func(ctx context.Context, sp1 decimal.Decimal) (sp2 string, err error) {
					return "", nil
				},
				BToC: func(ctx context.Context, mis string) (mos string, err error) {
					return "", nil
				},
			}, &simple.Actions{
				A: esmaq.Actions{
					OnEnter: func(_ context.Context) error {
						return errors.New("enter a error")
					},
					OnExit: func(_ context.Context) error {
						return errors.New("exit a error")
					},
				},
				B: esmaq.Actions{
					OnEnter: func(_ context.Context) error {
						return errors.New("enter b error")
					},
					OnExit: func(_ context.Context) error {
						return errors.New("exit b error")
					},
				},
				C: esmaq.Actions{
					OnEnter: func(_ context.Context) error {
						return errors.New("enter c error")
					},
					OnExit: func(_ context.Context) error {
						return errors.New("exit c error")
					},
				},
			})

			ctx := context.Background()
			_, _, err := sm.AToA(ctx, 0, 0, 0)
			assert.Error(t, err)

			_, _, _, err = sm.AToB(ctx, 0, 0, 0)
			assert.Error(t, err)

			_, err = sm.BToA(ctx, decimal.NewFromFloat(0))
			assert.Error(t, err)

			_, err = sm.BToC(ctx, "")
			assert.Error(t, err)
		})

		t.Run("action errors", func(t *testing.T) {
			sm := simple.NewSimple(&simple.Callbacks{
				AToA: func(ctx context.Context, iu uint, iu32 uint32, iu64 uint64) (of32 float32, of64 float32, err error) {
					return 0, 0, nil
				},
				AToB: func(ctx context.Context, ii int, ii32 int32, ii64 int64) (oi int, oi32 int32, oi64 int64, err error) {
					return 0, 0, 0, nil
				},
				BToA: func(ctx context.Context, sp1 decimal.Decimal) (sp2 string, err error) {
					return "", nil
				},
				BToC: func(ctx context.Context, mis string) (mos string, err error) {
					return "", nil
				},
			}, &simple.Actions{
				A: esmaq.Actions{
					OnEnter: func(_ context.Context) error {
						return errors.New("enter a error")
					},
					OnExit: func(_ context.Context) error {
						return errors.New("exit a error")
					},
				},
				B: esmaq.Actions{
					OnEnter: func(_ context.Context) error {
						return errors.New("enter b error")
					},
					OnExit: func(_ context.Context) error {
						return errors.New("exit b error")
					},
				},
				C: esmaq.Actions{
					OnEnter: func(_ context.Context) error {
						return errors.New("enter c error")
					},
					OnExit: func(_ context.Context) error {
						return errors.New("exit c error")
					},
				},
			})

			ctx := context.Background()
			ctx = simple.CtxWtFrom(ctx, simple.StateA)
			_, _, err := sm.AToA(ctx, 0, 0, 0)
			assert.Error(t, err)

			_, _, _, err = sm.AToB(ctx, 0, 0, 0)
			assert.Error(t, err)

			ctx = simple.CtxWtFrom(ctx, simple.StateB)
			_, err = sm.BToA(ctx, decimal.NewFromFloat(0))
			assert.Error(t, err)

			_, err = sm.BToC(ctx, "")
			assert.Error(t, err)
		})
	})
}
