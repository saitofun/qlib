package qbuilder

import (
	"reflect"

	"git.querycap.com/ss/lib/util/qstrings"
)

type Field interface {
}

type FieldEx struct {
	m        ModelDefine          // m model
	mrt      *reflect.Type        // mrt model reflect type
	ctx      *reflect.StructField // ctx struct info
	key      string               // key database field key
	name     string               // name field name
	tab      string               // tab table name
	quoted   string               // quoted `tab`.`name`
	db       string               // db database name
	alias    string               // alias builder alias
	selectEx string
}

func NewFieldEx(m ModelDefine, name string) (*FieldEx, bool) {
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
	return &FieldEx{
		m:      m,
		mrt:    &mrt,
		ctx:    &ctx,
		key:    key,
		name:   name,
		tab:    m.TableName(),
		quoted: "`" + m.TableName() + "`.`" + name + "`",
		db:     m.DatabaseName(),
		alias:  "",
	}, true
}

func (f *FieldEx) WithModel(m ModelDefine) *FieldEx {
	ret := *f
	ret.m = m
	return &ret
}

func (f *FieldEx) As(alias string) *FieldEx {
	ret := *f
	ret.alias = alias
	if alias != "" {
		// @todo
	}
	return &ret
}

func (f *FieldEx) Is(v interface{}) Ex {
	return FieldCondExpr(CondIS, f, v)
}

func (f *FieldEx) Eq(v interface{}) Ex {
	return FieldCondExpr(CondEQ, f, v)
}

func (f *FieldEx) NotEq(v interface{}) Ex {
	return FieldCondExpr(CondNOTEQ, f, v)
}

func (f *FieldEx) Gt(v interface{}) Ex {
	return FieldCondExpr(CondGT, f, v)
}

func (f *FieldEx) Gte(v interface{}) Ex {
	return FieldCondExpr(CondGTE, f, v)
}

// func (f *FieldEx) Lt(v interface{}) Ex {
// }
//
// func (f *FieldEx) Lte(v interface{}) Ex {
// }
//
// func (f *FieldEx) In(v ...interface{}) Ex {
// }
//
// func (f *FieldEx) InQuery(sub Ex) Ex {
// 	return &raw{
// 		append(append(append([]byte("? IN ("), sub.Expr()...)), ')'),
// 		append(append([]interface{}{}, f.quoted), sub.Args()...),
// 	}
// }
//
// func (f *FieldEx) NotIn(v ...interface{}) Ex {
// 	return &raw{
// 		CondNotInExpr(len(v)),
// 		append(append([]interface{}{}, f.quoted), v...),
// 	}
// }
//
// func (f *FieldEx) NotInQuery(sub Ex) Ex {
// 	return &raw{
// 		append(append(append([]byte("? NOT IN ("), sub.Expr()...)), ')'),
// 		append(append([]interface{}{}, f.quoted), sub.Args()...),
// 	}
// }
//
// func (f *FieldEx) Between(v1, v2 interface{}) Ex {
// 	return &raw{
// 		ExprCondBetween,
// 		[]interface{}{f.quoted, v1, v2},
// 	}
// }
//
// func (f *FieldEx) NotBetween(v1, v2 interface{}) Ex {
// 	return &raw{
// 		ExprCondNotBetween,
// 		[]interface{}{f.quoted, v1, v2},
// 	}
// }
//
// func (f *FieldEx) Like(v string) Ex {
// 	return &raw{
// 		ExprCondLike,
// 		[]interface{}{f.quoted, v}, // @todo ? string with quote
// 	}
// }
//
// func (f *FieldEx) LeftLike(v string) Ex {
// 	return &raw{
// 		ExprCondLLike,
// 		[]interface{}{f.quoted, v}, // @todo ? string with quote
// 	}
// }
//
// func (f *FieldEx) RightLike(v string) Ex {
// 	return &raw{
// 		ExprCondRLike,
// 		[]interface{}{f.quoted, v}, // @todo ? string with quote
// 	}
// }
//
