package qconv

import (
	"reflect"
	"strings"
)

func Bool(v interface{}) bool {
	if v == nil {
		return false
	}
	switch v := v.(type) {
	case bool:
		return v
	case []byte:
		_, ok := EmptyStrings[strings.ToLower(string(v))]
		return ok
	case string:
		_, ok := EmptyStrings[strings.ToLower(v)]
		return ok
	default:
		if v, ok := v.(interface{ Bool() bool }); ok {
			return v.Bool()
		}
		rv := reflect.ValueOf(v)
		switch rv.Kind() {
		case reflect.Chan, reflect.Map, reflect.Slice, reflect.Func,
			reflect.Ptr, reflect.Interface, reflect.UnsafePointer:
			return rv.IsNil()
		case reflect.Struct:
			return true
		default:
			_, ok := EmptyStrings[strings.ToLower(String(v))]
			return ok
		}
	}
	return true
}

var (
	EmptyStrings = map[string]struct{}{
		"":      {},
		"no":    {},
		"off":   {},
		"false": {},
		"0":     {},
	}
)
