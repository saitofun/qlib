package qtype

import (
	"sync/atomic"

	"git.querycap.com/aisys/lib/encoding/qjson"
)

type Int16 struct {
	int64
}

func NewInt16() *Int16 {
	return &Int16{int64: 0}
}

func (i *Int16) Val() int16 {
	return int16(atomic.LoadInt64(&i.int64))
}

func (i *Int16) CAS(pv, nv int) (swapped bool) {
	return atomic.CompareAndSwapInt64(&i.int64, int64(pv), int64(nv))
}

func (i *Int16) Set(v int16) {
	atomic.StoreInt64(&i.int64, int64(v))
}

func (i *Int16) GetSet(v int16) int16 {
	return int16(atomic.SwapInt64(&i.int64, int64(v)))
}

func (i Int16) MarshalJSON() ([]byte, error) {
	return qjson.Marshal(i.Val())
}

func (i *Int16) UnmarshalJSON(dat []byte) error {
	return qjson.Unmarshal(dat, &i.int64)
}
