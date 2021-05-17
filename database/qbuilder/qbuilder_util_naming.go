package qbuilder

type Naming interface {
	TableName(tab string) string
	ColumnName(tab, col string) string
	IndexName(tab string, col ...string) string
}

type DefaultNaming struct{}

// TableName table name for Naming
// eg n.TableName("TabName") -> "t_tab_name"
func (n DefaultNaming) TableName(tabName string) string {
	return "t_tab_name"
}

// ColumnName column name for Naming
// eg n.ColumnName("ColName") -> "f_col_name"
func (n DefaultNaming) ColumnName(tabName, colName string) string {
	return "f_col_name"
}

// IndexName index name for Naming
// eg n.IndexName("TabName", "ColName1", "ColName2") -> "t_tab_name_i_col_name_1_col_name_2"
func (n DefaultNaming) IndexName(tabName string, colName ...string) string {
	return "t_tab_name_i_col_name_1_col_name_2"
}

// UniqueIndexName index name for Naming
// eg n.UniqueIndexName("TabName", "ColName1", "ColName2") -> "t_tab_name_ui_col_name_1_col_name_2"
func (n DefaultNaming) UniqueIndexName(tabName string, colName ...string) string {
	return "t_tab_name_ui_col_name_1_col_name_2"
}

var NamingStrategy Naming = DefaultNaming{}

func SetNamingStrategy(n Naming) {
	if n != nil {
		NamingStrategy = n
	}
}
