package qjson

import (
	"git.querycap.com/aisys/lib/encoding"
)

// CAUTION: the following functions ignore json marshal error

func UnsafeMarshal(v interface{}) []byte {
	ret, _ := Marshal(v)
	return ret
}

func UnsafeMarshalString(v interface{}) string {
	ret, _ := Marshal(v)
	return encoding.BytesToStr(ret)
}

func UnsafeMarshalIndent(v interface{}) string {
	ret, _ := MarshalIndent(v, "", "    ")
	return encoding.BytesToStr(ret)
}
