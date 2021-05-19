package qsche

import (
	"context"
	"time"

	"git.querycap.com/ss/lib/container/qtype"
)

type timed struct {
	started *qtype.Bool
	du      time.Duration
	fn      func()
	ctx     context.Context
	cancel  context.CancelFunc
}

func NewTimedScheduler(fn func(), du time.Duration) Scheduler {
	ctx, cancel := context.WithCancel(context.Background())
	return &timed{
		started: qtype.NewBool(),
		du:      du,
		fn:      fn,
		ctx:     ctx,
		cancel:  cancel,
	}
}

func RunTimedScheduler(fn func(), du time.Duration) Scheduler {
	ret := NewTimedScheduler(fn, du)
	ret.Run()
	return ret
}

func (t *timed) WithContext(ctx context.Context) Scheduler {
	t.ctx, t.cancel = context.WithCancel(ctx)
	return t
}

func (t *timed) Started() bool { return t.started.Val() }

func (t *timed) Start() {
	if t.started.CAS(false, true) {
		for {
			select {
			case <-t.ctx.Done():
				return
			case <-time.After(t.du):
				go t.fn()
			}
		}
	}
}

func (t *timed) Run() { go t.Start() }

func (t *timed) Stop() {
	t.cancel()
}
