package qtype

import (
	"sync/atomic"

	"git.querycap.com/aisys/lib/encoding/qjson"
)

type UInt struct {
	uint64
}

func NewUInt() *UInt {
	return &UInt{uint64: 0}
}

func (i UInt) Val() uint {
	return uint(atomic.LoadUint64(&i.uint64))
}

func (i *UInt) CAS(pv, nv uint) bool {
	return atomic.CompareAndSwapUint64(&i.uint64, uint64(pv), uint64(nv))
}

func (i *UInt) Set(v uint) {
	atomic.StoreUint64(&i.uint64, uint64(v))
}

func (i *UInt) GetSet(v uint) uint {
	return uint(atomic.SwapUint64(&i.uint64, uint64(v)))
}

func (i UInt) MarshalJSON() ([]byte, error) {
	return qjson.Marshal(i.Val())
}

func (i *UInt) UnmarshalJSON(dat []byte) error {
	tmp := uint(0)
	err := qjson.Unmarshal(dat, &tmp)
	if err != nil {
		return err
	}
	i.Set(tmp)
	return nil
}
