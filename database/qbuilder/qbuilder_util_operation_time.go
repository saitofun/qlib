package qbuilder

import "github.com/saitofun/qlib/os/qtime"

type OperationTime struct {
	CreatedAt qtime.Time `db:"f_created_at,default='0'" json:"createdAt" `
	UpdatedAt qtime.Time `db:"f_updated_at,default='0'" json:"updatedAt"`
	DeletedAt qtime.Time `db:"f_deleted_at,default='0'" json:"deletedAt"`
}

func (v *OperationTime) OnUpdate() { v.UpdatedAt = qtime.Now() }

func (v *OperationTime) OnCreate() { v.CreatedAt, v.UpdatedAt = qtime.Now(), qtime.Now() }

func (v *OperationTime) OnDelete() {}

type OperationTimeWithSoftDelete struct{ OperationTime }

func (v *OperationTimeWithSoftDelete) OnDelete() { v.DeletedAt, v.UpdatedAt = qtime.Now(), qtime.Now() }

func (v *OperationTimeWithSoftDelete) SoftDelete() {}
