package qbuilder

type Primary struct {
	ID uint64 `db:"f_id,autoincrement" json:"-"`
}

func (Primary) Primary() []string { return []string{"ID"} }
