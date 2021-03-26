package qdb

type Querier interface {
	Query(interface{} /*model*/, ...WithCond) (interface{} /*record converted*/, error)
}
