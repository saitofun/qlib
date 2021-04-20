package qtype

import (
	"sync/atomic"

	"git.querycap.com/ss/lib/encoding"
)

type String struct {
	s atomic.Value
}

func NewString() *String {
	return NewStringWithVal("")
}

func NewStringWithVal(v string) *String {
	ret := &String{atomic.Value{}}
	ret.s.Store(v)
	return ret
}

func (s *String) Clone() *String {
	return NewStringWithVal(s.Val())
}

func (s *String) Val() string {
	return s.s.Load().(string)
}

func (s *String) Set(v string) string {
	ret := s.Val()
	s.s.Store(v)
	return ret
}

func (s *String) String() string {
	return s.Val()
}

func (s *String) Bytes() []byte {
	return encoding.StrToBytes(s.Val())
}

func (s String) MarshalJSON() ([]byte, error) {
	return encoding.StrToBytes(`"` + s.Val() + `"`), nil
}

func (s *String) UnmarshalJSON(v []byte) error {
	s.Set(encoding.BytesToStr(v))
	return nil
}
