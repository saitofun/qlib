package qsche

import (
	"time"

	"git.querycap.com/ss/lib/container/qqueue"
	"git.querycap.com/ss/lib/os/qtime"
)

const defaultPoolLimit = 1024

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

func (p *Workers) AddJob(j Job) *Context {
	ctx := NewContext(j)
	p.q.Push(ctx)
	ctx.Committed = qtime.Now()
	return ctx
}

func (p *Workers) AddWithDeadline(j Job, d time.Duration) *Context {
	ctx := NewContext(j)
	p.q.Push(ctx.WithDeadline(d))
	ctx.Committed = qtime.Now()
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
