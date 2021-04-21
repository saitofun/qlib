package qbuilder

import (
	"fmt"

	"git.querycap.com/ss/lib/container/qtype"
)

var (
	ExprCondIs          = qtype.NewBytesString("? IS ?")
	ExprCondIsClause    = qtype.NewBytesString("? IS ")
	ExprCondEq          = qtype.NewBytesString("? = ?")
	ExprCondEqClause    = qtype.NewBytesString("? = ")
	ExprCondNotEq       = qtype.NewBytesString("? <> ?")
	ExprCondNotEqClause = qtype.NewBytesString("? <> ")
	ExprCondLt          = qtype.NewBytesString("? < ?")
	ExprCondLtClause    = qtype.NewBytesString("? < ")
	ExprCondLte         = qtype.NewBytesString("? <= ?")
	ExprCondLteClause   = qtype.NewBytesString("? <= ")
	ExprCondGt          = qtype.NewBytesString("? > ?")
	ExprCondGtClause    = qtype.NewBytesString("? > ")
	ExprCondGte         = qtype.NewBytesString("? >= ?")
	ExprCondGteClause   = qtype.NewBytesString("? >= ")
	ExprCondBetween     = qtype.NewBytesString("? BETWEEN ? and ?")
	ExprCondNotBetween  = qtype.NewBytesString("? NOT BETWEEN ? and ?")
	ExprCondLike        = qtype.NewBytesString("? LIKE '%?%'")
	ExprCondLLike       = qtype.NewBytesString("? LIKE '?%'")
	ExprCondRLike       = qtype.NewBytesString("? LIKE '%?'")
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
	Expr           *qtype.Bytes
	WithClauseExpr *qtype.Bytes
	ArgLen         int
}

func (c *Cond) Type() CondType { return c.CondType }

var conditions = map[CondType]*Cond{
	CondIS:         {CondIS, ExprCondIs, ExprCondIsClause, 2},
	CondEQ:         {CondEQ, ExprCondEq, ExprCondEqClause, 2},
	CondNOTEQ:      {CondNOTEQ, ExprCondNotEq, ExprCondNotEqClause, 2},
	CondLT:         {CondLT, ExprCondLt, ExprCondLtClause, 2},
	CondLTE:        {CondLTE, ExprCondLte, ExprCondLteClause, 2},
	CondGT:         {CondGT, ExprCondGt, ExprCondGtClause, 2},
	CondGTE:        {CondGTE, ExprCondGte, ExprCondGteClause, 2},
	CondBETWEEN:    {CondBETWEEN, ExprCondBetween, nil, 3},
	CondNOTBETWEEN: {CondNOTBETWEEN, ExprCondNotBetween, nil, 3},
	CondLIKE:       {CondLIKE, ExprCondLike, nil, 2},
	CondLLIKE:      {CondLLIKE, ExprCondLLike, nil, 2},
	CondRLIKE:      {CondRLIKE, ExprCondRLike, nil, 2},
	CondIN:         {CondIN, nil, nil, -1},
}

func FieldCondExpr(typ CondType, f *FieldEx, args ...interface{}) Ex {
	var (
		cond = conditions[typ]
	)
	if cond == nil {
		panic(fmt.Sprintf("Unknown type: %v", typ))
	}
	if len(args) != cond.ArgLen-1 {
		panic(fmt.Sprintf("Unexpected arg len: %v-%v", len(args), cond.ArgLen))
	}
	ex, ok := args[0].(Ex)
	if len(args) == 2 && ok {
		return &raw{
			cond.WithClauseExpr.Clone().Append('(').AppendBytes(ex.Expr()).Append(')'),
			append(append([]interface{}{}, args[0]), ex.Args()...),
		}
	}
	return &raw{
		cond.Expr,
		append([]interface{}{f.key}, args...),
	}
}
