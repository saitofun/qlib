package qptr

func Bool(v bool) *bool          { return &v }
func Int(v int) *int             { return &v }
func Int8(v int8) *int8          { return &v }
func Int16(v int16) *int16       { return &v }
func Int32(v int32) *int32       { return &v }
func Int64(v int64) *int64       { return &v }
func Uint(v uint) *uint          { return &v }
func Uint8(v uint8) *uint8       { return &v }
func Uint16(v uint16) *uint16    { return &v }
func Uint32(v uint32) *uint32    { return &v }
func Uint64(v uint64) *uint64    { return &v }
func Float32(v float32) *float32 { return &v }
func Float64(v float64) *float64 { return &v }
func String(v string) *string    { return &v }

func BoolValue(v *bool) bool {
	if v != nil {
		return *v
	}
	return false
}

func IntValue(v *int) int {
	if v != nil {
		return *v
	}
	return 0
}

func Int8Value(v *int8) int8 {
	if v != nil {
		return *v
	}
	return 0
}

func Int16Value(v *int16) int16 {
	if v != nil {
		return *v
	}
	return 0
}

func Int32Value(v *int32) int32 {
	if v != nil {
		return *v
	}
	return 0
}

func Int64Value(v *int64) int64 {
	if v != nil {
		return *v
	}
	return 0
}

func UintValue(v *uint) uint {
	if v != nil {
		return *v
	}
	return 0
}

func Uint8Value(v *uint8) uint8 {
	if v != nil {
		return *v
	}
	return 0
}

func Uint16Value(v *uint16) uint16 {
	if v != nil {
		return *v
	}
	return 0
}

func Uint32Value(v *uint32) uint32 {
	if v != nil {
		return *v
	}
	return 0
}

func Uint64Value(v *uint64) uint64 {
	if v != nil {
		return *v
	}
	return 0
}

func Float32Value(v *float32) float32 {
	if v != nil {
		return *v
	}
	return 0
}

func Float64Value(v *float64) float64 {
	if v != nil {
		return *v
	}
	return 0
}

func StringValue(v *string) string {
	if v != nil {
		return *v
	}
	return ""
}
