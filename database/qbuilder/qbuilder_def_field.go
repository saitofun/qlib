package qbuilder

import (
	"reflect"
)

type Field struct {
	Name         string
	Column       string
	Schema       *Schema
	Tags         StructTags
	DataType     string // DataType golang datatype, eg: string
	SQLType      string // SQLType sql datatype, eg: TEXT
	Struct       reflect.StructField
	Type         reflect.Type
	IndirectType reflect.Type
}

func (f *Field) Clone() *Field                      { return nil }
func (f *Field) WithValue(interface{}) *Field       { return nil }
func (f *Field) WithStructValue(interface{}) *Field { return nil }
func (f *Field) WithSchema(s *Schema) *Field        { return f }
