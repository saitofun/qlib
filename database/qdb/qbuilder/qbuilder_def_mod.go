package qbuilder

import (
	"reflect"
	"sync"
)

type Model interface {
	DB() string
	Tab() string
}

type Table struct {
	db  string
	tab string
	rv  reflect.Type
}

func RegisterModel(v interface{}) *Table {
	rv := reflect.TypeOf(v)
	for rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		panic("expect reflect.Struct")
	}
	return nil
}

func (t *Table) ModelName() string {
	return t.rv.Name()
}

func RegisterOnce(m ...Model) {
	(&sync.Once{}).Do(func() {
	})
}
