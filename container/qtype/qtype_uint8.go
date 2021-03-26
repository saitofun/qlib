package qtype

import (
	"sync/atomic"

	"git.querycap.com/ss/lib/encoding/qjson"
)

type UInt8 struct {
	uint32
}

func NewUInt8() *UInt8 {
	return &UInt8{uint32: 0}
}

func (i UInt8) Val() uint8 {
	return uint8(atomic.LoadUint32(&i.uint32))
}

func (i *UInt8) CAS(pv, nv uint8) (swapped bool) {
	return atomic.CompareAndSwapUint32(&i.uint32, uint32(pv), uint32(nv))
}

func (i *UInt8) Set(v uint8) {
	atomic.StoreUint32(&i.uint32, uint32(v))
}

func (i *UInt8) GetSet(v uint8) uint8 {
	return uint8(atomic.SwapUint32(&i.uint32, uint32(v)))
}

func (i UInt8) MarshalJSON() ([]byte, error) {
	return qjson.Marshal(i.Val())
}

func (i *UInt8) UnmarshalJSON(dat []byte) error {
	tmp := uint8(0)
	err := qjson.Unmarshal(dat, &tmp)
	if err != nil {
		return err
	}
	i.Set(tmp)
	return nil
}
