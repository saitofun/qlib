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

type group struct {
	*sync.WaitGroup
}

func (g *group) Do(f func()) {
	go func() {
		g.Add(1)
		f()
		g.Done()
	}()
}
