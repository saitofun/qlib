package qbuilder

import (
	"sync"

	"github.com/saitofun/qlib/os/qsync"
)

var ArgListCache = struct {
	val map[int]string
	mtx *sync.Mutex
}{
	val: make(map[int]string),
	mtx: &sync.Mutex{},
}

func ArgListExpr(args ...interface{}) *expr {
	size := len(args)
	if size == 0 {
		return nil
	}
	ret := &expr{args: args}
	qsync.Guard(ArgListCache.mtx).Do(func() {
		ret.expr = ArgListCache.val[size]
	})
	if ret.expr == "" {
		ret.expr = "("
		for i := 1; i < size; i++ {
			ret.expr += "?,"
		}
		ret.expr += "?)"
		qsync.Guard(ArgListCache.mtx).Do(func() {
			ArgListCache.val[len(args)] = ret.expr
		})
	}
	return ret
}
