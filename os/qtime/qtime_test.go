package qtime_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/saitofun/qlib/os/qtime"
)

func TestTime_JSON(t *testing.T) {
	src := []byte(`"2020-11-11T11:11:11+08:00"`)
	v := qtime.Time{}
	if err := v.UnmarshalJSON(src); err != nil {
		t.Error(err)
		return
	}
	dst, err := v.MarshalJSON()
	if err != nil {
		t.Error(err)
		return
	}
	if string(src) != string(dst) {
		t.Error("unequal")
		return
	}

	dst, err = v.MarshalText()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(dst))

	if err = v.Scan(int64(0)); err != nil {
		t.Error(err)
		return
	}

	v = qtime.Now()
	unix := v.Unix()
	value, err := v.Value()
	if err != nil {
		t.Error(err)
		return
	}
	if unix != value.(int64) {
		t.Error("should equal")
		return
	}
	t.Log(value)

	vs := []qtime.Time{qtime.Zero, {}, qtime.UnixZero}
	empty := `""`
	for i, v := range vs {
		dst, err := v.MarshalJSON()
		if err != nil {
			t.Error(strconv.Itoa(i) + ": " + err.Error())
			return
		}
		t.Log(string(dst))
		if string(dst) != empty {
			t.Error(strconv.Itoa(i) + ": should empty")
			return
		}
		if err := v.UnmarshalJSON(dst); err != nil {
			t.Error(err)
			return
		}
		if !v.IsZero() {
			t.Error(strconv.Itoa(i) + ": should zero")
			return
		}
		value, err = v.Value()
		if err != nil {
			t.Error(strconv.Itoa(i) + ": " + err.Error())
			return
		}
		if value.(int64) != 0 {
			t.Error(strconv.Itoa(i) + ": should zero")
			return
		}
	}
}


func TestNewTime(t *testing.T) {
	fmt.Println(qtime.Now().Unix())
	fmt.Println(qtime.NowLocal().Unix())
}