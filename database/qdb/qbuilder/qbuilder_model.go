package qbuilder

import (
	"reflect"
	"sync"
)

type Model interface {
	TableName() []byte
	FieldsExpr() []byte
	FieldsExprWithQuote() []byte
	FieldNames() [][]byte
	FieldNamesWithQuote() [][]byte
}

type TypesCache struct {
	*sync.Mutex
	fields    map[reflect.Type][]string
	fieldExpr map[reflect.Type]string
	values    map[reflect.Type]func(interface{}) []string
	valueExpr map[reflect.Type]func(interface{}) string
}

func (t *TypesCache) GetFields(v interface{}, zero ...string) []string {
	return []string{}
}

func (t *TypesCache) GetFieldExpr(v interface{}) string {
	return ""
}

func (t *TypesCache) _(v interface{}) string {
	return ""
}
