package qptr

func NewInt(v int) *int             { return &v }
func NewInt8(v int8) *int8          { return &v }
func NewInt16(v int16) *int16       { return &v }
func NewInt32(v int32) *int32       { return &v }
func NewInt64(v int64) *int64       { return &v }
func NewUint(v uint) *uint          { return &v }
func NewUint8(v uint8) *uint8       { return &v }
func NewUint16(v uint16) *uint16    { return &v }
func NewUint32(v uint32) *uint32    { return &v }
func NewUint64(v uint64) *uint64    { return &v }
func NewFloat32(v float32) *float32 { return &v }
func NewFloat64(v float64) *float64 { return &v }
