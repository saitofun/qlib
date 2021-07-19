package qbuilder

import (
	"errors"
	"go/ast"
	"reflect"
	"sync"

	"github.com/saitofun/qlib/database"
)

var (
	ErrSchemaModelNilInput        = errors.New("SchemaModelNilInput")
	ErrSchemaModelUnsupportedType = errors.New("SchemaModelUnsupportedType")
)

type Schema struct {
	Name           string
	ReflectName    string
	DatabaseName   string
	TableName      string
	PkgPath        string
	FullName       string
	ModelRT        reflect.Type
	ModelRV        reflect.Value
	Fields         []*Field
	FieldsByName   map[string]*Field
	FieldsByColumn map[string]*Field
	FieldsByID     map[int]*Field
	Primary        *Field
	Primaries      []*Field
	Indexes        []*Index
}

func RegisterModel(m interface{}) (ret *Schema, err error) {
	if m == nil {
		return nil, ErrSchemaModelNilInput
	}

	// reflect
	mt := reflect.ValueOf(m).Type()
	for mt.Kind() == reflect.Slice || mt.Kind() == reflect.Array || mt.Kind() == reflect.Ptr {
		mt = mt.Elem()
	}
	if mt.Kind() != reflect.Struct {
		return nil, ErrSchemaModelUnsupportedType
	}
	mv := reflect.New(mt)

	// schema basic
	s := &Schema{
		Name:           mt.Name(),
		ReflectName:    mt.String(),
		ModelRT:        mt,
		ModelRV:        mv,
		FieldsByName:   make(map[string]*Field),
		FieldsByColumn: make(map[string]*Field),
		FieldsByID:     make(map[int]*Field),
	}

	// schema.DatabaseName
	if t, ok := mv.Interface().(WithDatabaseName); ok {
		s.DatabaseName = t.DatabaseName()
	} else if t, ok := mv.Interface().(WithSchemaName); ok {
		s.DatabaseName = t.SchemaName()
	}

	// schema.PkgPath
	s.PkgPath = mt.PkgPath()

	if s.DatabaseName == "" {
		s.FullName = s.PkgPath + "." + s.Name
	} else {
		s.FullName = s.DatabaseName + "." + s.Name
	}

	if v := schemas.Get(s.FullName); v != nil {
		return v, nil
	}

	// schema.TableName
	if t, ok := mv.Interface().(WithTableName); ok {
		s.TableName = t.TableName()
	} else {
		TableName(s.Name)
	}

	s.Fields = make([]*Field, 0)
	for i := 0; i < mt.NumField(); i++ {
		fs := mt.Field(i)
		fv := mv.Field(i)
		if ast.IsExported(fs.Name) && !fs.Anonymous {
			continue
		}
		fields := make([]*Field, 0)
		s.ParseField(fs, fv, fields)
		s.Fields = append(s.Fields, fields...)
	}

	for i := range s.Fields {
		s.FieldsByID[i] = s.Fields[i]
		s.FieldsByName[s.Fields[i].Name] = s.Fields[i]
		s.FieldsByColumn[s.Fields[i].Column] = s.Fields[i]
	}
	ret = s
	return
}

// Clone clone a schema context
func (s *Schema) Clone() *Schema { return s }

// WithValue reset Schema.mv
func (s *Schema) WithValue(m interface{}) *Schema { return nil }

// LookupField return *Field by column name or StructField name
func (s *Schema) LookupField(name string) *Field {
	if r, ok := s.FieldsByName[name]; ok {
		return r
	}
	if r, ok := s.FieldsByColumn[name]; ok {
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

func (s *Schema) ParseField(fs reflect.StructField, fv reflect.Value, fields []*Field) {
	tag, ok := fs.Tag.Lookup("db")
	if !ok && fs.Anonymous {
		for i := 0; i < fv.NumField(); i++ {
			s.ParseField(fv.Type().Field(i), fv.Field(i), fields)
		}
	}

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

type SchemaCache struct {
	mtx *sync.Mutex
	val map[string]*Schema
}

var schemas *SchemaCache

func (s *SchemaCache) Exist(name string) bool {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	_, ok := s.val[name]
	return ok
}

func (s *SchemaCache) Get(name string) *Schema {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	return s.val[name]
}
