package qsock

import (
	"sync"

	"git.querycap.com/ss/lib/net/qsock/qmsg"
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

func (r *Routes) GetHandlers(t qmsg.Type) []Handler {
	r.mu.Lock()
	defer r.mu.Unlock()
	ret, ok := r.v[t.String()]
	if ok {
		return ret
	}
	return nil
}

func (r *Routes) GetJobs(ev *Event) []Job {
	var handlers = r.GetHandlers(ev.Payload().Type())
	if len(handlers) == 0 {
		return nil
	}
	var ret []func()
	for _, h := range handlers {
		ret = append(ret, func() { h(ev) })
	}
	return ret
}
