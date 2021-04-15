package qbuilder

import (
	"bytes"
	"database/sql"

	"github.com/go-courier/sqlx/v2"
)

var (
	ExprSELECT     = []byte("SELECT")
	ExprUPDATE     = []byte("UPDATE")
	ExprDELETE     = []byte("DELETE FROM")
	ExprINSERT     = []byte("INSERT INTO")
	ExprVALUES     = []byte("VALUES")
	ExprSET        = []byte("SET")
	ExprFROM       = []byte("FROM")
	ExprWHERE      = []byte("WHERE")
	ExprGROUPBY    = []byte("GROUP BY")
	ExprORDERBY    = []byte("ORDER BY")
	ExprHAVING     = []byte("HAVING")
	ExprLIKE       = []byte("LIKE")
	ExprBETWEEN    = []byte("BETWEEN")
	ExprNOTBETWEEN = []byte("NOT BETWEEN")
	ExprDISDINCT   = []byte("DISTINCT")
	ExprLIMIT      = []byte("LIMIT")
	ExprOFFSET     = []byte("OFFSET")
	ExprAS         = []byte("AS")
	ExprDESC       = []byte("DESC")
	ExprIN         = []byte("IN")
	ExprAND        = []byte("AND")
	ExprOR         = []byte("OR")
	ExprIS         = []byte("IS")
	ExprJOIN       = []byte("JOIN")
	ExprINNERJOIN  = []byte("INNER JOIN")
	ExprLEFTJOIN   = []byte("LEFT JOIN")
	ExprRIGHTJOIN  = []byte("RIGHT JOIN")
	ExprUNIONJOIN  = []byte("UNION JOIN")

	ExprGlueSpace = []byte(" ")
	ExprGlueComma = []byte(",")
	ExprGlueEq    = []byte("=")
	ExprGlueLt    = []byte("<")
	ExprGlueGt    = []byte(">")
	ExprGlueLte   = []byte("<=")
	ExprGlueGte   = []byte("<=")

	ExprBracketL       = []byte("(")
	ExprBracketR       = []byte(")")
	ExprEnd            = []byte(";")
	ExprEndWithNewLine = []byte(";\n")
	ExprNewLine        = []byte("\n")
	ExprNewLinePrefix  = []byte("\n  ")
)

type expr struct {
	*bytes.Buffer
	Args []interface{}
}

func (e *expr) WriteExprs(raw ...[]byte) *expr {
	for i := range raw {
		e.Write(raw[i])
		e.Write(ExprGlueSpace)
	}
	e.Write(ExprGlueSpace)
	return e
}

func (e *expr) Begin(ex []byte) *expr {
	e.Buffer.Reset()
	e.Write(ex)
	e.Write(ExprGlueSpace)
	return e
}

func (e *expr) WriteExpr(ex []byte) *expr {
	e.Write(ex)
	e.Write(ExprGlueSpace)
	return e
}

func (e *expr) WriteExprAndContinue(ex []byte) *expr {
	return e.WriteExpr(ex).WriteNewLinePrefix()
}

func (e *expr) End() *expr {
	e.Write(ExprEndWithNewLine)
	e.WriteByte('\n')
	return e
}

func (e *expr) WriteNewLine() *expr {
	e.WriteByte('\n')
	return e
}

func (e *expr) WriteNewLinePrefix() *expr {
	e.WriteByte('\n')
	e.Write(ExprGlueSpace)
	return e
}

func (e *expr) WriteRaw(ex []byte) *expr {
	e.Write(ex)
	return e.End()
}

func (e *expr) WriteAdditions(additions ...Addition) *expr { return e }

func (e *expr) Select(v Model, additions ...Addition) *expr {
	e.Begin(ExprSELECT).WriteNewLinePrefix().
		WriteExprAndContinue(v.FieldsExpr()).
		WriteExprAndContinue(ExprFROM).
		WriteExprAndContinue(v.TableName())
	if len(additions) > 0 {
		e.WriteExprAndContinue(ExprWHERE).
			WriteAdditions(additions...).End()
	}
	return e
}

func (e *expr) Exec(db *sql.DB) (sql.Result, error) {
	defer e.Release()
	return db.Exec(e.String(), e.Args...)
}

func (e *expr) QueryAndScan(db *sql.DB, v interface{}) error {
	defer e.Release()
	rows, err := db.Query(e.String(), e.Args...)
	if err != nil {
		return err
	}
	return sqlx.Scan(rows, v)
}

func (e *expr) Release() {}
