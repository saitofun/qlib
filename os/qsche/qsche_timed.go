package qsche

import (
	"context"
	"time"
)

type timed struct {
	du     time.Duration
	fn     func()
	ctx    context.Context
	cancel context.CancelFunc
}

func NewTimedScheduler(fn func(), du time.Duration) Scheduler {
	ctx, cancel := context.WithCancel(context.Background())
	return &timed{
		du:     du,
		fn:     fn,
		ctx:    ctx,
		cancel: cancel,
	}
}

func RunTimedScheduler(fn func(), du time.Duration) Scheduler {
	ret := NewTimedScheduler(fn, du)
	go ret.Run()
	return ret
}

func (t *timed) Run() {
	for {
		select {
		case <-t.ctx.Done():
			return
		case <-time.After(t.du):
			go t.fn()
		}
	}
}

func (t *timed) Stop() {
	t.cancel()
}
