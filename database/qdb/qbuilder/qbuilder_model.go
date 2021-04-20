package qbuilder

type ModelDefine interface {
	DatabaseName() string
	TableName() string
}
