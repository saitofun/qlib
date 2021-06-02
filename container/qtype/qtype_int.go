package qtype

import (
	"strconv"
	"sync/atomic"

	"github.com/saitofun/qlib/encoding/qjson"
)

type Int struct {
	int64
}

func NewInt() *Int {
	return &Int{int64: 0}
}

func NewIntWithVal(v int) *Int {
	return &Int{int64: int64(v)}
}

func (i *Int) Clone() *Int {
	return NewIntWithVal(i.Val())
}

func (i *Int) Val() int {
	return int(atomic.LoadInt64(&i.int64))
}

func (i *Int) CAS(pv, nv int) (swapped bool) {
	return atomic.CompareAndSwapInt64(&i.int64, int64(pv), int64(nv))
}

func (i *Int) Set(v int) int {
	return int(atomic.SwapInt64(&i.int64, int64(v)))
}

func (i *Int) String() string {
	return strconv.Itoa(i.Val())
}

func (i *Int) Add(delta int) int {
	return int(atomic.AddInt64(&i.int64, int64(delta)))
}

func (i Int) MarshalJSON() ([]byte, error) {
	return qjson.Marshal(i.Val())
}

func (i *Int) UnmarshalJSON(dat []byte) error {
	return qjson.Unmarshal(dat, &i.int64)
}
