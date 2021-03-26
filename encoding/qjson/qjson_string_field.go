package qjson

import (
	"git.querycap.com/aisys/lib/encoding"
)

var Stringer = struct {
	Marshal   func(interface{}) ([]byte, error)
	Unmarshal func([]byte, interface{}) error
}{
	Marshal: func(v interface{}) ([]byte, error) {
		raw, err := Marshal(v)
		if err != nil {
			return nil, err
		}
		return Marshal(encoding.BytesToStr(raw))
	},
	Unmarshal: func(data []byte, v interface{}) error {
		var raw string
		var err = Unmarshal(data, &raw)
		if err != nil {
			return err
		}
		return Unmarshal(encoding.StrToBytes(raw), v)
	},
}
