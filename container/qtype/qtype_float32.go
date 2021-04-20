package qtype

import (
	"math"
	"strconv"
	"sync/atomic"

	"git.querycap.com/ss/lib/encoding/qjson"
)

type Float32 struct {
	uint32
}

func NewFloat32() *Float32 {
	return &Float32{0}
}

func NewFloat32WithVal(v float32) *Float32 {
	return &Float32{math.Float32bits(v)}
}

func (f *Float32) Clone() *Float32 {
	return NewFloat32WithVal(f.Val())
}

func (f *Float32) Val() float32 {
	return math.Float32frombits(atomic.LoadUint32(&f.uint32))
}

func (f *Float32) CAS(pv, nv float32) (swapped bool) {
	return atomic.CompareAndSwapUint32(
		&f.uint32, math.Float32bits(pv), math.Float32bits(nv))
}

func (f *Float32) Set(v float32) float32 {
	return math.Float32frombits(atomic.SwapUint32(&f.uint32, math.Float32bits(v)))
}

func (f *Float32) Add(delta float32) float32 {
	var ret float32
	for {
		old := math.Float32frombits(f.uint32)
		ret = old + delta
		if atomic.CompareAndSwapUint32(&f.uint32, math.Float32bits(old), math.Float32bits(ret)) {
			break
		}
	}
	return ret
}

func (f *Float32) String() string {
	return strconv.FormatFloat(float64(f.Val()), 'f', -1, 32)
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
