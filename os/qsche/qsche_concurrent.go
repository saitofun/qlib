package qsche

import (
	"context"
	"sync"

	"git.querycap.com/ss/lib/container/qtype"
)

type concurrent struct {
	*Workers
	started *qtype.Bool
	con     int
	ctx     context.Context
	cancel  context.CancelFunc
}

func NewConScheduler(concurrency int) WorkersScheduler {
	ctx, cancel := context.WithCancel(context.Background())
	return &concurrent{
		Workers: NewWorkers(concurrency),
		started: qtype.NewBool(),
		con:     concurrency,
		ctx:     ctx,
		cancel:  cancel,
	}
}

func RunConScheduler(concurrency int) WorkersScheduler {
	ret := NewConScheduler(concurrency)
	go ret.Run()
	return ret
}

func (c *concurrent) WithContext(ctx context.Context) Scheduler {
	c.ctx, c.cancel = context.WithCancel(ctx)
	return c
}

func (c *concurrent) routine() {
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

func (c *concurrent) run() {
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			if ctx := c.Pop(); ctx != nil {
				go ctx.Exec(c.ctx)
			}
		}
	}
}

func (c *concurrent) Started() bool { return c.started.Val() }

func (c *concurrent) Run() {
	if c.started.CAS(false, true) {
		if c.con == 0 {
			c.run()
		}
		wg := sync.WaitGroup{}
		for i := 0; i < c.con; i++ {
			go func() {
				wg.Add(1)
				c.routine()
			}()
		}
		wg.Wait()
	}
}

func (c *concurrent) Stop() {
	c.cancel()
	c.Workers.Close()
}
