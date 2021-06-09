// Package qmsg provides interfaces for universal message
package qmsg

type Type interface {
	String() string
}

type ID interface {
	String() string
}

type Message interface {
	ID() ID
	Type() Type
}

type WithTimestamp interface {
	GetTimestamp() int64
}

type WithErrorCheck interface {
	Err() error
}

type WithReused interface {
	Renew()
	RenewWithID(ID) // for responding
}
