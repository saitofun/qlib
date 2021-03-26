package qtime_test

import (
	"fmt"
	"testing"

	"git.querycap.com/aisys/lib/qtime"
)

func TestTime_UnmarshalJSON(t *testing.T) {
	v := qtime.Now()
	err := v.UnmarshalJSON([]byte("\"2020-11-11 11:11:11.000\""))
	fmt.Println(err, v)
}
