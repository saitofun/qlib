package qdb

type WithIndex interface {
	Index() interface{}
}

type WithUniqueIndex interface {
	UniqueIndex() interface{}
}

type UnionIndex struct {
	Indexes []WithIndex
}

func (UnionIndex) Index() interface{} {
	return nil
}

type UnionUniqueIndex struct {
	UniqueIndexes []WithUniqueIndex
}

func (UnionUniqueIndex) UniqueIndex() interface{} {
	return nil
}
