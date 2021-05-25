package qsche

import (
	"context"
	"sync"
	"time"

	"git.querycap.com/ss/lib/os/qsync"
)

type (
	Fn        func()
	FnWithErr func() error
)

// FnJob job: func()
type FnJob struct {
	id string
	fn Fn
}

func NewFnJob(f Fn, id ...string) *FnJob  { return &FnJob{"", f} }
func (f *FnJob) Do() (interface{}, error) { f.fn(); return nil, nil }
func (f *FnJob) WithID(id string) *FnJob  { f.id = id; return f }

// FnWithErrJob job: func() error
type FnWithErrJob struct {
	id string
	fn FnWithErr
}

func NewFnWithErrJob(f FnWithErr, id ...string) *FnWithErrJob {
	return &FnWithErrJob{"", f}
}

func (f *FnWithErrJob) Do() (interface{}, error)       { return nil, f.fn() }
func (f *FnWithErrJob) WithID(id string) *FnWithErrJob { f.id = id; return f }

// Job sche schedule unit
type Job interface{ Do() (interface{}, error) }

type NamedJob interface {
	Job
	Name() string
}

type NamedJobWithStat interface {
	NamedJob
	statisticMarker()
}

type statistic interface{ statisticMarker() }

func AsNamedJobWithStat(j NamedJob) NamedJobWithStat {
	return &struct {
		NamedJob
		statistic
	}{NamedJob: j}
}

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
	Context() context.Context
	WithContext(ctx context.Context) Scheduler
}

// WorkersScheduler worker manage and schedule control
type WorkersScheduler interface {
	Scheduler
	Add(Job) *Context
	AddWithDeadline(Job, time.Time) *Context
	AddWithTimeout(Job, time.Duration) *Context
	WaitGroup(...Job) []*Context
}

func WaitGroup(sche WorkersScheduler, jobs ...Job) (ret []*Context) {
	ch := make(chan *Context, len(jobs))
	wg := &sync.WaitGroup{}
	for i := range jobs {
		select {
		case <-sche.Context().Done():
			return nil
		default:
			qsync.Group(wg).Do(func() {
				ctx := sche.Add(jobs[i])
				ch <- ctx
				<-ctx.Done()
			})
		}
	}
	wg.Wait()
	for i := 0; i < len(jobs); i++ {
		ret = append(ret, <-ch)
	}
	return
}
