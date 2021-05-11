package qstack

import (
	"git.querycap.com/ss/lib/container/qlist"
	"git.querycap.com/ss/lib/container/qtype"
)

type Stack struct {
	act    chan struct{}
	lst    *qlist.List
	closed *qtype.Bool
}

func New() *Stack {
	return &Stack{
		act:    make(chan struct{}, 8096),
		lst:    qlist.NewSafe(),
		closed: qtype.NewBool(),
	}
}

func (s *Stack) Push(v interface{}) {
	if !s.closed.Val() {
		s.act <- struct{}{}
		s.lst.PushBack(v)
	}
}

func (s *Stack) Pop() (ret interface{}) {
	<-s.act
	ret = s.lst.PopBack()
	return
}

func (s *Stack) Len() int {
	if !s.closed.Val() {
		return s.lst.Len()
	}
	return 0
}

func (s *Stack) Close() {
	s.closed.Set(true)
	close(s.act)
	s.lst.Clear()
}
