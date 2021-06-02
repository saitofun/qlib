package qqueue

import (
	"math"

	"github.com/saitofun/qlib/container/qlist"
	"github.com/saitofun/qlib/container/qtype"
)

type unlimited struct {
	lst    *qlist.List
	act    chan struct{}
	qch    chan interface{}
	closed *qtype.Bool
}

func New() Queue {
	ret := &unlimited{
		lst:    qlist.NewSafe(),
		act:    make(chan struct{}, math.MaxInt32),
		qch:    make(chan interface{}, gCap),
		closed: qtype.NewBool(),
	}
	go ret.sync()
	return ret
}

const (
	gCap   = 4096
	gBatch = 32
)

func (q *unlimited) Push(v interface{}) {
	if !q.closed.Val() {
		q.lst.PushBack(v)
		q.act <- struct{}{}
	}
}

func (q *unlimited) Pop() interface{} {
	return <-q.qch
}

func (q *unlimited) Len() int {
	return len(q.qch) + len(q.act)
}

func (q *unlimited) Close() {
	q.closed.Set(true)
	close(q.act)
	close(q.qch)
	q.lst.Clear()
}

func (q *unlimited) sync() {
	defer func() {
		if q.closed.Val() {
			_ = recover()
		}
	}()
	for !q.closed.Val() {
		<-q.act
		if !q.closed.Val() {
			bat := q.lst.PopFrontN(gBatch)
			for _, v := range bat {
				q.qch <- v
			}
			for i := 0; i < len(bat); i++ {
				<-q.act
			}
		} else {
			break
		}
	}
}

func (q *unlimited) Closed() bool { return q.closed.Val() }
