package encoding

import (
	"encoding/json"
)

type Int interface {
	Int() int64
}

type Uint interface {
	Uint() uint64
}

type Bool interface {
	Bool() bool
}

type Float interface {
	Float() float64
}

type String interface {
	String() string
}

type Bytes interface {
	Bytes() []byte
}

type Value interface {
	Value() interface{}
}

type JsonMarshaler = json.Marshaler
