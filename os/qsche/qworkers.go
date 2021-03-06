package qsche

import (
	"time"

	"github.com/saitofun/qlib/container/qqueue"
	"github.com/saitofun/qlib/container/qtype"
	"github.com/saitofun/qlib/os/qtime"
)

const defaultPoolLimit = 64

type Workers struct {
	q   qqueue.Queue
	lmt int
	seq *qtype.Int64
}

func NewWorkers(lmt ...int) *Workers {
	limit := defaultPoolLimit
	if len(lmt) > 0 && lmt[0] > 0 {
		limit = lmt[0]
	}
	return &Workers{q: qqueue.NewLimited(limit), lmt: limit, seq: qtype.NewInt64()}
}

func (p *Workers) Limit() int { return p.lmt }

func (p *Workers) Add(j Job) (ctx *Context) {
	ctx = NewContext(j, p.seq.Add(1))
	p.q.Push(ctx)
	return ctx
}

func (p *Workers) AddWithDeadline(j Job, deadline time.Time) (ctx *Context) {
	ctx = NewContext(j, p.seq.Add(1))
	p.q.Push(ctx.WithDeadline(deadline))
	return ctx
}

func (p *Workers) AddWithTimeout(j Job, timeout time.Duration) (ctx *Context) {
	ctx = NewContext(j, p.seq.Add(1))
	p.q.Push(ctx.WithTimeout(timeout))
	return ctx
}

func (p *Workers) Pop() *Context {
	ctx := p.q.Pop()
	if ctx == nil {
		return nil
	}
	ctx.(*Context).stat[1] = qtime.NewTime()
	return ctx.(*Context)
}

func (p *Workers) Close() { p.q.Close() }

func (p *Workers) Closed() bool { return p.q.Closed() }

func (p *Workers) Len() int { return p.q.Len() }
