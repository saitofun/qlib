package qtype

import (
	"strconv"
	"sync/atomic"

	"github.com/saitofun/qlib/encoding/qjson"
)

type UInt64 struct {
	uint64
}

func NewUInt64() *UInt64 {
	return &UInt64{0}
}

func NewUInt64WithVal(v uint64) *UInt64 {
	return &UInt64{v}
}

func (i *UInt64) Clone() *UInt64 {
	return NewUInt64WithVal(i.Val())
}

func (i UInt64) Val() uint64 {
	return atomic.LoadUint64(&i.uint64)
}

func (i *UInt64) CAS(pv, nv uint64) (swapped bool) {
	return atomic.CompareAndSwapUint64(&i.uint64, pv, nv)
}

func (i *UInt64) Set(v uint64) uint64 {
	return atomic.SwapUint64(&i.uint64, v)
}

func (i *UInt64) Add(delta uint64) uint64 {
	return atomic.AddUint64(&i.uint64, delta)
}

func (i *UInt64) String() string {
	return strconv.FormatUint(uint64(i.Val()), 10)
}

func (i UInt64) MarshalJSON() ([]byte, error) {
	return qjson.Marshal(i.Val())
}

func (i *UInt64) UnmarshalJSON(dat []byte) error {
	tmp := uint64(0)
	err := qjson.Unmarshal(dat, &tmp)
	if err != nil {
		return err
	}
	i.Set(tmp)
	return nil
}
