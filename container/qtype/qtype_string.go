package qtype

import (
	"sync"
	"sync/atomic"

	"git.querycap.com/ss/lib/encoding"
)

type String struct {
	mu  *sync.Mutex
	str atomic.Value
}

func NewString() *String {
	return &String{
		mu: &sync.Mutex{},
	}
}

func (s *String) Val() string {
	str := s.str.Load()
	if str == nil {
		return ""
	}
	return str.(string)
}

func (s *String) CAS(pv, nv string) (swapped bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	str := s.Val()
	if str == pv {
		s.str.Store(nv)
		return true
	}
	return false
}

func (s *String) GetSet(v string) string {
	s.mu.Lock()
	defer s.mu.Unlock()
	str := s.Val()
	s.Set(v)
	return str
}

func (s *String) Set(v string) {
	s.str.Store(v)
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
