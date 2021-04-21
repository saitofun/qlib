package qbuilder

import (
	"git.querycap.com/ss/lib/container/qtype"
)

var (
	// DMLs
	ExprSELECT   = qtype.NewBytesString("SELECT")
	ExprUPDATE   = qtype.NewBytesString("UPDATE")
	ExprDELETE   = qtype.NewBytesString("DELETE FROM")
	ExprINSERT   = qtype.NewBytesString("INSERT INTO")
	ExprVALUES   = qtype.NewBytesString("VALUES")
	ExprSET      = qtype.NewBytesString("SET")
	ExprFROM     = qtype.NewBytesString("FROM")
	ExprWHERE    = qtype.NewBytesString("WHERE")
	ExprGROUPBY  = qtype.NewBytesString("GROUP BY")
	ExprORDERBY  = qtype.NewBytesString("ORDER BY")
	ExprDESC     = qtype.NewBytesString("DESC")
	ExprHAVING   = qtype.NewBytesString("HAVING")
	ExprDISDINCT = qtype.NewBytesString("DISTINCT")
	ExprLIMIT    = qtype.NewBytesString("LIMIT")
	ExprOFFSET   = qtype.NewBytesString("OFFSET")

	// alias
	ExprAS = qtype.NewBytesString("AS")

	// condition operators
	ExprIN         = qtype.NewBytesString("IN")
	ExprAND        = qtype.NewBytesString("AND")
	ExprOR         = qtype.NewBytesString("OR")
	ExprXOR        = qtype.NewBytesString("XOR")
	ExprIS         = qtype.NewBytesString("IS")
	ExprISNOTNULL  = qtype.NewBytesString("IS NOT NULL")
	ExprISNULL     = qtype.NewBytesString("IS NULL")
	ExprLIKE       = qtype.NewBytesString("LIKE")
	ExprBETWEEN    = qtype.NewBytesString("BETWEEN")
	ExprNOTBETWEEN = qtype.NewBytesString("NOT BETWEEN")

	// join query
	ExprJOIN      = qtype.NewBytesString("JOIN")
	ExprINNERJOIN = qtype.NewBytesString("INNER JOIN")
	ExprLEFTJOIN  = qtype.NewBytesString("LEFT JOIN")
	ExprRIGHTJOIN = qtype.NewBytesString("RIGHT JOIN")
	ExprUNIONJOIN = qtype.NewBytesString("UNION JOIN")
)

// Ex SQL Expressions
type Ex interface {
	Expr() *qtype.Bytes
	Args() []interface{}
}

// raw Raw SQL Expression
type raw struct {
	expr *qtype.Bytes
	args []interface{}
}

func NewRawEx(ex string, args ...interface{}) *raw {
	return &raw{qtype.NewBytesString(ex), args}
}

func newRawEx(ex *qtype.Bytes, args ...interface{}) *raw {
	return &raw{ex, args}
}

func (r *raw) Expr() *qtype.Bytes {
	return r.expr
}

func (r *raw) Args() []interface{} {
	return r.args
}

type CondEx interface {
	Ex
	CondType() CondType
}

type expr struct {
	expr *qtype.Bytes
	args []interface{}
}

func (e *expr) Expr() *qtype.Bytes  { return e.expr }
func (e *expr) Args() []interface{} { return e.args }
