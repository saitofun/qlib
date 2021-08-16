package qtype_test

import (
	"sync/atomic"
	"testing"

	"github.com/saitofun/qlib/container/qtype"
)

func BenchmarkAtomic_Store(b *testing.B) {
	v := atomic.Value{}
	for i := 0; i < b.N; i++ {
		v.Store(i)
	}
}

func BenchmarkAny_Set(b *testing.B) {
	v := qtype.New()
	for i := 0; i < b.N; i++ {
		v.Set(i)
	}
}

func BenchmarkBool_SetTure(b *testing.B) {
	v := qtype.NewBool()
	for i := 0; i < b.N; i++ {
		v.Set(true)
	}
}

func BenchmarkBool_SetFalse(b *testing.B) {
	v := qtype.NewBool()
	for i := 0; i < b.N; i++ {
		v.Set(false)
	}
}

func BenchmarkFloat32_Set(b *testing.B) {
	v := qtype.NewFloat32()
	for i := 0; i < b.N; i++ {
		v.Set(float32(i))
	}
}

func BenchmarkFloat64_Set(b *testing.B) {
	v := qtype.NewFloat64()
	for i := float64(0); i < float64(b.N); i++ {
		v.Set(i)
	}
}

func BenchmarkInt_Set(b *testing.B) {
	v := qtype.NewInt()
	for i := 0; i < b.N; i++ {
		v.Set(i)
	}
}

func BenchmarkInt8_Set(b *testing.B) {
	v := qtype.NewInt8()
	for i := int8(0); i < int8(b.N); i++ {
		v.Set(i)
	}
}

func BenchmarkInt16_Set(b *testing.B) {
	v := qtype.NewInt16()
	for i := 0; i < b.N; i++ {
		v.Set(int16(i))
	}
}

func BenchmarkInt32_Set(b *testing.B) {
	v := qtype.NewInt32()
	for i := int32(0); i < int32(b.N); i++ {
		v.Set(i)
	}
}

func BenchmarkInt64_Set(b *testing.B) {
	v := qtype.NewInt64()
	for i := 0; i < b.N; i++ {
		v.Set(int64(i))
	}
}

func BenchmarkUint_Set(b *testing.B) {
	v := qtype.NewUInt()
	for i := uint(0); i < uint(b.N); i++ {
		v.Set(uint(i))
	}
}

func BenchmarkUInt8_Set(b *testing.B) {
	v := qtype.NewUInt8()
	for i := 0; i < b.N; i++ {
		v.Set(uint8(i))
	}
}

func BenchmarkUInt16_Set(b *testing.B) {
	v := qtype.NewUInt16()
	for i := 0; i < b.N; i++ {
		v.Set(uint16(i))
	}
}

func BenchmarkUint32_Set(b *testing.B) {
	v := qtype.NewUInt32()
	for i := uint32(0); i < uint32(b.N); i++ {
		v.Set(i)
	}
}

func BenchmarkString_Set(b *testing.B) {
	v := qtype.NewString()
	for i := 0; i < b.N; i++ {
		v.Set("")
	}
}

func BenchmarkAtomic_Load(b *testing.B) {
	v := atomic.Value{}
	for i := 0; i < b.N; i++ {
		v.Load()
	}
}

func BenchmarkAny_Val(b *testing.B) {
	v := qtype.New()
	for i := 0; i < b.N; i++ {
		v.Val()
	}
}

func BenchmarkBool_Val(b *testing.B) {
	v := qtype.NewBool()
	for i := 0; i < b.N; i++ {
		v.Val()
	}
}

func BenchmarkFloat32_Val(b *testing.B) {
	v := qtype.NewFloat32()
	for i := 0; i < b.N; i++ {
		v.Val()
	}
}

func BenchmarkFloat64_Val(b *testing.B) {
	v := qtype.NewFloat64()
	for i := float64(0); i < float64(b.N); i++ {
		v.Val()
	}
}

func BenchmarkInt_Val(b *testing.B) {
	v := qtype.NewInt()
	for i := 0; i < b.N; i++ {
		v.Val()
	}
}

func BenchmarkInt8_Val(b *testing.B) {
	v := qtype.NewInt8()
	for i := int8(0); i < int8(b.N); i++ {
		v.Val()
	}
}

func BenchmarkInt16_Val(b *testing.B) {
	v := qtype.NewInt16()
	for i := 0; i < b.N; i++ {
		v.Set(int16(i))
	}
}

func BenchmarkInt32_Val(b *testing.B) {
	v := qtype.NewInt32()
	for i := int32(0); i < int32(b.N); i++ {
		v.Val()
	}
}

func BenchmarkInt64_Val(b *testing.B) {
	v := qtype.NewInt64()
	for i := 0; i < b.N; i++ {
		v.Val()
	}
}

func BenchmarkUint_Val(b *testing.B) {
	v := qtype.NewUInt()
	for i := uint(0); i < uint(b.N); i++ {
		v.Val()
	}
}

func BenchmarkUInt8_Val(b *testing.B) {
	v := qtype.NewUInt8()
	for i := 0; i < b.N; i++ {
		v.Val()
	}
}

func BenchmarkUInt16_Val(b *testing.B) {
	v := qtype.NewUInt16()
	for i := 0; i < b.N; i++ {
		v.Val()
	}
}

func BenchmarkUint32_Val(b *testing.B) {
	v := qtype.NewUInt32()
	for i := uint32(0); i < uint32(b.N); i++ {
		v.Val()
	}
}

func BenchmarkString_Val(b *testing.B) {
	v := qtype.NewString()
	for i := 0; i < b.N; i++ {
		v.Val()
	}
}
