package qtype

import (
	"sync/atomic"

	"git.querycap.com/ss/lib/encoding/qconv"
	"git.querycap.com/ss/lib/encoding/qjson"
)

type Any struct {
	v atomic.Value
}

func New() *Any {
	return &Any{}
}

func NewWithVal(v interface{}) *Any {
	ret := &Any{}
	ret.v.Store(v)
	return ret
}

func (i *Any) Clone() *Any {
	return NewWithVal(i.Val())
}

func (i *Any) Val() interface{} {
	return i.v.Load()
}

func (i *Any) Set(v interface{}) interface{} {
	ret := i.v.Load()
	i.v.Store(v)
	return ret
}

func (i *Any) String() string {
	return qconv.String(i.Val())
}

func (i Any) MarshalJSON() ([]byte, error) {
	return qjson.Marshal(i.Val())
}

func (i *Any) UnmarshalJSON(v []byte) error {
	val := (interface{})(nil)
	err := qjson.Unmarshal(v, &val)
	if err != nil {
		return err
	}
	i.Set(val)
	return nil
}
