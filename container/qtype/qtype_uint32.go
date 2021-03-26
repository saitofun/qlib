package qtype

import (
	"sync/atomic"

	"git.querycap.com/aisys/lib/encoding/qjson"
)

type UInt32 struct {
	uint32
}

func NewUInt32() *UInt32 {
	return &UInt32{uint32: 0}
}

func (i UInt32) Val() uint32 {
	return atomic.LoadUint32(&i.uint32)
}

func (i *UInt32) CAS(pv, nv uint32) (swapped bool) {
	return atomic.CompareAndSwapUint32(&i.uint32, pv, nv)
}

func (i *UInt32) Set(v uint32) {
	atomic.StoreUint32(&i.uint32, v)
}

func (i *UInt32) GetSet(v uint32) uint32 {
	return atomic.SwapUint32(&i.uint32, v)
}

func (i UInt32) MarshalJSON() ([]byte, error) {
	return qjson.Marshal(i.Val())
}

func (i *UInt32) UnmarshalJSON(dat []byte) error {
	tmp := uint32(0)
	err := qjson.Unmarshal(dat, &tmp)
	if err != nil {
		return err
	}
	i.Set(tmp)
	return nil
}
