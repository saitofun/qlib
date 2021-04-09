package rwmutex

import "sync"

type RWMutex struct {
	*sync.RWMutex
}

func New() *RWMutex {
	return &RWMutex{}
}

func NewSafe() *RWMutex {
	return &RWMutex{RWMutex: &sync.RWMutex{}}
}

func (mu *RWMutex) IsSafe() bool {
	return mu.RWMutex != nil
}

func (mu *RWMutex) Lock() {
	if mu.RWMutex != nil {
		mu.RWMutex.Lock()
	}
}

func (mu *RWMutex) Unlock() {
	if mu.RWMutex != nil {
		mu.RWMutex.Unlock()
	}
}

func (mu *RWMutex) RLock() {
	if mu.RWMutex != nil {
		mu.RWMutex.RLock()
	}
}

func (mu *RWMutex) RUnlock() {
	if mu.RWMutex != nil {
		mu.RWMutex.RUnlock()
	}
}
