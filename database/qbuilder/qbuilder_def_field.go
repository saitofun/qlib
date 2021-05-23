package qbuilder

import (
	"fmt"
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

func (f *Field) Database() string {
	if f.Schema != nil {
		return f.Schema.Database
	}
	return ""
}

func (f *Field) ColumnName() string       { return f.Column }
func (f *Field) QuotedColumnName() string { return fmt.Sprintf("`%s`.`%s`", f.Schema.Table, f.Column) }
