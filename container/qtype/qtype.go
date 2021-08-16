package qtype

import "reflect"

type WithType interface {
	Type() reflect.Type
}

type WithJSON interface {
	MarshalJSON() ([]byte, error)
	UnmarshalJSON([]byte) error
}
