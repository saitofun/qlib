package mutex

import (
	"sync"
)

type Mutex struct {
	*sync.Mutex
}

func New() *Mutex {
	return &Mutex{
		Mutex: nil,
	}
}

func NewSafe() *Mutex {
	return &Mutex{
		Mutex: &sync.Mutex{},
	}
}

func (mu *Mutex) IsSafe() bool {
	return mu.Mutex != nil
}

func (mu *Mutex) Lock() {
	if mu.Mutex != nil {
		mu.Mutex.Lock()
	}
}

func (mu *Mutex) Unlock() {
	if mu.Mutex != nil {
		mu.Mutex.Unlock()
	}
}
