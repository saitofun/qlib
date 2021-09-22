package qbuilder

import (
	"reflect"
)

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

// Naming naming interface
type Naming interface {
	TableName(tab string) string
	ColumnName(tab, col string) string
	IndexName(tab string, col ...string) string
	UniqueIndexName(tab string, col ...string) string
}

type Datatype struct {
	Type    reflect.Type
	SqlType func(dialect string) string
}

type Database struct {
	Name string
	*Tables
}

// Register register table to database
func (d *Database) Register(tables ...*Table) {}

// Table get table by name
func (d *Database) Table(name string) *Table { return nil }

type Table struct {
	*Database
	*Fields
	Name string
}

func (t *Table) Register(fields ...*Field) {}

func (t *Table) Field(string) *Field { return nil }

type Tables struct {
	*Database
	tables []*Table
	values map[string]int
}

func (t *Tables) Add(table *Table) {}

type Field struct {
	*Table
	Name     string
	Datatype // @todo
}

func Parse() *Field { return nil }

type Fields struct {
	*Table
	fields []*Field
	values map[string]int
}

func (f *Fields) Add(field *Field) {}

type Index struct {
	Name string
	*Table
	*Fields
}

func Alias() {}

type expr struct {
	expr string
	args []interface{}
}
