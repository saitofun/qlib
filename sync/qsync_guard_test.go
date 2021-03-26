package sync_test

import (
	"testing"

	"git.querycap.com/ss/lib/sync"
)

func TestGuard(t *testing.T) {
	var v Tv
	v.Do(func() { v.csa(10, 11) })
}

type Tv struct {
	*sync.Guard
	Val int
}

func (v *Tv) csa(pv int, nv int) (swapped bool) {
	if v.Val == pv {
		v.Val = nv
		swapped = true
		return
	}
	return
}
