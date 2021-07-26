package qbuilder_test

import (
	"reflect"
	"testing"

	"github.com/saitofun/qlib/database/qbuilder"
)

type model struct {
	qbuilder.Primary
	qbuilder.OperationTime
}

func (m model) Database() string { return "test_db" }
func (m model) Schema() string   { return "test_schema" }

type modelp struct {
	qbuilder.Primary
	qbuilder.OperationTime
}

func (m *modelp) Database() string { return "test_db" }
func (m *modelp) Schema() string   { return "test_schema" }

func TestInterfaces(t *testing.T) {
	t.Log(reflect.TypeOf(model{}).Name())
	t.Log(reflect.TypeOf(model{}).String())
}
