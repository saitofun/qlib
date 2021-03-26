package qtype

import (
	"math"
	"reflect"
	"sync/atomic"

	"git.querycap.com/ss/lib/encoding/qjson"
)

type Float64 struct {
	*uint64
}

func NewFloat64() *Float64 {
	return &Float64{uint64: new(uint64)}
}

func (f *Float64) Type() reflect.Type {
	return reflect.TypeOf(float64(0))
}

func (f *Float64) Val() float64 {
	return math.Float64frombits(atomic.LoadUint64(f.uint64))
}

func (f *Float64) CAS(pv, nv float64) (swapped bool) {
	return atomic.CompareAndSwapUint64(
		f.uint64, math.Float64bits(pv), math.Float64bits(nv))
}

func (f *Float64) Set(v float64) {
	atomic.StoreUint64(f.uint64, math.Float64bits(v))
}

func (f *Float64) GetSet(v float64) float64 {
	return math.Float64frombits(atomic.SwapUint64(f.uint64, math.Float64bits(v)))
}

func (f *Float64) MarshalJSON() ([]byte, error) {
	return qjson.Marshal(f.Val())
}

func (f *Float64) UnmarshalJSON(v []byte) error {
	tmp := 0.0
	err := qjson.Unmarshal(v, &tmp)
	if err != nil {
		return err
	}
	f.Set(tmp)
	return nil
}
