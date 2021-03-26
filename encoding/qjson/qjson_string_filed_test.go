package qjson_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"git.querycap.com/aisys/lib/encoding/qjson"
	"git.querycap.com/aisys/lib/qtime"
)

type Struct struct {
	A Embedded   `json:"a,string"`
	B qtime.Time `json:"b"`
}

type Embedded struct {
	C int    `json:"c,string"`
	D string `json:"d,int"`
}

// func (v Embedded) MarshalJSON() ([]byte, error) {
// 	type tmp Embedded
// 	var val = tmp(v)
// 	return qjson.Stringer.Marshal(val)
// }
//
// func (v *Embedded) UnmarshalJSON(data []byte) error {
// 	type tmp Embedded
// 	var val = (*tmp)(v)
// 	return qjson.Stringer.Unmarshal(data, val)
// }

func TestJson(t *testing.T) {
	var val = Struct{
		A: Embedded{C: 10, D: "10"},
		B: qtime.Time{Time: time.Now()},
	}
	var str, _ = json.MarshalIndent(val, "", "  ")
	fmt.Println(string(str))
	val = Struct{}
	if err := qjson.Unmarshal(str, &val); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
}
