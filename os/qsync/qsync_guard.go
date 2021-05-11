package qsync

import "sync"

type guard struct {
	*sync.Mutex
}

func (g *guard) Do(f func()) {
	g.Lock()
	defer g.Unlock()
	f()
}

func Guard(mutex *sync.Mutex) *guard {
	return &guard{mutex}
}
