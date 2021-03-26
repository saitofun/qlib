package qroutines

import (
	"context"
	"sync/atomic"
)

type Job = func()

type LimitedWorkerPool struct {
	jobs   chan func()
	cnt    *int32
	lmt    int32
	ctx    context.Context
	cancel context.CancelFunc
}

func NewLimitedWorkerPool(lmt int) *LimitedWorkerPool {
	ctx, cancel := context.WithCancel(context.Background())
	ret := &LimitedWorkerPool{
		jobs:   make(chan func(), lmt),
		lmt:    int32(lmt),
		cnt:    new(int32),
		ctx:    ctx,
		cancel: cancel,
	}
	go ret.run()
	return ret
}

func (p *LimitedWorkerPool) Add(jobs ...Job) {
	for _, j := range jobs {
		if j != nil {
			p.jobs <- j
		}
	}
}

func (p *LimitedWorkerPool) run() {
	for {
		select {
		case <-p.ctx.Done():
			return
		default:
		}
		j := <-p.jobs
		if j != nil {
			go func(j Job) {
				atomic.AddInt32(p.cnt, 1)
				defer atomic.AddInt32(p.cnt, -1)
				j()
			}(j)
		}
	}
}

func (p *LimitedWorkerPool) Release() {
	p.cancel()
	close(p.jobs)
}

var (
	closed = int32(1)
)

type Pool struct {
	lmt    int
	cnt    int32
	closed int32
}
