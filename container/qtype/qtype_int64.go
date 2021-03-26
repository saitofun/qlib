package qtype

import (
	"sync/atomic"

	"git.querycap.com/aisys/lib/encoding/qjson"
)

type Int64 struct {
	int64
}

func NewInt64() *Int64 {
	return &Int64{int64: 0}
}

func (i *Int64) Val() int64 {
	return atomic.LoadInt64(&i.int64)
}

func (i *Int64) CAS(pv, nv int64) (swapped bool) {
	return atomic.CompareAndSwapInt64(&i.int64, pv, nv)
}

func (i *Int64) Set(v int64) {
	atomic.StoreInt64(&i.int64, v)
}

func (i *Int64) GetSet(v int64) int64 {
	return atomic.SwapInt64(&i.int64, v)
}

func (i Int64) MarshalJSON() ([]byte, error) {
	return qjson.Marshal(i.Val())
}

func (i *Int64) UnmarshalJSON(dat []byte) error {
	return qjson.Unmarshal(dat, &i.int64)
}
