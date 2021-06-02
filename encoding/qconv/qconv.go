package qconv

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/saitofun/qlib/encoding"
	"github.com/saitofun/qlib/encoding/qjson"
	"github.com/saitofun/qlib/os/qtime"
)

func Int(v interface{}) int64     { return 0 }
func Uint(v interface{}) uint64   { return 0 }
func Float(v interface{}) float64 { return 0 }
func Bool(v interface{}) bool     { return true }
func Bytes(v interface{}) []byte  { return nil }
func Runes(v interface{}) []rune  { return nil }
func String(v interface{}) string {
	switch val := v.(type) {
	case int:
		return strconv.Itoa(val)
	case int8:
		return strconv.Itoa(int(val))
	case int16:
		return strconv.Itoa(int(val))
	case int32:
		return strconv.Itoa(int(val))
	case int64:
		return strconv.FormatInt(val, 10)
	case uint:
		return strconv.FormatUint(uint64(val), 10)
	case uint8:
		return strconv.FormatUint(uint64(val), 10)
	case uint16:
		return strconv.FormatUint(uint64(val), 10)
	case uint32:
		return strconv.FormatUint(uint64(val), 10)
	case uint64:
		return strconv.FormatUint(val, 10)
	case float32:
		return strconv.FormatFloat(float64(val), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(val)
	case string:
		return val
	case *string:
		return *val
	case []byte:
		return string(val)
	case *[]byte:
		return string(*val)
	case time.Time:
		if val.IsZero() {
			return ""
		}
		return val.String()
	case *time.Time:
		if val == nil || val.IsZero() {
			return ""
		}
		return val.String()
	case qtime.Time:
		return val.String()
	case *qtime.Time:
		if val == nil {
			return ""
		}
		return val.String()
	default:
		if val == nil {
			return ""
		}
		if vf, ok := val.(encoding.String); ok {
			return vf.String()
		}
		if vf, ok := val.(error); ok {
			return vf.Error()
		}
		rv := reflect.ValueOf(v)
		kind := rv.Kind()
		switch kind {
		case reflect.Chan, reflect.Map, reflect.Slice, reflect.Func,
			reflect.Ptr, reflect.Interface, reflect.UnsafePointer:
			if rv.IsNil() {
				return ""
			}
		case reflect.String:
			return rv.String()
		}
		if kind == reflect.Ptr {
			return String(rv.Elem().Interface())
		}
		if content, err := qjson.Marshal(val); err != nil {
			return fmt.Sprint(val)
		} else {
			return string(content)
		}
	}
}
