package qbuilder

import (
	"reflect"
	"sync"
)

type ModelCtx struct {
	rt         reflect.Type
	rv         reflect.Value
	table      []byte
	fieldNames [][]byte
	fieldExpr  []byte
}

func NewModelCache(v Model) *ModelCtx {
	return &ModelCtx{
		rt:        reflect.TypeOf(v),
		rv:        reflect.ValueOf(v),
		table:     v.TableName(),
		fieldExpr: v.FieldsExpr(),
	}
}

type ModelCache struct {
	mu  *sync.RWMutex
	val map[string]*ModelCtx
}

var models *ModelCache

func init() {
	models = &ModelCache{
		mu:  &sync.RWMutex{},
		val: make(map[string]*ModelCtx),
	}
}
