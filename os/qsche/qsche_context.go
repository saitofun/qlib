package qsche

import (
	"context"
	"time"

	"github.com/saitofun/qlib/os/qtime"
)

type Context struct {
	stat                   // stat commit,committed,execute,finished
	deadline *qtime.Time   // deadline context deadline
	resc     chan *Result  // resc result chan
	res      *Result       // res result
	done     chan struct{} // done context done chan
	Job
}

func NewContext(j Job) *Context {
	ret := &Context{
		Job:  j,
		resc: make(chan *Result, 1),
		res:  &Result{},
		done: make(chan struct{}, 1),
	}
	ret.stat[0] = qtime.NewTime()
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
		c.res = &Result{}
		c.res.Val, c.res.error = c.Job.Do()
		c.stat[2] = qtime.NewTime()
		c.resc <- c.res
	}
	c.done <- struct{}{}
}

func (c *Context) WithDeadline(deadline time.Time) *Context {
	c.deadline = &qtime.Time{Time: deadline}
	return c
}

func (c *Context) WithTimeout(timeout time.Duration) *Context {
	c.deadline = &qtime.Time{Time: c.stat[0].Add(timeout)}
	return c
}

func (c *Context) Result() (interface{}, error) {
	r := <-c.resc
	return r.Val, r.error
}

// stat 0 commit 1 scheduled 3 done
type stat [3]*qtime.Time

// Latency sub committed and commit
func (s *stat) Latency() time.Duration {
	if s[0] != nil && s[1] != nil {
		return s[1].Time.Sub(s[0].Time)
	}
	return -1
}

// Cost sub done and scheduled
func (s *stat) Cost() time.Duration {
	if s[1] != nil && s[2] != nil {
		return s[2].Time.Sub(s[1].Time)

	}
	return -1
}

// Total sub commit and
func (s *stat) Total() time.Duration {
	if s[0] != nil && s[2] != nil {
		return s[2].Time.Sub(s[0].Time)

	}
	return -1
}

func (s *stat) Stat() [2]int64 {
	return [2]int64{s.Latency().Milliseconds(), s.Cost().Milliseconds()}
}
