package encoding

import (
	"reflect"
	"unsafe"
)

/** convert between []byte and string without memory occupied */

func BytesToStr(v []byte) (r string) {
	if v == nil {
		return ""
	}
	pb := (*reflect.SliceHeader)(unsafe.Pointer(&v))
	ps := (*reflect.StringHeader)(unsafe.Pointer(&r))
	ps.Data = pb.Data
	ps.Len = pb.Len
	return
}

func StrToBytes(v string) (r []byte) {
	if v == "" {
		return nil
	}
	pb := (*reflect.SliceHeader)(unsafe.Pointer(&r))
	ps := (*reflect.SliceHeader)(unsafe.Pointer(&v))
	pb.Data = ps.Data
	pb.Len = ps.Len
	pb.Cap = ps.Len
	return
}
