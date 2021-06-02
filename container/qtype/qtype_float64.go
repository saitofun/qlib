package qtype

import (
	"math"
	"strconv"
	"sync/atomic"

	"github.com/saitofun/qlib/encoding/qjson"
)

type Float64 struct {
	uint64
}

func NewFloat64() *Float64 {
	return &Float64{0}
}

func NewFloat64WithVal(v float64) *Float64 {
	return &Float64{math.Float64bits(v)}
}

func (f *Float64) Clone() *Float64 {
	return NewFloat64WithVal(f.Val())
}

func (f *Float64) Val() float64 {
	return math.Float64frombits(atomic.LoadUint64(&f.uint64))
}

func (f *Float64) CAS(pv, nv float64) (swapped bool) {
	return atomic.CompareAndSwapUint64(
		&f.uint64, math.Float64bits(pv), math.Float64bits(nv))
}

func (f *Float64) Set(v float64) float64 {
	return math.Float64frombits(atomic.SwapUint64(&f.uint64, math.Float64bits(v)))
}

func (f *Float64) Add(delta float64) float64 {
	var new float64
	for {
		old := math.Float64frombits(f.uint64)
		new = old + delta
		if atomic.CompareAndSwapUint64(&f.uint64, math.Float64bits(old), math.Float64bits(new)) {
			break
		}
	}
	return new
}

func (f *Float64) String() string {
	return strconv.FormatFloat(f.Val(), 'f', -1, 64)
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
