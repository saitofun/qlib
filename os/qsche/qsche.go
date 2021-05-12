package qsche

import (
	"context"
	"time"

	"git.querycap.com/ss/lib/os/qtime"
)

type Job interface {
	Do() (interface{}, error)
}

type Fn struct {
	fn func()
}

func NewFn(fn func()) *Fn { return &Fn{fn} }

func (f Fn) Do() (interface{}, error) {
	f.fn()
	return nil, nil
}

type FnWithErr struct {
	fn func() error
}

func NewFnWithErr(fn func() error) *FnWithErr { return &FnWithErr{fn} }

func (f FnWithErr) Do() (interface{}, error) {
	return nil, f.fn()
}

type FnWithVal struct {
	fn func() interface{}
}

func NewFnWithVal(fn func() interface{}) *FnWithVal { return &FnWithVal{fn} }

func (f FnWithVal) Do() (interface{}, error) {
	return f.fn(), nil
}

type FnWithResult struct {
	fn func() (interface{}, error)
}

func NewFnWithResult(fn func() (interface{}, error)) *FnWithResult { return &FnWithResult{fn} }

func (f FnWithResult) Do() (interface{}, error) {
	return f.fn()
}

type Result struct {
	Val interface{}
	error
}

type Context struct {
	CommitAt  qtime.Time
	Committed qtime.Time
	Scheduled qtime.Time
	Finished  qtime.Time
	Deadline  *qtime.Time
	result    chan *Result
	Job
}

func NewContext(j Job) *Context {
	return &Context{
		CommitAt: qtime.Now(),
		Job:      j,
		result:   make(chan *Result, 1),
	}
}

func (c *Context) WithDeadline(d time.Duration) *Context {
	if d > 0 {
		if c.CommitAt.IsZero() {
			c.CommitAt = qtime.Now()
		}
		c.Deadline = &qtime.Time{Time: c.Committed.Add(d)}
	}
	return c
}

func (c *Context) Exec(ctx context.Context) {
	res := &Result{}
	if c.Deadline != nil {
		ctx, _ = context.WithDeadline(ctx, c.Deadline.Time)
	}
	select {
	case <-ctx.Done():
		res.error = ctx.Err()
	default:
		c.Scheduled = qtime.Now()
		res.Val, res.error = c.Job.Do()
		c.Finished = qtime.Now()
		c.result <- res
	}
}

func (c *Context) Result() (interface{}, error) {
	r := <-c.result
	return r.Val, r.error
}

type Scheduler interface {
	Run()
	Stop()
}

type WorkersScheduler interface {
	Scheduler
	Add(Job) *Context
	AddWithDeadline(Job, time.Duration) *Context
}
