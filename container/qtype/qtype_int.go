package qtype

import (
	"sync/atomic"

	"git.querycap.com/aisys/lib/encoding/qjson"
)

type Int struct {
	int64
}

func NewInt() *Int {
	return &Int{int64: 0}
}

func (i *Int) Val() int {
	return int(atomic.LoadInt64(&i.int64))
}

func (i *Int) CAS(pv, nv int) (swapped bool) {
	return atomic.CompareAndSwapInt64(&i.int64, int64(pv), int64(nv))
}

func (i *Int) Set(v int) {
	atomic.StoreInt64(&i.int64, int64(v))
}

func (i *Int) GetSet(v int) int {
	return int(atomic.SwapInt64(&i.int64, int64(v)))
}

func (i Int) MarshalJSON() ([]byte, error) {
	return qjson.Marshal(i.Val())
}

func (i *Int) UnmarshalJSON(dat []byte) error {
	return qjson.Unmarshal(dat, &i.int64)
}
