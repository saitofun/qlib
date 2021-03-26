package qdb

type Model interface {
	Name() string
	Fields() []Field
}

type ModelImpl struct {
	PrimaryID
	Fields        []Field
	Indexes       []WithIndex
	UniqueIndexes []WithUniqueIndex
	Comment       string
}

type TestModel struct {
	ID     uint64 `db:"primary_id"`
	Field1 string `db:"index"`
	Field2 string `db:"unique_index"`
}
