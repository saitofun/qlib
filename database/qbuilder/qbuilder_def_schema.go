package qbuilder

import (
	"errors"
	"go/ast"
	"reflect"

	"git.querycap.com/ss/lib/database"
)

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

var (
	ErrSchemaModelNilInput        = errors.New("SchemaModelNilInput")
	ErrSchemaModelUnsupportedType = errors.New("SchemaModelUnsupportedType")
)

func Model(m interface{}) (ret *Schema, err error) {
	if m == nil {
		return nil, ErrSchemaModelNilInput
	}
	mt := reflect.ValueOf(m).Type()
	for mt.Kind() == reflect.Slice || mt.Kind() == reflect.Array || mt.Kind() == reflect.Ptr {
		mt = mt.Elem()
	}
	if mt.Kind() != reflect.Struct {
		return nil, ErrSchemaModelUnsupportedType
	}
	s := &Schema{ModelRT: mt}

	// TODO cache[mt.String()].EXISTED
	mv := reflect.New(mt)
	if t, ok := mv.Interface().(database.T); ok {
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
func (s *Schema) FieldByCol(col string) *Field { return s.columnFields[col] }

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

func (s *Schema) ParseField(fs reflect.StructField) *Field {
	f := &Field{
		Name:         fs.Name,
		Column:       "",
		Schema:       s,
		Tags:         ParseTags(fs.Tag, ","),
		DataType:     "",
		SQLType:      "",
		Struct:       fs,
		Type:         fs.Type,
		IndirectType: fs.Type,
	}
	for f.IndirectType.Kind() == reflect.Ptr {
		f.IndirectType = f.IndirectType.Elem()
	}

	fv := reflect.New(f.IndirectType)

	if c, ok := fv.Interface().(database.C); ok {
		f.Column = c.ColumnName()
	} else {
		if f.Column = f.Tags.ColumnName(); f.Column == "" {
			f.Column = NamingStrategy.ColumnName(s.Table, f.Name)
		}
	}

	if t, ok := fv.Interface().(database.SQLType); ok {
		f.SQLType = t.SQLType("")
	} else {
		switch f.IndirectType.Kind() {
		case reflect.Bool:
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		case reflect.Float32, reflect.Float64:
		case reflect.String:
		case reflect.Struct:
		case reflect.Slice, reflect.Array:
		}
	}
	return f
}

// ArgsEx argument list expression: eg (?,?,...)
func (s *Schema) ArgsEx() string { return "" }

// AssignEx assignment list expression: eg: f_a=?,f_b=?,...
func (s *Schema) AssignEx() string { return "" }

// ColumnsEx column list expression: eg (f_a,f_b,...)
func (s *Schema) ColumnsEx() string { return "" }

// Insert insert SQL query: INSERT INTO t_tab (f_a,f_b,...) VALUES (?,?,..);
func (s *Schema) Insert(m ...interface{}) Ex { return nil }

// Update update by primary SQL query: UPDATE t_tab SET f_a=?,f_b=?,... WHERE f_a=? AND f_b<?...
func (s *Schema) Update(m interface{}, cond ...Cond) Ex { return nil }

// Delete delete by primary SQL query: DELETE FROM t_tab WHERE f_id=?
func (s *Schema) Delete(m interface{}, cond ...Cond) Ex { return nil }

// Select select by primary SQL query: SELECT * FROM t_tab WHERE f_a=? AND f_b>?...
func (s *Schema) Select(m interface{}, cond ...Cond) Ex { return nil }
