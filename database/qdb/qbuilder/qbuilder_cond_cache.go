package qbuilder

import (
	"sync"
)

type condInCache struct {
	*sync.Mutex
	vIns    map[int][]byte
	vNotIns map[int][]byte
}

func (c *condInCache) GetIn(n int) []byte {
	c.Lock()
	defer c.Unlock()
	if ret, ok := c.vIns[n]; ok {
		return ret
	}
	ret := []byte("? IN (")
	for i := 1; i < n; i++ {
		ret = append(ret, '?', ',')
	}
	ret = append(ret, '?', ')')
	c.vIns[n] = ret
	return ret
}

func (c *condInCache) GetNotIn(n int) []byte {
	c.Lock()
	defer c.Unlock()
	if ret, ok := c.vNotIns[n]; ok {
		return ret
	}
	ret := []byte("? NOT IN (")
	for i := 1; i < n; i++ {
		ret = append(ret, '?', ',')
	}
	ret = append(ret, '?', ')')
	c.vNotIns[n] = ret
	return ret
}

var cacheCondIn = &condInCache{
	Mutex:   &sync.Mutex{},
	vIns:    make(map[int][]byte),
	vNotIns: make(map[int][]byte),
}

func CondInExpr(tars int) []byte {
	return cacheCondIn.GetIn(tars)
}

func CondNotInExpr(tars int) []byte {
	return cacheCondIn.GetNotIn(tars)
}
