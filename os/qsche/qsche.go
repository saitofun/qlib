package qsche

import (
	"context"
	"time"
)

type FnJob struct{ fn Fn }

type Fn func()

func NewFnJob(f Fn) *FnJob { return &FnJob{f} }

func (f FnJob) Do() (interface{}, error) { f.fn(); return nil, nil }

type FnWithErrJob struct{ fn FnWithErr }

type FnWithErr func() error

func NewFnWithErrJob(f FnWithErr) *FnWithErrJob { return &FnWithErrJob{f} }

func (f FnWithErrJob) Do() (interface{}, error) { return nil, f.fn() }

type FnWithValJob struct{ fn FnWithVal }

type FnWithVal func() interface{}

func NewFnWithValJob(f FnWithVal) *FnWithValJob { return &FnWithValJob{f} }

func (f FnWithValJob) Do() (interface{}, error) { return f.fn(), nil }

type FnWithResultJob struct{ fn FnWithResult }

type FnWithResult func() (interface{}, error)

func NewFnWithResultJob(fn func() (interface{}, error)) *FnWithResultJob { return &FnWithResultJob{fn} }

func (f FnWithResultJob) Do() (interface{}, error) { return f.fn() }

// Job sche schedule unit
type Job interface{ Do() (interface{}, error) }

// Result job result
type Result struct {
	Val interface{}
	error
}

// Scheduler scheduler without worker operations
type Scheduler interface {
	Start()
	Run()
	Started() bool
	Stop()
	WithContext(ctx context.Context) Scheduler
}

// WorkersScheduler worker manage and schedule control
type WorkersScheduler interface {
	Scheduler
	Add(Job) *Context
	AddWithDeadline(Job, time.Time) *Context
	AddWithTimeout(Job, time.Duration) *Context
}
