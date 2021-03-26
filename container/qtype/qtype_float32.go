package qtype

import (
	"math"
	"reflect"
	"sync/atomic"

	"git.querycap.com/aisys/lib/encoding/qjson"
)

type Float32 struct {
	*uint32
}

func NewFloat32() *Float32 {
	return &Float32{uint32: new(uint32)}
}

func (f *Float32) Type() reflect.Type {
	return reflect.TypeOf(float32(0))
}

func (f *Float32) Val() float32 {
	return math.Float32frombits(atomic.LoadUint32(f.uint32))
}

func (f *Float32) CAS(pv, nv float32) (swapped bool) {
	return atomic.CompareAndSwapUint32(
		f.uint32, math.Float32bits(pv), math.Float32bits(nv))
}

func (f *Float32) Set(v float32) {
	atomic.StoreUint32(f.uint32, math.Float32bits(v))
}

func (f *Float32) GetSet(v float32) float32 {
	return math.Float32frombits(atomic.SwapUint32(f.uint32, math.Float32bits(v)))
}

func (f Float32) MarshalJSON() ([]byte, error) {
	return qjson.Marshal(f.Val())
}

func (f *Float32) UnmarshalJSON(v []byte) error {
	tmp := float32(0.0)
	err := qjson.Unmarshal(v, tmp)
	if err != nil {
		return err
	}
	f.Set(tmp)
	return nil
}
