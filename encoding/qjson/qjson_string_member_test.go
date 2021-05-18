package qjson_test

import (
	"fmt"
	"testing"
	"time"

	"git.querycap.com/ss/lib/encoding/qjson"
	"git.querycap.com/ss/lib/os/qtime"
)

type Struct struct {
	A Embedded   `json:"a,string"`
	B qtime.Time `json:"b"`
}

type Embedded struct {
	C                        int    `json:"c,string"`
	D                        string `json:"d,int"`
	qjson.StringMemberMarker `json:"-"`
}

func (v Embedded) MarshalJSON() ([]byte, error) {
	if v.Stepped() {
		return nil, nil
	}
	v.Step()
	return qjson.Marshal(v)
}

func (v *Embedded) UnmarshalJSON(data []byte) error {
	return qjson.Unmarshal(data, v)
}

func TestJson(t *testing.T) {
	var val = Struct{
		A: Embedded{C: 10, D: "10"},
		B: qtime.Time{Time: time.Now()},
	}
	var str, _ = qjson.MarshalIndent(val, "", "  ")
	fmt.Println(string(str))
	val = Struct{}
	if err := qjson.Unmarshal(str, &val); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
}
