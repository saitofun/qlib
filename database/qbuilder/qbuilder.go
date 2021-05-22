package qbuilder

type Query struct {
	Expr string
	Args []interface{}
}

func Join(queries ...*Query) *Query {
	ret := &Query{}
	for _, q := range queries {
		ret.Expr += q.Expr
		ret.Expr += " "
		ret.Args = append(ret.Args, q.Args...)
	}
	return ret
}
