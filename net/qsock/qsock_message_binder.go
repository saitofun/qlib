package qsock

import (
	"sync"
	"time"

	"git.querycap.com/aisys/lib/net/qsock/qmsg"
)

type Binder struct {
	*sync.Mutex
	mapping map[qmsg.ID]chan qmsg.Message
}

func NewBinder() *Binder {
	return &Binder{
		Mutex:   &sync.Mutex{},
		mapping: make(map[qmsg.ID]chan qmsg.Message),
	}
}

func (b *Binder) New(id qmsg.ID) {
	b.Lock()
	defer b.Unlock()

	b.mapping[id] = make(chan qmsg.Message, 1)
}

func (b *Binder) get(id qmsg.ID) <-chan qmsg.Message {
	b.Lock()
	defer b.Unlock()

	if c, ok := b.mapping[id]; ok {
		return c
	}
	return nil
}

func (b *Binder) del(id qmsg.ID) {
	b.Lock()
	defer b.Unlock()

	if _, ok := b.mapping[id]; ok {
		delete(b.mapping, id)
	}
}

func (b *Binder) Push(msg qmsg.Message) bool {
	b.Lock()
	defer b.Unlock()
	if c, ok := b.mapping[msg.ID()]; ok && c != nil {
		c <- msg
		return true
	}
	return false
}

func (b *Binder) Wait(id qmsg.ID, d time.Duration) (qmsg.Message, error) {
	defer b.del(id)

	c := b.get(id)
	if c == nil {
		return nil, EMessageUnbound
	}

	select {
	case ret := <-c:
		return ret, nil
	case <-time.After(d):
		return nil, EMessageTimeout
	}
}

func (b *Binder) Remove(id qmsg.ID) {
	b.del(id)
}

func (b *Binder) Reset() {
	b.Lock()
	defer b.Unlock()
	for k := range b.mapping {
		delete(b.mapping, k)
	}
}

func (b *Binder) Len() int {
	b.Lock()
	defer b.Unlock()
	return len(b.mapping)
}
