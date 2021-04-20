package qtype

import (
	"strconv"
	"sync/atomic"

	"git.querycap.com/ss/lib/encoding/qjson"
)

type Int16 struct {
	int32
}

func NewInt16() *Int16 {
	return &Int16{0}
}

func NewInt16WithVal(v int16) *Int16 {
	return &Int16{int32(v)}
}

func (i *Int16) Clone() *Int16 {
	return NewInt16WithVal(i.Val())
}

func (i *Int16) Val() int16 {
	return int16(atomic.LoadInt32(&i.int32))
}

func (i *Int16) CAS(pv, nv int16) (swapped bool) {
	return atomic.CompareAndSwapInt32(&i.int32, int32(pv), int32(nv))
}

func (i *Int16) Set(v int16) int16 {
	return int16(atomic.SwapInt32(&i.int32, int32(v)))
}

func (i *Int16) Add(delta int16) int16 {
	return int16(atomic.AddInt32(&i.int32, int32(delta)))
}

func (i *Int16) String() string {
	return strconv.Itoa(int(i.Val()))
}

func (i Int16) MarshalJSON() ([]byte, error) {
	return qjson.Marshal(i.Val())
}

func (i *Int16) UnmarshalJSON(dat []byte) error {
	v := int16(0)
	if err := qjson.Unmarshal(dat, &v); err != nil {
		return err
	}
	i.Set(v)
	return nil
}
