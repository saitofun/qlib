package qbuilder

var (
	ExprCondIs         = []byte("? IS ?")
	ExprCondEq         = []byte("? = ?")
	ExprCondLt         = []byte("? < ?")
	ExprCondLte        = []byte("? <= ?")
	ExprCondGt         = []byte("? > ?")
	ExprCondGte        = []byte("? >= ?")
	ExprCondBetween    = []byte("? BETWEEN ? and ?")
	ExprCondNotBetween = []byte("? NOT BETWEEN ? and ?")
	ExprCondLike       = []byte("? LIKE '%?%'")
	ExprCondLeftLike   = []byte("? LIKE '?%'")
	ExprCondRightLike  = []byte("? LIKE '%?'")
)

type CondCmp struct {
	args [2]interface{}
}

type CondIs CondCmp

func (*CondIs) Expr() []byte {
	return ExprCondEq
}

func (c *CondIs) Args() []interface{} {
	return c.args[0:2]
}

type CondEq CondCmp

func (*CondEq) Expr() []byte {
	return ExprCondEq
}

func (c *CondEq) Args() []interface{} {
	return c.args[0:2]
}

type CondLt CondCmp

func (*CondLt) Expr() []byte {
	return ExprCondLt
}

func (c *CondLt) Args() []interface{} {
	return c.args[0:2]
}

type CondLte CondCmp

func (*CondLte) Lte() []byte {
	return ExprCondLte
}

func (c *CondLte) Args() []interface{} {
	return c.args[0:2]
}

type CondGt CondCmp

func (*CondCmp) Gt() []byte {
	return ExprCondGt
}

func (c *CondGt) Args() []interface{} {
	return c.args[0:2]
}

type CondGte CondCmp

func (*CondGte) Gte() []byte {
	return ExprCondGte
}

func (c *CondCmp) Args() []interface{} {
	return c.args[0:2]
}

type CondIn struct {
	src interface{}
	tar []interface{}
}

func (c *CondIn) Expr() []byte {
	return CondInExpr(len(c.tar))
}

func (c *CondIn) Args() []interface{} {
	return append(append([]interface{}{}, c.src), c.tar...)
}

type CondInQuery struct {
	src interface{}
	tar *expr
}

func (c *CondInQuery) Expr() []byte {
	return append(append([]byte("? IN ("), c.tar.Bytes()...), ')')
}

func (c *CondInQuery) Args() []interface{} {
	return append(append([]interface{}{}, c.src), c.tar.args...)
}

type CondLike struct {
	src interface{}
	tar string
}

func (*CondLike) Expr() []byte {
	return ExprCondLike
}

func (c *CondLike) Args() []interface{} {
	return []interface{}{c.src, c.tar}
}

type CondLeftLike CondLike

func (*CondLeftLike) Expr() []byte {
	return ExprCondLeftLike
}

func (c *CondLeftLike) Args() []interface{} {
	return []interface{}{c.src, c.tar}
}

type CondRightLike CondLike

func (*CondRightLike) Expr() []byte {
	return ExprCondRightLike
}

func (c *CondRightLike) Args() []interface{} {
	return []interface{}{c.src, c.tar}
}

