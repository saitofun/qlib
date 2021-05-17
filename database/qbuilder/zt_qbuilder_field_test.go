package qbuilder_test

import (
	"database/sql"
)

type m struct {
	ID   uint64 `db:"f_id"`
	Path string `db:"f_Path"`
}

var (
	// _  qbuilder.ModelDefine = (*m)(nil)
	db *sql.DB
)

func (m *m) DatabaseName() string { return "event" }
func (m *m) TableName() string    { return "event" }

// func init() {
// 	var err error
// 	db, err = sql.Open("sqlite3", "/Users/sincos/sincos/tmp/ss/event")
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func TestFieldExpr(t *testing.T) {
// 	m := &m{1, "100"}
// 	fID, ok := qbuilder.NewFieldEx(m, "ID")
// 	if !ok {
// 		t.Error("!ok")
// 		return
// 	}
// 	fPath, ok := qbuilder.NewFieldEx(m, "Path")
// 	exprs := []qbuilder.Ex{
// 		fID.Is("xxx"),
// 		fID.Eq(1),
// 		fID.NotEq(1),
// 		fID.Gt(100),
// 		// fID.Gt(qbuilder.NewRawEx([]byte("select ? from ?"), "t_tmp_field", "t_tmp")),
// 		fID.Gte(100),
// 		fID.Lt(100),
// 		fID.Lte(100),
// 	}
// 	for _, e := range exprs {
// 		t.Logf("expr: %s\n", string(e.Expr()))
// 		t.Logf("args: %v\n", e.Args())
// 	}
// 	ex := qbuilder.Select(fPath).From(m).Where(fPath.Eq(1))
// 	db.Query()
// }
//
