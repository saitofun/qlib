package qbuilder

import (
	"sync"
)

type condInCache struct {
	*sync.Mutex
	v map[int][]byte
}

func (c *condInCache) Get(n int) []byte {
	c.Lock()
	defer c.Unlock()
	if ret, ok := c.v[n]; ok {
		return ret
	}
	ret := make([]byte, 0)
	for i := 1; i < n; i++ {
		ret = append(ret, '?', ',')
	}
	ret = append(ret, '?')
	c.v[n] = ret
	return ret
}

var cacheCondIn = &condInCache{
	Mutex: &sync.Mutex{},
	v:     make(map[int][]byte),
}

func CondInExpr(tars int) []byte {
	return cacheCondIn.Get(tars)
}
