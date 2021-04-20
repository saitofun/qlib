package qdb

type Model interface {
	Name() string
	Database() string
	Fields() []Field
}
