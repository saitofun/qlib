package qtype_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"git.querycap.com/ss/lib/container/qtype"
)

func TestNew(t *testing.T) {
	// b := qtype.NewBool()
	// dat, err := b.MarshalJSON()
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }
	// fmt.Println(string(dat))

	// dat = []byte("true")
	// err = b.UnmarshalJSON(dat)
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }
	// fmt.Println(len(dat), b.Val())

	// i := qtype.NewInt()
	// dat, err = i.MarshalJSON()
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }
	// fmt.Println(string(dat))

	// dat = []byte("-1")
	// err = i.UnmarshalJSON(dat)
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }
	// fmt.Println(len(dat), i.Val())

	// f := qtype.NewFloat64()
	// f.SetLan(0.00002)
	// dat, err = f.MarshalJSON()
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }
	// fmt.Println(string(dat))

	// dat = []byte("0.0000000002")
	// err = f.UnmarshalJSON(dat)
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }
	// fmt.Println(len(dat), f.Val())

	// s := qtype.NewString()
	// s.SetLan("0.00002")
	// dat, err = s.MarshalJSON()
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }
	// fmt.Println(string(dat))

	// dat = []byte(`"0.002"`)
	// err = s.UnmarshalJSON(dat)
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }
	// fmt.Println(len(dat), s.Val())
	var v = struct {
		INT     qtype.Int     `json:"int"`
		INT8    qtype.Int8    `json:"int8"`
		INT16   qtype.Int16   `json:"int16"`
		INT32   qtype.Int32   `json:"int32"`
		INT64   qtype.Int64   `json:"int64"`
		UINT    qtype.UInt    `json:"uint"`
		UINT8   qtype.UInt8   `json:"uint8"`
		UINT16  qtype.UInt16  `json:"uint16"`
		UINT32  qtype.UInt32  `json:"uint32"`
		UINT64  qtype.UInt64  `json:"uint64"`
		FLOAT32 qtype.Float32 `json:"float32"`
		FLOAT64 qtype.Float64 `json:"float64"`
		STRING  qtype.String  `json:"string"`
		BOOL    qtype.Bool    `json:"bool"`
	}{
	}
	b, _ := json.MarshalIndent(v, "", "    ")
	fmt.Println(string(b))
}
