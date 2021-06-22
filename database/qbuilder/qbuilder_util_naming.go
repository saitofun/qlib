package qbuilder

import (
	"strings"

	"github.com/saitofun/qlib/util/qstrings"
)

type DefaultNaming struct{}

// TableName table name for Naming
// eg n.TableName("TabName") -> "t_tab_name"
func (n DefaultNaming) TableName(tab string) string {
	return "t_" + qstrings.ToSnakeString(tab)
}

// ColumnName column name for Naming
// eg n.ColumnName("ColName") -> "f_col_name"
func (n DefaultNaming) ColumnName(tabName, colName string) string {
	return "f_" + qstrings.ToSnakeString(colName)
}

// IndexName index name for Naming
// eg n.IndexName("TabName", "ColName1", "ColName2") -> "t_tab_name_i_col_name_1_col_name_2"
func (n DefaultNaming) IndexName(tabName string, colName ...string) string {
	joiners := []string{"t", qstrings.ToSnakeString(tabName), "i"}
	for _, v := range colName {
		joiners = append(joiners, qstrings.ToSnakeString(v))
	}
	return strings.Join(joiners, "_")
}

// UniqueIndexName index name for Naming
// eg n.UniqueIndexName("TabName", "ColName1", "ColName2") -> "t_tab_name_ui_col_name_1_col_name_2"
func (n DefaultNaming) UniqueIndexName(tabName string, colName ...string) string {
	joiners := []string{"t", qstrings.ToSnakeString(tabName), "ui"}
	for _, v := range colName {
		joiners = append(joiners, qstrings.ToSnakeString(v))
	}
	return strings.Join(joiners, "_")
}

func TableName(tab string) string {
	return NamingStrategy.TableName(tab)
}

func ColumnName(tab string, col string) string {
	return NamingStrategy.ColumnName(tab, col)
}

func IndexName(tab string, col ...string) string {
	return NamingStrategy.IndexName(tab, col...)
}

func UniqueIndexName(tab string, col ...string) string {
	return NamingStrategy.UniqueIndexName(tab, col...)
}

var NamingStrategy Naming = DefaultNaming{}

func SetNamingStrategy(n Naming) { NamingStrategy = n }
