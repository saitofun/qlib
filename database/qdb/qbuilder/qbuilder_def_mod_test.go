package qbuilder

import (
	"reflect"
	"testing"
)

func TestReflectValueName(t *testing.T) {
	type V struct {
		A int
		B int
		C int
	}

	v := reflect.TypeOf(V{})
	t.Log(v.Name())
	t.Log(v.String())

	rt := reflect.StructOf([]reflect.StructField{v.Field(0), v.Field(2)})
	t.Log(rt.Name())
	t.Log(rt.String())

	svt := reflect.SliceOf(reflect.New(rt).Elem().Type())
	rv := reflect.MakeSlice(svt, 0, 0)
	t.Log(rv.Type().Name())
	t.Log(rv.Type().String())
	t.Log(rv.Interface())
}
