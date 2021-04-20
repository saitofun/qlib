package qbuilder

import (
	"reflect"

	"git.querycap.com/ss/lib/util/qstrings"
)

type FieldDefine interface {
}

type Field struct {
	ModelDefine                      // m model
	mrt         *reflect.Type        // mrt model reflect type
	ctx         *reflect.StructField // ctx struct info
	key         string               // key database field key
	name        string               // name field name
	tab         string               // tab table name
	quoted      string               // quoted `tab`.`name`
	db          string               // db database name
}

func NewField(m ModelDefine, name string) (*Field, bool) {
	mrt := reflect.TypeOf(m)
	if mrt.Kind() == reflect.Ptr {
		mrt = mrt.Elem()
	}
	ctx, ok := mrt.FieldByName(name)
	if !ok {
		return nil, false
	}
	key, ok := ctx.Tag.Lookup("db")
	if !ok {
		key = "f_" + qstrings.ToSnakeString(name)
	}
	return &Field{
		ModelDefine: m,
		mrt:         &mrt,
		ctx:         &ctx,
		key:         key,
		name:        name,
		tab:         m.TableName(),
		quoted:      "`" + m.TableName() + "`.`" + name + "`",
		db:          m.DatabaseName(),
	}, true
}

func (f *Field) WithModel(m ModelDefine) *Field {
	return &Field{
		ModelDefine: m,
		mrt:         f.mrt,
		ctx:         f.ctx,
		key:         f.key,
		name:        f.name,
		tab:         f.tab,
		db:          f.db,
	}
}

func (f *Field) Is(v interface{}) Ex {
	if ex, ok := v.(Ex); ok {
		return &raw{
			append(ExprCondIsClause, ex.Expr()...),
			append([]interface{}{f.quoted}, ex.Args()...),
		}
	}
	if fv, ok := v.(*Field); ok {
		v = fv.quoted
	}
	return &raw{ExprCondIs, []interface{}{f.quoted, v}}
}

func (f *Field) Eq(v interface{}) Ex {
	if ex, ok := v.(Ex); ok {
		return &raw{
			append(ExprCondEqClause, ex.Expr()...),
			append([]interface{}{f.quoted}, ex.Args()...),
		}
	}
	if fv, ok := v.(*Field); ok {
		v = fv.quoted
	}
	return &raw{ExprCondEq, []interface{}{f.quoted, v}}
}

func (f *Field) NotEq(v interface{}) Ex {
	if ex, ok := v.(Ex); ok {
		return &raw{
			append(ExprCondNotEqClause, ex.Expr()...),
			append([]interface{}{f.quoted}, ex.Args()...),
		}
	}
	if fv, ok := v.(*Field); ok {
		v = fv.quoted
	}
	return &raw{ExprCondNotEq, []interface{}{f.quoted, v}}
}

func (f *Field) Gt(v interface{}) Ex {
	if ex, ok := v.(Ex); ok {
		return &raw{
			append(append(append(ExprCondGtClause, '('), ex.Expr()...), ')'),
			append([]interface{}{f.quoted}, ex.Args()...),
		}
	}
	if fv, ok := v.(*Field); ok {
		v = fv.quoted
	}
	return &raw{ExprCondGt, []interface{}{f.quoted, v}}
}

func (f *Field) Gte(v interface{}) Ex {
	if ex, ok := v.(Ex); ok {
		return &raw{
			append(ExprCondGteClause, ex.Expr()...),
			append([]interface{}{f.quoted}, ex.Args()...),
		}
	}
	if fv, ok := v.(*Field); ok {
		v = fv.quoted
	}
	return &raw{ExprCondGte, []interface{}{f.quoted, v}}
}

func (f *Field) Lt(v interface{}) Ex {
	if ex, ok := v.(Ex); ok {
		return &raw{
			append(ExprCondLtClause, ex.Expr()...),
			append([]interface{}{f.quoted}, ex.Args()...),
		}
	}
	if fv, ok := v.(*Field); ok {
		v = fv.quoted
	}
	return &raw{ExprCondLt, []interface{}{f.quoted, v}}
}

func (f *Field) Lte(v interface{}) Ex {
	if ex, ok := v.(Ex); ok {
		return &raw{
			append(ExprCondLteClause, ex.Expr()...),
			append([]interface{}{f.quoted}, ex.Args()...),
		}
	}
	if fv, ok := v.(*Field); ok {
		v = fv.quoted
	}
	return &raw{ExprCondLte, []interface{}{f.quoted, v}}
}

func (f *Field) In(v ...interface{}) Ex {
	return &raw{
		CondInExpr(len(v)),
		append(append([]interface{}{}, f.name), v...),
	}
}

func (f *Field) InQuery(sub Ex) Ex {
	return &raw{
		append(append(append([]byte("? IN ("), sub.Expr()...)), ')'),
		append(append([]interface{}{}, f.quoted), sub.Args()...),
	}
}

func (f *Field) NotIn(v ...interface{}) Ex {
	return &raw{
		CondNotInExpr(len(v)),
		append(append([]interface{}{}, f.quoted), v...),
	}
}

func (f *Field) NotInQuery(sub Ex) Ex {
	return &raw{
		append(append(append([]byte("? NOT IN ("), sub.Expr()...)), ')'),
		append(append([]interface{}{}, f.quoted), sub.Args()...),
	}
}

func (f *Field) Between(v1, v2 interface{}) Ex {
	return &raw{
		ExprCondBetween,
		[]interface{}{f.quoted, v1, v2},
	}
}

func (f *Field) NotBetween(v1, v2 interface{}) Ex {
	return &raw{
		ExprCondNotBetween,
		[]interface{}{f.quoted, v1, v2},
	}
}

func (f *Field) Like(v string) Ex {
	return &raw{
		ExprCondLike,
		[]interface{}{f.quoted, v}, // @todo ? string with quote
	}
}

func (f *Field) LeftLike(v string) Ex {
	return &raw{
		ExprCondLeftLike,
		[]interface{}{f.quoted, v}, // @todo ? string with quote
	}
}

func (f *Field) RightLike(v string) Ex {
	return &raw{
		ExprCondRightLike,
		[]interface{}{f.quoted, v}, // @todo ? string with quote
	}
}
