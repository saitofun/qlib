package qtype

import (
	"sync/atomic"

	"git.querycap.com/aisys/lib/encoding/qjson"
)

type UInt16 struct {
	int64
}

func NewUInt16() *UInt16 {
	return &UInt16{int64: 0}
}

func (i UInt16) Val() int {
	return int(atomic.LoadInt64(&i.int64))
}

func (i *UInt16) CAS(pv, nv uint16) (swapped bool) {
	return atomic.CompareAndSwapInt64(&i.int64, int64(pv), int64(nv))
}

func (i *UInt16) Set(v int16) {
	atomic.StoreInt64(&i.int64, int64(v))
}

func (i *UInt16) GetSet(v int16) int16 {
	return int16(atomic.SwapInt64(&i.int64, int64(v)))
}

func (i UInt16) MarshalJSON() ([]byte, error) {
	return qjson.Marshal(i.Val())
}

func (i *UInt16) UnmarshalJSON(dat []byte) error {
	tmp := int16(0)
	err := qjson.Unmarshal(dat, &tmp)
	if err != nil {
		return err
	}
	i.Set(tmp)
	return nil
}
