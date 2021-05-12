package qsche

import (
	"context"
	"sync"
)

type concurrent struct {
	*Workers
	con    int
	ctx    context.Context
	cancel context.CancelFunc
}

func NewConScheduler(con int, lmt ...int) WorkersScheduler {
	ctx, cancel := context.WithCancel(context.Background())
	return &concurrent{
		Workers: NewWorkers(lmt...),
		con:     con,
		ctx:     ctx,
		cancel:  cancel,
	}
}

func RunConScheduler(con int, lmt ...int) WorkersScheduler {
	ret := NewConScheduler(con, lmt...)
	go ret.Run()
	return ret
}

func (c *concurrent) run() {
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

func (c *concurrent) Run() {
	wg := sync.WaitGroup{}
	for i := 0; i < c.con; i++ {
		go func() {
			wg.Add(1)
			c.run()
		}()
	}
	wg.Wait()
}

func (c *concurrent) Stop() {
	c.cancel()
	c.Workers.Close()
}
