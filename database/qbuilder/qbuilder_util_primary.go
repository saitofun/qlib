package qbuilder

type Primary struct {
	ID uint64 `db:"f_id,autoincrement" json:"-"`
}

func (v *Primary) Primary() []string { return []string{"ID"} }

func (v *Primary) PrimaryID() uint64 { return v.ID }

