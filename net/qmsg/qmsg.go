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
	GetTs() int64
	SetTs(int64)
}

type WithErrorCheck interface {
	Err() error
}
