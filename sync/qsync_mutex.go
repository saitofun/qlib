package sync

import (
	"sync"
	
	"git.querycap.com/aisys/lib/container/qtype"
)

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
