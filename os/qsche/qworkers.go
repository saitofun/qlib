package qsche

import (
	"time"

	"git.querycap.com/ss/lib/container/qqueue"
	"git.querycap.com/ss/lib/os/qtime"
)

const defaultPoolLimit = 64

type Workers struct {
	q qqueue.Queue
}

func NewWorkers(lmt ...int) *Workers {
	limit := defaultPoolLimit
	if len(lmt) > 0 && lmt[0] > 0 {
		limit = lmt[0]
	}
	return &Workers{q: qqueue.NewLimited(limit)}
}

func (p *Workers) Add(j Job) (ctx *Context) {
	ctx = NewContext(j)
	p.q.Push(ctx)
	ctx.Stages[1] = qtime.Now()
	return ctx
}

func (p *Workers) AddWithDeadline(j Job, deadline time.Time) (ctx *Context) {
	ctx = NewContext(j)
	p.q.Push(ctx.WithDeadline(deadline))
	ctx.Stages[1] = qtime.Now()
	return ctx
}

func (p *Workers) AddWithTimeout(j Job, timeout time.Duration) (ctx *Context) {
	ctx = NewContext(j)
	p.q.Push(ctx.WithTimeout(timeout))
	ctx.Stages[1] = qtime.Now()
	return ctx
}

func (p *Workers) Pop() *Context {
	ctx := p.q.Pop()
	if ctx == nil {
		return nil
	}
	return ctx.(*Context)
}

func (p *Workers) Close() { p.q.Close() }

func (p *Workers) Closed() bool { return p.q.Closed() }
