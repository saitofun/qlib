package qbuilder

var (
	ExprCondIs          = []byte("? IS ?")
	ExprCondIsClause    = []byte("? IS ")
	ExprCondEq          = []byte("? = ?")
	ExprCondEqClause    = []byte("? = ")
	ExprCondNotEq       = []byte("? <> ?")
	ExprCondNotEqClause = []byte("? <> ")
	ExprCondLt          = []byte("? < ?")
	ExprCondLtClause    = []byte("? < ")
	ExprCondLte         = []byte("? <= ?")
	ExprCondLteClause   = []byte("? <= ")
	ExprCondGt          = []byte("? > ?")
	ExprCondGtClause    = []byte("? > ")
	ExprCondGte         = []byte("? >= ?")
	ExprCondGteClause   = []byte("? >= ")
	ExprCondBetween     = []byte("? BETWEEN ? and ?")
	ExprCondNotBetween  = []byte("? NOT BETWEEN ? and ?")
	ExprCondLike        = []byte("? LIKE '%?%'")
	ExprCondLLike       = []byte("? LIKE '?%'")
	ExprCondRLike       = []byte("? LIKE '%?'")
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

var cond = map[CondType]*Cond{
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

