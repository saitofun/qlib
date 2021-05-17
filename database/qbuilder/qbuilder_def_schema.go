package qbuilder

import (
	"go/ast"
	"reflect"
)

type T interface {
	TableName() string
}

type D interface {
	DatabaseName() string
}

type C interface {
	ColumnName() string
}

type Primary interface {
	Primary() string // Primary primary field name
}

type PrimaryID interface {
	PrimaryID() int // PrimaryID primary field id
}

type Schema struct {
	Name         string
	Table        string
	ModelRT      reflect.Type
	ModelRV      reflect.Value
	fields       []*Field
	nameFields   map[string]*Field
	columnFields map[string]*Field
	primary      *Field
	primaries    []*Field
	indexes      []*Field
}

func Model(m interface{}) (ret *Schema) {
	s := &Schema{}
	if m == nil {
		return // TODO nil
	}
	mt := reflect.ValueOf(m).Type()
	for mt.Kind() == reflect.Slice || mt.Kind() == reflect.Array || mt.Kind() == reflect.Ptr {
		mt = mt.Elem()
	}
	if mt.Kind() != reflect.Struct {
		return // TODO unsupported datatype
	}
	s.ModelRT = mt

	// TODO cache[mt.String()].EXISTED

	mv := reflect.New(mt)
	if t, ok := mv.Interface().(T); ok {
		s.Table = t.TableName()
	} else {
		s.Table = NamingStrategy.TableName(mt.Name())
	}
	s.ModelRV = mv

	s.fields = make([]*Field, 0, mt.NumField())
	for i := 0; i < mt.NumField(); i++ {
		if fs := mt.Field(i); ast.IsExported(fs.Name) {
			s.fields = append(s.fields, s.ParseField(fs))
		}
	}

	ret = s
	return
}

// Clone clone a schema context
func (s *Schema) Clone() *Schema { return s }

// WithValue reset Schema.mv
func (s *Schema) WithValue(m interface{}) *Schema { return nil }

// Field return *Field by StructField index
func (s *Schema) Field(i int) *Field { return s.fields[i] }

// FieldByName return *Field by StructField name
func (s *Schema) FieldByName(name string) *Field { return s.nameFields[name] }

// FieldByColumn return *Field by column name
func (s *Schema) FieldByColumn(col string) *Field { return s.columnFields[col] }

// LookupField return *Field by column name or StructField name
func (s *Schema) LookupField(name string) *Field {
	if r, ok := s.nameFields[name]; ok {
		return r
	}
	if r, ok := s.columnFields[name]; ok {
		return r
	}
	return nil
}

/*
type A struct {
	FieldA int
	FieldB string
	FieldC float64
}

v10 := s.NewSelectDestWithNames("FieldA")
v10 should be reflect.ValueOf(&struct{FieldA int}{})

v11 := s.NewSelectDestSliceWithNames("FieldA")
v11 should be reflect.ValueOf(&([]struct{FieldA int}{}))

v20 := s.NewSelectDestWithIndex(2)
v20 should be reflect.ValueOf(&struct{FieldB string}{})

v21 := s.NewSelectDestSliceWithIndex(2)
v21 should be reflect.ValueOf(&([]struct{FieldB string}{}))

*/

// NewSelectDestByNames return a subset struct by field name
func (s *Schema) NewSelectDestByNames(fn ...string) reflect.Value {
	return reflect.ValueOf(nil)
}

// NewSelectDestSliceByNames return a subset struct slice by field name
func (s *Schema) NewSelectDestSliceByNames(fn ...string) reflect.Value {
	return reflect.ValueOf(nil)
}

// NewSelectDestByIndexes return a subset struct by field index
func (s *Schema) NewSelectDestByIndexes(fi ...int) reflect.Value {
	return reflect.ValueOf(nil)
}

// NewSelectDestSliceByIndexes return a subset struct slice by field index
func (s *Schema) NewSelectDestSliceByIndexes(idx ...int) reflect.Value {
	return reflect.ValueOf(nil)
}

func (s *Schema) NewSelectDest() reflect.Value {
	return reflect.New(s.ModelRT)
}

func (s *Schema) NewSelectDestSlice() reflect.Value {
	slice := reflect.MakeSlice(reflect.SliceOf(reflect.PtrTo(s.ModelRT)), 0, 0)
	ret := reflect.New(slice.Type())
	ret.Elem().Set(slice)
	return ret
}

func (s *Schema) NewSelectDestSliceWithCap(cap int) reflect.Value {
	slice := reflect.MakeSlice(reflect.SliceOf(reflect.PtrTo(s.ModelRT)), 0, cap)
	ret := reflect.New(slice.Type())
	ret.Elem().Set(slice)
	return ret
}

func (s *Schema) ParseField(f reflect.StructField) *Field {
	return nil
}