package qsche

import "context"

type chain struct {
	*Workers
	ctx    context.Context
	cancel context.CancelFunc
}

func NewChainedScheduler(lmt ...int) WorkersScheduler {
	ctx, cancel := context.WithCancel(context.Background())
	return &chain{
		Workers: NewWorkers(lmt...),
		ctx:     ctx,
		cancel:  cancel,
	}
}

func RunChainedScheduler(lmt ...int) WorkersScheduler {
	ret := NewChainedScheduler(lmt...)
	go ret.Run()
	return ret
}

func (c *chain) Run() {
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

func (c *chain) Stop() {
	c.cancel()
	c.Workers.Close()
}
