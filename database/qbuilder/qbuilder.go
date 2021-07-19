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

type WithOperationTime interface {
	OnUpdate()
	OnCreate()
	OnDelete()
}

// WithDatabaseName model's schema name
type WithDatabaseName interface {
	DatabaseName() string
}

// WithSchemaName model's schema name
type WithSchemaName interface {
	SchemaName() string
}

type WithTableName interface {
	TableName() string
}

// WithIndexes model's indexes
type WithIndexes interface {
	Indexes() []Index
}

type Naming interface {
	TableName(tab string) string
	ColumnName(tab, col string) string
	IndexName(tab string, col ...string) string
	UniqueIndexName(tab string, col ...string) string
}
