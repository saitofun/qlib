package qtime

import (
	"math"
	"time"
)

type ElapsedAccumulator struct {
	initial time.Time
	peek    time.Duration
	minimum time.Duration
	last    time.Duration
	total   time.Duration
	average time.Duration

	op int64
	do int64
}

func NewElapsedAccumulator() *ElapsedAccumulator {
	return &ElapsedAccumulator{
		initial: time.Now(),
		peek:    0,
		minimum: math.MaxInt64,
	}
}

func (v *ElapsedAccumulator) Do(f func(), op int64) {
	span := FuncElapsed(f)

	if span > v.peek {
		v.peek = span
	} else if v.minimum > span {
		v.minimum = span
	}
	v.last = span
	v.total += span
	v.op += op
	v.do += 1
}

func (v *ElapsedAccumulator) Last() time.Duration {
	return v.last
}

func (v *ElapsedAccumulator) Peek() time.Duration {
	return v.peek
}

func (v *ElapsedAccumulator) Minimum() time.Duration {
	return v.minimum
}

func (v *ElapsedAccumulator) Total() time.Duration {
	return v.total
}

func (v *ElapsedAccumulator) DoCount() int64 {
	return v.do
}

func (v *ElapsedAccumulator) OpCount() int64 {
	return v.op
}

func (v *ElapsedAccumulator) AverageOp() time.Duration {
	if v.op == 0 {
		return 0
	}
	return v.total / time.Duration(v.op)
}

func (v *ElapsedAccumulator) AverageDo() time.Duration {
	if v.do == 0 {
		return 0
	}
	return v.total / time.Duration(v.do)
}

func (v *ElapsedAccumulator) SinceInitial() time.Duration {
	return time.Since(v.initial)
}

func FuncElapsed(f func()) time.Duration {
	t := time.Now()
	f()
	return time.Since(t)
}
