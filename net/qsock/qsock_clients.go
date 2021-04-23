package qsock

import (
	"sync"
)

type clients struct {
	*sync.RWMutex
	connections map[string]*Node
	events      chan *Event
}

func NewClientManager() *clients {
	return &clients{
		RWMutex:     &sync.RWMutex{},
		connections: make(map[string]*Node),
		events:      make(chan *Event, 2048),
	}
}

func (cs *clients) New(n *Node) {
	cs.Lock()
	defer cs.Unlock()

	cs.connections[n.ID()] = n

	go func(n *Node) {
		defer cs.Remove(n.ID())
		for {
			if n.closed {
				return
			}
			msg, _ := n.ReadMessage(0)
			if msg == nil {
				break
			}
			cs.events <- &Event{n, msg}
		}
	}(n)
}

func (cs *clients) Remove(id string, reason ...interface{}) {
	cs.Lock()
	defer cs.Unlock()

	n := cs.connections[id]
	if n != nil {
		n.Stop(reason...)
	}
	delete(cs.connections, id)
}

func (cs *clients) Get(id string) *Node {
	cs.RLock()
	defer cs.RUnlock()
	ret, ok := cs.connections[id]
	if ok {
		return ret
	}
	return nil
}

func (cs *clients) Length() (int, int) {
	cs.RLock()
	defer cs.RUnlock()
	return len(cs.connections), len(cs.events)
}

func (cs *clients) Reset() {
	cs.Lock()
	defer cs.Unlock()

	for id, n := range cs.connections {
		if n != nil {
			n.Stop("server reset clients")
		}
		delete(cs.connections, id)
	}
}

func (cs *clients) SetNodeKey(id, key string) {
	cs.Lock()
	defer cs.Unlock()
	node := cs.connections[id]
	delete(cs.connections, id)
	cs.connections[key] = node
	node.id = key
}
