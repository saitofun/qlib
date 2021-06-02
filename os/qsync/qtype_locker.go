package qsync

import "github.com/saitofun/qlib/container/qtype"

const (
	free = false
	busy = true
)

type Locker qtype.Bool

func NewLocker() *Locker { return (*Locker)(qtype.NewBool()) }

func (l *Locker) Lock() bool { return (*qtype.Bool)(l).CAS(free, busy) }

func (l *Locker) Unlock() { (*qtype.Bool)(l).Set(free) }
