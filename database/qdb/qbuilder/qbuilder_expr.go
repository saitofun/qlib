package qbuilder

import (
	"bytes"
)

var (
	// DMLs
	ExprSELECT   = []byte("SELECT")
	ExprUPDATE   = []byte("UPDATE")
	ExprDELETE   = []byte("DELETE FROM")
	ExprINSERT   = []byte("INSERT INTO")
	ExprVALUES   = []byte("VALUES")
	ExprSET      = []byte("SET")
	ExprFROM     = []byte("FROM")
	ExprWHERE    = []byte("WHERE")
	ExprGROUPBY  = []byte("GROUP BY")
	ExprORDERBY  = []byte("ORDER BY")
	ExprDESC     = []byte("DESC")
	ExprHAVING   = []byte("HAVING")
	ExprDISDINCT = []byte("DISTINCT")
	ExprLIMIT    = []byte("LIMIT")
	ExprOFFSET   = []byte("OFFSET")

	// alias
	ExprAS = []byte("AS")

	// condition operators
	ExprIN         = []byte("IN")
	ExprAND        = []byte("AND")
	ExprOR         = []byte("OR")
	ExprXOR        = []byte("XOR")
	ExprIS         = []byte("IS")
	ExprISNOTNULL  = []byte("IS NOT NULL")
	ExprISNULL     = []byte("IS NULL")
	ExprLIKE       = []byte("LIKE")
	ExprBETWEEN    = []byte("BETWEEN")
	ExprNOTBETWEEN = []byte("NOT BETWEEN")

	// join query
	ExprJOIN      = []byte("JOIN")
	ExprINNERJOIN = []byte("INNER JOIN")
	ExprLEFTJOIN  = []byte("LEFT JOIN")
	ExprRIGHTJOIN = []byte("RIGHT JOIN")
	ExprUNIONJOIN = []byte("UNION JOIN")

	// glues
	ExprGlueSpace = []byte(" ")
	ExprGlueComma = []byte(",")
	ExprGlueEq    = []byte("=")
	ExprGlueLt    = []byte("<")
	ExprGlueGt    = []byte(">")
	ExprGlueLte   = []byte("<=")
	ExprGlueGte   = []byte("<=")

	// quotes
	ExprBracketL = []byte("(")
	ExprBracketR = []byte(")")

	// end and breaks
	ExprEnd            = []byte(";")
	ExprEndWithNewLine = []byte(";\n")
	ExprNewLine        = []byte("\n")
	ExprNewLinePrefix  = []byte("\n  ")
)

// Condition Expressions
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
	ExprCondLeftLike    = []byte("? LIKE '?%'")
	ExprCondRightLike   = []byte("? LIKE '%?'")
)

const (
	CondIS = iota + 1
	CondEQ
	CondNOTEQ
	CondLT
	CondLTE
	CondGT
	CondGTE
	CondBETWEEN
	CondNOTBETWEEN
	CondLIKE
	CondLEFTLIKE
	CondRIGHTLIKE
)

type Ex interface {
	Expr() []byte
	Args() []interface{}
}

type raw struct {
	expr []byte
	args []interface{}
}

func NewRawEx(ex []byte, args ...interface{}) *raw {
	return &raw{ex, args}
}

func (r *raw) Expr() []byte {
	return r.expr
}

func (r *raw) Args() []interface{} {
	return r.args
}

type expr struct {
	*bytes.Buffer
	args []interface{}
}

func (e *expr) Expr() []byte        { return e.Bytes() }
func (e *expr) Args() []interface{} { return e.args }

func Clause(e Ex) Ex {
	return &raw{
		append(append(append([]byte{}, '('), e.Expr()...), ')'),
		e.Args(),
	}
}
