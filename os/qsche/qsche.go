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
	Stages   [4]qtime.Time // Stages: commit,committed,execute,finished
	deadline *qtime.Time
	result   chan *Result
	res      *Result
	done     chan struct{}
	Job
}

func NewContext(j Job) *Context {
	ret := &Context{
		Job:    j,
		result: make(chan *Result, 1),
		done:   make(chan struct{}, 1),
	}
	ret.Stages[0] = qtime.Now()
	return ret
}

func (c *Context) Deadline() (deadline time.Time, ok bool) {
	if c.deadline != nil {
		deadline = c.deadline.Time
		ok = true
	}
	return
}

func (c *Context) Value(key interface{}) interface{} { return nil }

func (c *Context) Done() <-chan struct{} { return c.done }

func (c *Context) Err() error {
	if c.res == nil {
		return nil
	}
	return c.res.error
}

func (c *Context) Exec(ctx context.Context) {
	if c.deadline != nil {
		ctx, _ = context.WithDeadline(ctx, c.deadline.Time)
	}
	select {
	case <-ctx.Done():
		c.res.error = ctx.Err()
	default:
		c.Stages[2] = qtime.Now()
		c.res = &Result{}
		c.res.Val, c.res.error = c.Job.Do()
		c.Stages[3] = qtime.Now()
		c.result <- c.res
	}
	c.done <- struct{}{}
}

func (c *Context) WithDeadline(deadline time.Time) *Context {
	c.deadline = &qtime.Time{Time: deadline}
	return c
}

func (c *Context) WithTimeout(timeout time.Duration) *Context {
	c.deadline = &qtime.Time{Time: c.Stages[0].Add(timeout)}
	return c
}

func (c *Context) Result() (interface{}, error) {
	r := <-c.result
	return r.Val, r.error
}

type Scheduler interface {
	Start()
	Run()
	Started() bool
	Stop()
	WithContext(ctx context.Context) Scheduler
}

type WorkersScheduler interface {
	Scheduler
	Add(Job) *Context
	AddWithDeadline(Job, time.Time) *Context
	AddWithTimeout(Job, time.Duration) *Context
}
