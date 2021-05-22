package qjson

import (
	"git.querycap.com/ss/lib/encoding"
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

func UnmarshalString(content string, v interface{}) error {
	return Unmarshal(encoding.StrToBytes(content), v)
}
