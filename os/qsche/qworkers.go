package qsche

import (
	"time"

	"git.querycap.com/ss/lib/container/qqueue"
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

func (p *Workers) Add(j Job) (ctx *Context) {
	defer func() { ctx.Committed() }()
	ctx = NewContext(j)
	p.q.Push(ctx)
	return ctx
}

func (p *Workers) AddWithDeadline(j Job, deadline time.Time) (ctx *Context) {
	defer func() { ctx.Committed() }()
	ctx = NewContext(j)
	p.q.Push(ctx.WithDeadline(deadline))
	return ctx
}

func (p *Workers) AddWithTimeout(j Job, timeout time.Duration) (ctx *Context) {
	defer func() { ctx.Committed() }()
	ctx = NewContext(j)
	p.q.Push(ctx.WithTimeout(timeout))
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
