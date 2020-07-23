package gen

import (
	"reflect"

	"github.com/dave/jennifer/jen"
)

type typeInfo struct {
	t        reflect.Type
	v        interface{}
	name     string
	pkgPath  string
	nillable bool
}

func newTypeInfo(v interface{}) *typeInfo {
	t := reflect.TypeOf(v)
	t2 := indirect(t)
	return &typeInfo{
		t:        t2,
		v:        v,
		name:     t2.Name(),
		pkgPath:  t2.PkgPath(),
		nillable: isNillable(t),
	}
}

func getParamC(id string, v interface{}) jen.Code {
	info := newTypeInfo(v)
	c := jen.Id(id)
	if info.nillable {
		c = c.Op("*")
	}

	// built-in types
	if info.pkgPath == "" {
		switch info.t.Kind() {
		case reflect.Int:
			return c.Int()
		case reflect.Int32:
			return c.Int32()
		case reflect.Int64:
			return c.Int64()
		case reflect.Uint:
			return c.Uint()
		case reflect.Uint32:
			return c.Uint32()
		case reflect.Uint64:
			return c.Uint64()
		case reflect.Float32:
			return c.Float32()
		case reflect.Float64:
			return c.Float64()
		case reflect.String:
			return c.String()
		}
	}

	return c.Qual(info.pkgPath, info.name)
}

func getZeroValC(v interface{}) jen.Code {
	info := newTypeInfo(v)
	if info.nillable {
		return jen.Nil()
	}

	// built-in types
	if info.pkgPath == "" {
		switch info.t.Kind() {
		case reflect.Int, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64:
			return jen.Lit(0)
		case reflect.String:
			return jen.Lit("")
		}
	}

	return jen.Qual(info.pkgPath, info.name).Op("{").Op("}")
}

func isNillable(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Slice, reflect.Array, reflect.Ptr, reflect.Map:
		return true
	}
	return false
}
