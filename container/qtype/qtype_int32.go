package qtype

import (
	"sync/atomic"

	"git.querycap.com/aisys/lib/encoding/qjson"
)

type Int32 struct {
	int32
}

func NewInt32() *Int32 {
	return &Int32{int32: 0}
}

func (i *Int32) Val() int32 {
	return atomic.LoadInt32(&i.int32)
}

func (i *Int32) CAS(pv, nv int32) (swapped bool) {
	return atomic.CompareAndSwapInt32(&i.int32, pv, nv)
}

func (i *Int32) Set(v int32) {
	atomic.StoreInt32(&i.int32, v)
}

func (i *Int32) GetSet(v int32) int32 {
	return atomic.SwapInt32(&i.int32, v)
}

func (i Int32) MarshalJSON() ([]byte, error) {
	return qjson.Marshal(i.Val())
}

func (i *Int32) UnmarshalJSON(dat []byte) error {
	return qjson.Unmarshal(dat, &i.int32)
}
