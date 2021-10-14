package qqueue

type Queue interface {
	Push(v interface{})
	TryPush(v interface{}) bool
	Pop() interface{}
	TryPop() interface{}
	Len() int
	Close()
	Closed() bool
}
