package container

import (
	"git.querycap.com/ss/lib/container/qtype"
)

type _boolean interface {
	Clone() *qtype.Bool
	Val() (val bool)
	Set(new bool) (old bool)
	CAS(old, new bool) (swapped bool)
}

type _int interface {
	Clone() *qtype.Int
	Val() int
	Set(new int) (old int)
	CAS(old, new int) (swapped bool)
	Add(delta int) (new int)
}

type _int8 interface {
	Clone() *qtype.Int8
	Val() int8
	Set(new int8) (old int8)
	CAS(old, new int8) (swapped bool)
	Add(delta int8) (new int8)
}

type _int16 interface {
	Clone() *qtype.Int16
	Val() int16
	Set(new int16) (old int16)
	CAS(old, new int16) (swapped bool)
	Add(delta int16) (new int16)
}

type _int32 interface {
	Clone() *qtype.Int32
	Val() int32
	Set(new int32) (old int32)
	CAS(old, new int32) (swapped bool)
	Add(delta int32) (new int32)
}

type _int64 interface {
	Clone() *qtype.Int64
	Val() int64
	Set(new int64) (old int64)
	CAS(old, new int64) (swapped bool)
	Add(delta int64) (new int64)
}

type _uint interface {
	Clone() *qtype.UInt
	Val() uint
	Set(new uint) (old uint)
	CAS(old, new uint) (swapped bool)
	Add(delta uint) (new uint)
}

type _uint8 interface {
	Clone() *qtype.UInt8
	Val() uint8
	Set(new uint8) (old uint8)
	CAS(old, new uint8) (swapped bool)
	Add(delta uint8) (new uint8)
}

type _uint16 interface {
	Clone() *qtype.UInt16
	Val() uint16
	Set(new uint16) (old uint16)
	CAS(old, new uint16) (swapped bool)
	Add(delta uint16) (new uint16)
}

type _uint32 interface {
	Clone() *qtype.UInt32
	Val() uint32
	Set(new uint32) (old uint32)
	CAS(old, new uint32) (swapped bool)
	Add(delta uint32) (new uint32)
}

type _uint64 interface {
	Clone() *qtype.UInt64
	Val() uint64
	Set(new uint64) (old uint64)
	CAS(old, new uint64) (swapped bool)
	Add(delta uint64) (new uint64)
}

type _float32 interface {
	Clone() *qtype.Float32
	Val() float32
	Set(new float32) (old float32)
	CAS(old, new float32) (swapped bool)
	Add(delta float32) (new float32)
}

type _float64 interface {
	Clone() *qtype.Float64
	Val() float64
	Set(new float64) (old float64)
	CAS(old, new float64) (swapped bool)
	Add(delta float64) (new float64)
}

type _string interface {
	Clone() *qtype.String
	Val() string
	Set(new string) (old string)
}

var (
	_ _boolean = (*qtype.Bool)(nil)
	_ _int     = (*qtype.Int)(nil)
	_ _int8    = (*qtype.Int8)(nil)
	_ _int16   = (*qtype.Int16)(nil)
	_ _int32   = (*qtype.Int32)(nil)
	_ _int64   = (*qtype.Int64)(nil)
	_ _uint    = (*qtype.UInt)(nil)
	_ _uint8   = (*qtype.UInt8)(nil)
	_ _uint16  = (*qtype.UInt16)(nil)
	_ _uint32  = (*qtype.UInt32)(nil)
	_ _uint64  = (*qtype.UInt64)(nil)
	_ _float32 = (*qtype.Float32)(nil)
	_ _float64 = (*qtype.Float64)(nil)
	_ _string  = (*qtype.String)(nil)
)
