package qbuilder_test

import (
	"testing"

	"git.querycap.com/ss/lib/database/qdb/qbuilder"
)

type m struct {
	A int    `db:"f_a"`
	B string `db:"f_b"`
}

//go:linkname

var _ qbuilder.ModelDefine = (*m)(nil)

func (m *m) DatabaseName() string { return "tmp" }
func (m *m) TableName() string    { return "tmp" }

func TestFieldExpr(t *testing.T) {
	m := &m{1, "100"}
	fa, ok := qbuilder.NewField(m, "A")
	if !ok {
		t.Error("!ok")
		return
	}
	exprs := []qbuilder.Ex{
		fa.Is("xxx"),
		fa.Eq(1),
		fa.NotEq(1),
		fa.Gt(100),
		fa.Gt(qbuilder.NewRawEx([]byte("select ? from ?"), "t_tmp_field", "t_tmp")),
		fa.Gte(100),
		fa.Lt(100),
		fa.Lte(100),
	}
	for _, e := range exprs {
		t.Logf("expr: %s\n", string(e.Expr()))
		t.Logf("args: %v\n", e.Args())
	}
}
