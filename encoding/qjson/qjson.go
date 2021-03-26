package qjson

type Encoder interface {
	Marshaler
	Unmarshaler
}

type Marshaler interface {
	MarshalJSON() ([]byte, error)
}

type Unmarshaler interface {
	UnmarshalJSON([]byte) error
}

type EmptyEncoder interface{}
