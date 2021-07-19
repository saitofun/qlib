package qbuilder

import (
	"sync"
)

type CondType int

const (
	CondIS CondType = iota + 1
	CondEQ
	CondNOTEQ
	CondLT
	CondLTE
	CondGT
	CondGTE
	CondBETWEEN
	CondNOTBETWEEN
	CondLIKE
	CondLLIKE
	CondRLIKE
	CondIN
)

type Cond struct {
	CondType
	Expr           []byte
	WithClauseExpr []byte
	ArgLen         int
}

func (c *Cond) Type() CondType { return c.CondType }

type FieldCond interface {
	Field() *Field
	Type() CondType
}

var FieldCondCache = struct {
	val map[string]map[string][14]string
	mtx *sync.Mutex
}{
	val: make(map[string]map[string][14]string),
	mtx: &sync.Mutex{},
}

func FieldCondExpr(f *Field, t CondType, arg ...interface{}) CondEx {
	// qsync.Guard(FieldCondCache.mtx).Do(func() {
	// 	if len(FieldCondCache.val) == 0 {

	// 	}
	// 	if _, ok := FieldCondCache.val[f.Schema.Database]; !ok {
	// 		FieldCondCache.val[f.Schema.Database] = make(map[string][14]string)
	// 	}
	// })

	// switch t {
	// case CondIS, CondEQ, CondNOTEQ, CondLT, CondLTE, CondGT, CondGTE:
	// case CondBETWEEN, CondNOTBETWEEN:
	// case CondLIKE, CondLLIKE, CondRLIKE:
	// case CondIN:
	// }
	return AsCond(nil)
}
