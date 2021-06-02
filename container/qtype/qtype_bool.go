package qtype

import (
	"reflect"
	"sync/atomic"

	"github.com/saitofun/qlib/encoding/qjson"
)

type Bool struct{ int32 }

var (
	fJSON  = []byte("false")
	tJSON  = []byte("true")
	tInt32 = int32(1)
)

func NewBool() *Bool {
	return &Bool{}
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
	return atomic.LoadInt32(&b.int32) == 1
}

func (b *Bool) CAS(pv, nv bool) (swapped bool) {
	var old, new int32
	if pv {
		old = tInt32
	}
	if nv {
		new = tInt32
	}
	return atomic.CompareAndSwapInt32(&b.int32, old, new)
}

func (b *Bool) Set(v bool) (old bool) {
	if v {
		old = atomic.SwapInt32(&b.int32, 1) == 1
	} else {
		old = atomic.SwapInt32(&b.int32, 0) == 1
	}
	return
}

func (b *Bool) Type() reflect.Type {
	return reflect.TypeOf(true)
}

func (b *Bool) MarshalJSON() ([]byte, error) {
	if b.Val() {
		return tJSON, nil
	}
	return fJSON, nil
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
