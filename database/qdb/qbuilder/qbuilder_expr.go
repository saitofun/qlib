package qbuilder

import (
	"bytes"
	"database/sql"

	"github.com/go-courier/sqlx/v2"
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

type expr struct {
	*bytes.Buffer
	args []interface{}
}

func InstanceExpr(query string, args ...interface{}) *expr {
	return &expr{Buffer: bytes.NewBufferString(query), args: args}
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
	return db.Exec(e.String(), e.args...)
}

func (e *expr) QueryAndScan(db *sql.DB, v interface{}) error {
	defer e.Release()
	rows, err := db.Query(e.String(), e.args...)
	if err != nil {
		return err
	}
	return sqlx.Scan(rows, v)
}

func (e *expr) Release() {}
