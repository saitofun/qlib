package qqueue

import "git.querycap.com/ss/lib/container/qtype"

type limited struct {
	qch    chan interface{}
	closed *qtype.Bool
}

func NewLimited(lmt int) Queue {
	return &limited{
		qch:    make(chan interface{}, lmt),
		closed: qtype.NewBool(),
	}
}

func (q *limited) Push(v interface{}) {
	q.qch <- v
}

func (q *limited) Pop() interface{} {
	return <-q.qch
}

func (q *limited) Len() int {
	if q.closed.Val() {
		return 0
	}
	return len(q.qch)
}

func (q *limited) Close() {
	q.closed.Set(false)
	close(q.qch)
}

func (q *limited) Closed() bool { return q.closed.Val() }
