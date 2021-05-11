package qsche

import "context"

type chain struct {
	*Workers
	ctx    context.Context
	cancel context.CancelFunc
}

func NewChainedScheduler(lmt ...int) Scheduler {
	ctx, cancel := context.WithCancel(context.Background())
	return &chain{
		Workers: NewWorkers(lmt...),
		ctx:     ctx,
		cancel:  cancel,
	}
}

func (c *chain) Run() {
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			if ctx := c.Pop(); ctx != nil {
				ctx.Do(c.ctx)
			}
		}
	}
}

func (c *chain) Stop() {
	c.cancel()
	c.Workers.Close()
}
