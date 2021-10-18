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

type NamedMessage interface {
	Name() string
}

type WithTimestamp interface {
	GetTs() int64
	SetTs()
}

type WithErrorCheck interface {
	Err() error
}
