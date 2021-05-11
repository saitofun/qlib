package qqueue

type Queue interface {
	Push(v interface{})
	Pop() interface{}
	Len() int
	Close()
	Closed() bool
}
