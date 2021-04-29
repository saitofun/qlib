package qbuilder

// Ex SQL Expressions
type Ex interface {
	Expr() []byte
	Args() []interface{}
}

// raw Raw SQL Expression
type raw struct {
	expr []byte
	args []interface{}
}

func NewRawEx(ex string, args ...interface{}) *raw {
	return nil
}

func newRawEx(ex []byte, args ...interface{}) *raw {
	return nil
}

func (r *raw) Expr() []byte {
	return nil
}

func (r *raw) Args() []interface{} {
	return r.args
}

type CondEx interface {
	Ex
	CondType() CondType
}

type expr struct {
	expr []byte
	args []interface{}
}

func (e *expr) Expr() []byte        { return e.expr }
func (e *expr) Args() []interface{} { return e.args }
