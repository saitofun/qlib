package qtype

import (
	"strconv"
	"sync/atomic"

	"git.querycap.com/ss/lib/encoding/qjson"
)

type UInt struct {
	uint64
}

func NewUInt() *UInt {
	return &UInt{0}
}

func NewUIntWithVal(v uint) *UInt {
	return &UInt{uint64(v)}
}

func (i *UInt) Clone() *UInt {
	return NewUIntWithVal(i.Val())
}

func (i *UInt) Val() uint {
	return uint(atomic.LoadUint64(&i.uint64))
}

func (i *UInt) CAS(pv, nv uint) bool {
	return atomic.CompareAndSwapUint64(&i.uint64, uint64(pv), uint64(nv))
}

func (i *UInt) Set(v uint) uint {
	return uint(atomic.SwapUint64(&i.uint64, uint64(v)))
}

func (i *UInt) Add(delta uint) uint {
	return uint(atomic.AddUint64(&i.uint64, uint64(delta)))
}

func (i *UInt) String() string {
	return strconv.FormatUint(uint64(i.Val()), 10)
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
