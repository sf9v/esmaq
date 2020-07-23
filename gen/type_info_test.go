package gen

import (
	"math/big"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/stevenferrer/esmaq/gen/internal"
	"github.com/stretchr/testify/assert"
)

func Test_getParamC(t *testing.T) {
	type args struct {
		id string
		v  interface{}
	}
	tests := []struct {
		name string
		args args
		want jen.Code
	}{
		{
			name: "int",
			args: args{
				id: "i",
				v:  int(0),
			},
			want: jen.Id("i").Int(),
		},
		{
			name: "*int",
			args: args{
				id: "i",
				v:  internal.IntPtr(0),
			},
			want: jen.Id("i").Op("*").Int(),
		},
		{
			name: "int32",
			args: args{
				id: "i",
				v:  int32(0),
			},
			want: jen.Id("i").Int32(),
		},
		{
			name: "*int32",
			args: args{
				id: "i",
				v:  internal.Int32Ptr(0),
			},
			want: jen.Id("i").Op("*").Int32(),
		},
		{
			name: "int64",
			args: args{
				id: "i",
				v:  int64(0),
			},
			want: jen.Id("i").Int64(),
		},
		{
			name: "*int64",
			args: args{
				id: "i",
				v:  internal.Int64Ptr(0),
			},
			want: jen.Id("i").Op("*").Int64(),
		},
		{
			name: "uint",
			args: args{
				id: "u",
				v:  uint(0),
			},
			want: jen.Id("u").Uint(),
		},
		{
			name: "*uint",
			args: args{
				id: "u",
				v:  internal.UintPtr(0),
			},
			want: jen.Id("u").Op("*").Uint(),
		},
		{
			name: "uint32",
			args: args{
				id: "u",
				v:  uint32(0),
			},
			want: jen.Id("u").Uint32(),
		},
		{
			name: "*uint32",
			args: args{
				id: "u",
				v:  internal.Uint32Ptr(0),
			},
			want: jen.Id("u").Op("*").Uint32(),
		},
		{
			name: "uint64",
			args: args{
				id: "u",
				v:  uint64(0),
			},
			want: jen.Id("u").Uint64(),
		},
		{
			name: "*uint64",
			args: args{
				id: "u",
				v:  internal.Uint64Ptr(0),
			},
			want: jen.Id("u").Op("*").Uint64(),
		},
		{
			name: "float32",
			args: args{
				id: "u",
				v:  float32(0),
			},
			want: jen.Id("u").Float32(),
		},
		{
			name: "*float32",
			args: args{
				id: "u",
				v:  internal.Float32Ptr(0),
			},
			want: jen.Id("u").Op("*").Float32(),
		},
		{
			name: "float64",
			args: args{
				id: "u",
				v:  float64(0),
			},
			want: jen.Id("u").Float64(),
		},
		{
			name: "*float64",
			args: args{
				id: "u",
				v:  internal.Float64Ptr(0),
			},
			want: jen.Id("u").Op("*").Float64(),
		},
		{
			name: "string",
			args: args{
				id: "s",
				v:  "",
			},
			want: jen.Id("s").String(),
		},
		{
			name: "*string",
			args: args{
				id: "s",
				v:  internal.StrPtr(""),
			},
			want: jen.Id("s").Op("*").String(),
		},
		{
			name: "big.Int",
			args: args{
				id: "bi",
				v:  big.Int{},
			},
			want: jen.Id("bi").Qual("math/big", "Int"),
		},

		{
			name: "*big.Int",
			args: args{
				id: "bi",
				v:  big.NewInt(0),
			},
			want: jen.Id("bi").Op("*").Qual("math/big", "Int"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getParamC(tt.args.id, tt.args.v)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_getZeroValC(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want jen.Code
	}{
		{
			name: "int",
			args: args{
				v: int(0),
			},
			want: jen.Lit(0),
		},
		{
			name: "*int",
			args: args{
				v: internal.IntPtr(0),
			},
			want: jen.Nil(),
		},
		{
			name: "int32",
			args: args{
				v: int32(0),
			},
			want: jen.Lit(0),
		},
		{
			name: "*int32",
			args: args{
				v: internal.Int32Ptr(0),
			},
			want: jen.Nil(),
		},
		{
			name: "int64",
			args: args{
				v: int64(0),
			},
			want: jen.Lit(0),
		},
		{
			name: "*int64",
			args: args{
				v: internal.Int64Ptr(0),
			},
			want: jen.Nil(),
		},
		{
			name: "uint",
			args: args{
				v: uint(0),
			},
			want: jen.Lit(0),
		},
		{
			name: "*uint",
			args: args{
				v: internal.UintPtr(0),
			},
			want: jen.Nil(),
		},
		{
			name: "uint32",
			args: args{
				v: uint32(0),
			},
			want: jen.Lit(0),
		},
		{
			name: "*uint32",
			args: args{
				v: internal.Uint32Ptr(0),
			},
			want: jen.Nil(),
		},
		{
			name: "uint64",
			args: args{
				v: uint64(0),
			},
			want: jen.Lit(0),
		},
		{
			name: "*uint64",
			args: args{
				v: internal.Uint64Ptr(0),
			},
			want: jen.Nil(),
		},
		{
			name: "float32",
			args: args{
				v: float32(0),
			},
			want: jen.Lit(0),
		},
		{
			name: "*float32",
			args: args{
				v: internal.Float32Ptr(0),
			},
			want: jen.Nil(),
		},
		{
			name: "float64",
			args: args{
				v: float64(0),
			},
			want: jen.Lit(0),
		},
		{
			name: "*float64",
			args: args{
				v: internal.Float64Ptr(0),
			},
			want: jen.Nil(),
		},
		{
			name: "string",
			args: args{
				v: "",
			},
			want: jen.Lit(""),
		},
		{
			name: "*string",
			args: args{
				v: internal.StrPtr(""),
			},
			want: jen.Nil(),
		},
		{
			name: "big.Int",
			args: args{
				v: big.Int{},
			},
			want: jen.Qual("math/big", "Int").Op("{").Op("}"),
		},
		{
			name: "*big.Int",
			args: args{
				v: big.NewInt(0),
			},
			want: jen.Nil(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getZeroValC(tt.args.v)
			assert.Equal(t, tt.want, got)
		})
	}
}
