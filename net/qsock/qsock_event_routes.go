package qsock

import (
	"sync"

	"git.querycap.com/ss/lib/net/qsock/qmsg"
	"git.querycap.com/ss/lib/os/qsche"
)

type Handler func(*Event)

func HandlerFunc(h Handler, ev *Event) qsche.Fn { return func() { h(ev) } }

type Job func()

type Routes struct {
	mu *sync.Mutex
	v  map[qmsg.Type][]Handler
}

func NewRoutes() *Routes {
	return &Routes{
		mu: &sync.Mutex{},
		v:  make(map[qmsg.Type][]Handler),
	}
}

func (r *Routes) Register(t qmsg.Type, fns ...Handler) {
	if r == nil {
		return
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	r.v[t] = append(r.v[t], fns...)
}

func (r *Routes) Handlers(t qmsg.Type) []Handler {
	if r == nil {
		return nil
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.v[t]
}
