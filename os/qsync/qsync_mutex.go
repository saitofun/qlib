package qsync

import (
	"sync"

	"github.com/saitofun/qlib/container/qtype"
)

// @TODO

type Mutex struct {
	qtype.Int32
	sync.Mutex
}

func (m *Mutex) TryLock() bool {
	if m.CAS(0, 1) {
		m.Lock()
		return true
	}
	return false
}

func (m *Mutex) Lock() {
}

func (m *Mutex) Unlock() {
}
