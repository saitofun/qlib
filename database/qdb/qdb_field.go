package qdb

type Field interface {
	Name() string
	Type() string
	Limitation() interface{}
}

type WithCond interface {
	Cond() interface{}
}

type WithDefault interface {
	Default() interface{}
}

type PrimaryID interface {
	PrimaryID() uint64
}

type ForeignKey interface {
	From() Field
	Current() Field
	ModelFrom() Model
	ModelCurrent() Model
}
