package qtype_test

import (
	"testing"

	"git.querycap.com/ss/lib/container/qtype"
	"git.querycap.com/ss/lib/encoding/qjson"
)

func TestBytes_Basic(t *testing.T) {
	b1 := qtype.NewBytesString("123")
	b2 := qtype.NewBytesString("456")
	b3 := qtype.NewBytesString("789")

	b1.AppendBytes(b2, b3)
	t.Log(b1.String())
	t.Log(qjson.UnsafeMarshalString(qtype.GetBytesPoolStatus()))
}
