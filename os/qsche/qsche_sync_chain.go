package qsche

import (
	"context"

	"github.com/saitofun/qlib/container/qtype"
)

type chain struct {
	*Workers
	id      string
	started *qtype.Bool
	ctx     context.Context
	cancel  context.CancelFunc
}

func NewChainedScheduler(id string, lmt ...int) WorkersScheduler {
	ctx, cancel := context.WithCancel(context.Background())
	return &chain{
		Workers: NewWorkers(lmt...),
		id:      id,
		started: qtype.NewBool(),
		ctx:     ctx,
		cancel:  cancel,
	}
}

func RunChainedScheduler(id string, lmt ...int) WorkersScheduler {
	ret := NewChainedScheduler(id, lmt...)
	ret.Run()
	return ret
}

func (c *chain) Context() context.Context { return c.ctx }

func (c *chain) WithContext(ctx context.Context) Scheduler {
	c.ctx, c.cancel = context.WithCancel(ctx)
	return c
}

func (c *chain) Start() {
	if c.started.CAS(false, true) {
		for {
			select {
			case <-c.ctx.Done():
				return
			default:
				if ctx := c.Pop(); ctx != nil {
					ctx.Exec(c.ctx)
				}
			}
		}
	}
}

func (c *chain) Run() { go c.Start() }

func (c *chain) Started() bool { return c.started.Val() }

func (c *chain) Stop() {
	c.cancel()
	c.Workers.Close()
}

func (c *chain) WaitGroup(jobs ...Job) []*Context {
	return WaitGroup(c, jobs...)
}
