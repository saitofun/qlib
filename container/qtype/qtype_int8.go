package qtype

import (
	"sync/atomic"

	"git.querycap.com/ss/lib/encoding/qjson"
)

type Int8 struct {
	int32
}

func NewInt8() *Int8 {
	return &Int8{int32: 0}
}

func (i Int8) Val() int8 {
	return int8(atomic.LoadInt32(&i.int32))
}

func (i *Int8) CAS(pv, nv int8) (swapped bool) {
	return atomic.CompareAndSwapInt32(&i.int32, int32(pv), int32(nv))
}

func (i *Int8) Set(v int8) {
	atomic.StoreInt32(&i.int32, int32(v))
}

func (i *Int8) GetSet(v int8) int8 {
	return int8(atomic.SwapInt32(&i.int32, int32(v)))
}

func (i Int8) MarshalJSON() ([]byte, error) {
	return qjson.Marshal(i.Val())
}

func (i *Int8) UnmarshalJSON(dat []byte) error {
	return qjson.Unmarshal(dat, &i.int32)
}
