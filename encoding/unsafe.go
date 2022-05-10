package encoding

import (
	"unsafe"
)

// BytesToStr converts from []byte to string without memory copy
func BytesToStr(v []byte) (r string) {
	return *(*string)(unsafe.Pointer(&v))
}

// StrToBytes convert from string to []byte without memory copy
func StrToBytes(v string) (r []byte) {
	return *(*[]byte)(unsafe.Pointer(&v))
}
