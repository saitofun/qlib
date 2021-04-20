package qtype

import (
	"strconv"
	"sync/atomic"

	"git.querycap.com/ss/lib/encoding/qjson"
)

type UInt16 struct {
	uint32
}

func NewUInt16() *UInt16 {
	return &UInt16{0}
}

func NewUInt16WithVal(v uint16) *UInt16 {
	return &UInt16{uint32(v)}
}

func (i *UInt16) Clone() *UInt16 {
	return NewUInt16WithVal(i.Val())
}

func (i UInt16) Val() uint16 {
	return uint16(atomic.LoadUint32(&i.uint32))
}

func (i *UInt16) CAS(pv, nv uint16) (swapped bool) {
	return atomic.CompareAndSwapUint32(&i.uint32, uint32(pv), uint32(nv))
}

func (i *UInt16) Set(v uint16) uint16 {
	return uint16(atomic.SwapUint32(&i.uint32, uint32(v)))
}

func (i *UInt16) Add(delta uint16) uint16 {
	return uint16(atomic.AddUint32(&i.uint32, uint32(delta)))
}

func (i UInt16) MarshalJSON() ([]byte, error) {
	return qjson.Marshal(i.Val())
}

func (i *UInt16) String() string {
	return strconv.FormatUint(uint64(i.Val()), 10)
}

func (i *UInt16) UnmarshalJSON(dat []byte) error {
	tmp := uint16(0)
	err := qjson.Unmarshal(dat, &tmp)
	if err != nil {
		return err
	}
	i.Set(tmp)
	return nil
}
