package qbuilder

// Ex SQL Expressions
type Ex interface {
	Expr() string
	Args() []interface{}
}

type expr struct {
	expr string
	args []interface{}
}

func (e *expr) Expr() string        { return e.expr }
func (e *expr) Args() []interface{} { return e.args }

type CondMarker interface {
	condMarker()
}

type CondEx struct {
	Ex
	CondMarker
}

func AsCond(ex Ex) CondEx {
	return struct {
		Ex
		CondMarker
	}{Ex: ex}
}
