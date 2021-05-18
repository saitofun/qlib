package database

// D database name
type D interface {
	DatabaseName() string
}

// T table name
type T interface {
	TableName() string
}

// C database column
type C interface {
	ColumnName() string
}

// Primary primary field name
type Primary interface {
	Primary() string
}

// PrimaryID primary field index
type PrimaryID interface {
	PrimaryID() int
}

type SQLType interface {
	SQLType(dialect DialectName) string //
}

type DialectName string

const (
	SQLITE     DialectName = "sqlite"
	MySQL                  = "mysql"
	PostgreSQL             = "postgresql"
	SQLServer              = "sqlserver"
)
