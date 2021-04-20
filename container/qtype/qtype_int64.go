package qtype

import (
	"strconv"
	"sync/atomic"

	"git.querycap.com/ss/lib/encoding/qjson"
)

type Int64 struct {
	int64
}

func NewInt64() *Int64 {
	return &Int64{int64: 0}
}

func NewInt64WithVal(v int64) *Int64 {
	return &Int64{int64: v}
}

func (i *Int64) Clone() *Int64 {
	return NewInt64WithVal(i.Val())
}

func (i *Int64) Val() int64 {
	return atomic.LoadInt64(&i.int64)
}

func (i *Int64) CAS(pv, nv int64) (swapped bool) {
	return atomic.CompareAndSwapInt64(&i.int64, pv, nv)
}

func (i *Int64) Set(v int64) int64 {
	return atomic.SwapInt64(&i.int64, v)
}

func (i *Int64) Add(delta int64) int64 {
	return atomic.AddInt64(&i.int64, delta)
}

func (i *Int64) String() string {
	return strconv.FormatInt(int64(i.Val()), 10)
}

func (i Int64) MarshalJSON() ([]byte, error) {
	return qjson.Marshal(i.Val())
}

func (i *Int64) UnmarshalJSON(dat []byte) error {
	return qjson.Unmarshal(dat, &i.int64)
}
