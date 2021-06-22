package qbuilder

type Query struct {
	Expr string
	Args []interface{}
}

func JoinQueries(queries ...*Query) *Query {
	ret := &Query{}
	for _, q := range queries {
		ret.Expr += q.Expr
		ret.Expr += " "
		ret.Args = append(ret.Args, q.Args...)
	}
	return ret
}

type WithPrimary interface {
	Primary() string
	PrimaryID() uint64
}

type WithSoftDelete interface {
	SoftDelete()
}

type WithOperationTimes interface {
	OnUpdate()
	OnCreate()
	OnDelete()
}

type Naming interface {
	TableName(tab string) string
	ColumnName(tab, col string) string
	IndexName(tab string, col ...string) string
	UniqueIndexName(tab string, col ...string) string
}
