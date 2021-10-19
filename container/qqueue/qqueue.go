package qqueue

type Queue interface {
	Push(v interface{})
	TryPush(v interface{}) bool
	Pop() interface{}
	TryPop() interface{}
	WaitPop() <-chan interface{}
	Len() int
	Close()
	Closed() bool
}
