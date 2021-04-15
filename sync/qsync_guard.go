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

func (g *Guard) Do(f func()) {
	g.mu.Lock()
	defer g.mu.Unlock()
	f()
}
