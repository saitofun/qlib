package sync

import "sync"

type Guard struct {
	mu *sync.Mutex
}

func New() *Guard {
	return &Guard{mu: &sync.Mutex{}}
}

func NewWith(mu *sync.Mutex) *Guard {
	return &Guard{mu: mu}
}

func (gl *Guard) Do(f func()) {
	gl.mu.Lock()
	defer gl.mu.Unlock()
	f()
}
