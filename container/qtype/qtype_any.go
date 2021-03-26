package qtype

import (
	"sync"
	"sync/atomic"

	"git.querycap.com/ss/lib/encoding/qjson"
)

type Any struct {
	mu  *sync.Mutex
	val atomic.Value
}

func NewAny() *Any {
	return &Any{
		mu: &sync.Mutex{},
	}
}

func (i *Any) Val() interface{} {
	return i.val.Load()
}

func (i *Any) Set(v interface{}) {
	i.val.Store(v)
}

func (i *Any) GetSet(v interface{}) interface{} {
	i.mu.Lock()
	defer i.mu.Unlock()
	ret := i.Val()
	i.Set(v)
	return ret
}

func (i *Any) CAS(pv, nv interface{}) (swapped bool) {
	i.mu.Lock()
	defer i.mu.Unlock()
	if pv == i.Val() {
		i.Set(nv)
		return true
	}
	return false
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
