package qconv

import (
	"bytes"
	"encoding/binary"
	"strconv"
)

// Int convert any to int64
func Int(v interface{}) int64 {
	if v == nil {
		return 0
	}
	if v, ok := v.(int64); ok {
		return v
	}

	switch v := v.(type) {
	case int:
		return int64(v)
	case int8:
		return int64(v)
	case int16:
		return int64(v)
	case int32:
		return int64(v)
	case uint:
		return int64(v)
	case uint8:
		return int64(v)
	case uint16:
		return int64(v)
	case uint32:
		return int64(v)
	case uint64:
		return int64(v)
	case bool:
		return int64(Btoi(v))
	case []byte:
		if len(v) >= 4 {
			ret := int64(0)
			buf := bytes.NewBuffer(v[0:4])
			if err := binary.Read(buf, binary.LittleEndian, &ret); err == nil {
				return ret
			}
		}
		return 0
	default:
		if i, ok := v.(interface{ Int() int64 }); ok {
			return i.Int()
		}
		str := String(v)
		if i, err := strconv.ParseInt(str, 10, 64); err == nil {
			return i
		}
		if i, err := strconv.ParseInt(str, 16, 64); err == nil {
			return i
		}
		return int64(Float(v))
	}
}

// Uint convert any to int64
func Uint(v interface{}) uint64 {
	if v == nil {
		return 0
	}
	if v, ok := v.(uint64); ok {
		return v
	}

	switch v := v.(type) {
	case int:
		return uint64(v)
	case int8:
		return uint64(v)
	case int16:
		return uint64(v)
	case int32:
		return uint64(v)
	case uint:
		return uint64(v)
	case uint8:
		return uint64(v)
	case uint16:
		return uint64(v)
	case uint32:
		return uint64(v)
	case uint64:
		return uint64(v)
	case bool:
		return uint64(Btoi(v))
	case []byte:
		if len(v) >= 4 {
			ret := uint64(0)
			buf := bytes.NewBuffer(v[0:4])
			if err := binary.Read(buf, binary.LittleEndian, &ret); err == nil {
				return ret
			}
		}
		return 0
	default:
		if i, ok := v.(interface{ Uint() uint64 }); ok {
			return i.Uint()
		}
		str := String(v)
		if i, err := strconv.ParseUint(str, 10, 64); err == nil {
			return i
		}
		if i, err := strconv.ParseUint(str, 16, 64); err == nil {
			return i
		}
		return uint64(Float(v))
	}
}
