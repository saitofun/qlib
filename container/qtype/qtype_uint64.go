package qtype

import (
	"sync/atomic"

	"git.querycap.com/ss/lib/encoding/qjson"
)

type UInt64 struct {
	uint64
}

func NewUInt64() *UInt64 {
	return &UInt64{uint64: 0}
}

func (i UInt64) Val() uint64 {
	return atomic.LoadUint64(&i.uint64)
}

func (i *UInt64) CAS(pv, nv uint64) (swapped bool) {
	return atomic.CompareAndSwapUint64(&i.uint64, pv, nv)
}

func (i *UInt64) Set(v uint64) {
	atomic.StoreUint64(&i.uint64, v)
}

func (i *UInt64) GetSet(v uint64) uint64 {
	return atomic.SwapUint64(&i.uint64, v)
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
