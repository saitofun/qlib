package qsche

import (
	"context"
	"time"
)

type timed struct {
	du     time.Duration
	job    Job
	ctx    context.Context
	cancel context.CancelFunc
}

func NewTimedScheduler(fn func(), du time.Duration) Scheduler {
	ctx, cancel := context.WithCancel(context.Background())
	return &timed{
		du:     du,
		job:    &Fn{fn},
		ctx:    ctx,
		cancel: cancel,
	}
}

func (t *timed) Run() {
	for {
		select {
		case <-t.ctx.Done():
			return
		case <-time.After(t.du):
			t.job.Do()
		}
	}
}

func (t *timed) Stop() {
	t.cancel()
}
