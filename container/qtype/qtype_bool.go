package qtype

import (
	"reflect"
	"sync/atomic"

	"git.querycap.com/aisys/lib/encoding/qjson"
)

type Bool struct {
	*int32
}

var (
	vs = map[bool]int32{true: 1, false: 0}
	f  = []byte("false")
	t  = []byte("true")
)

func NewBool() *Bool {
	return &Bool{int32: new(int32)}
}

func NewBoolWithVal(v bool) *Bool {
	ret := NewBool()
	ret.Set(v)
	return ret
}

func (b *Bool) Clone() *Bool {
	return NewBoolWithVal(b.Val())
}

func (b *Bool) Val() bool {
	return atomic.LoadInt32(b.int32) == 1
}

func (b *Bool) CAS(pv, nv bool) (swapped bool) {
	return atomic.CompareAndSwapInt32(b.int32, vs[pv], vs[nv])
}

func (b *Bool) Set(v bool) {
	atomic.StoreInt32(b.int32, vs[v])
}

func (b *Bool) GetSet(v bool) bool {
	return atomic.SwapInt32(b.int32, vs[v]) == 1
}

func (b *Bool) Type() reflect.Type {
	return reflect.TypeOf(true)
}

func (b *Bool) MarshalJSON() ([]byte, error) {
	if b.Val() {
		return t, nil
	}
	return f, nil
}

func (b *Bool) UnmarshalJSON(v []byte) error {
	val := false
	err := qjson.Unmarshal(v, &val)
	if err != nil {
		return err
	}
	b.Set(val)
	return nil
}
