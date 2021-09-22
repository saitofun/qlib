package database

type Datatype interface{ Datatype(dialect string) string }

type Field struct{}

type Index struct{}

type Table struct{}

type Option map[string]string
