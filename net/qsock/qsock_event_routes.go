package qsock

import (
	"sync"

	"git.querycap.com/ss/lib/net/qsock/qmsg"
	"git.querycap.com/ss/lib/os/qsche"
	"git.querycap.com/ss/lib/os/qsync"
)

type (
	Handler = func(*Event)
	Job     = func()
)

type Routes struct {
	mu *sync.Mutex
	v  map[string][]Handler
}

func NewRoutes() *Routes {
	return &Routes{
		mu: &sync.Mutex{},
		v:  make(map[string][]Handler),
	}
}

func (r *Routes) Register(t qmsg.Type, fns ...Handler) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.v[t.String()] = append(r.v[t.String()], fns...)
}

func (r *Routes) EventJobs(ev *Event) (ret []qsche.Job) {
	var handlers []Handler
	qsync.Guard(r.mu).Do(func() {
		handlers = r.v[ev.Payload().Type().String()]
	})
	if len(handlers) == 0 {
		return nil
	}
	for _, h := range handlers {
		ret = append(ret, qsche.NewFn(func() { h(ev) }))
	}
	return ret
}
