package qdriver

import "sync"

type Driver interface {
	Open()
	Tables()
	TableColumns()
	TableIndexes()
	Options()
	Datatype()
}

var drivers = struct {
	val map[string]Driver
	mtx sync.Mutex
}{}
