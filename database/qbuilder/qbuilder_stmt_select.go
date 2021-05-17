package qbuilder

// type SelectStmt struct {
// 	Fields    []Field
// 	Models    []ModelDefine
// 	Cond      CondExpr
// 	Orders    []OrderDefine
// 	Joins     []Joins
// 	WithLimit int
// 	*expr
// }
//
// func (s *SelectStmt) From(m ...ModelDefine) *SelectStmt {
// 	s.Models = m
// 	return s
// }
//
// func (s *SelectStmt) Where(cond CondExpr) *SelectStmt {
// 	s.Cond = cond
// 	return s
// }
//
// func (s *SelectStmt) OrderBy(orders ...OrderDefine) *SelectStmt {
// 	s.Orders = orders
// 	return s
// }
//
// func (s *SelectStmt) Limit(lmt int) *SelectStmt {
// 	s.WithLimit = lmt
// 	return s
// }
//
// func Select(f ...Field) *SelectStmt {
// 	return &SelectStmt{Fields: f}
// }
