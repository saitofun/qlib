package qstrings_test

import (
	"testing"

	"github.com/saitofun/qlib/util/qstrings"
)

func TestToCamelString(t *testing.T) {
	t.Log(qstrings.ToCamelString("sss_sss_sss"))
	t.Log(qstrings.ToSnakeString("sss_sss_sss"))
	t.Log(qstrings.ToSnakeString("SssSssSss"))
	t.Log(qstrings.ToCamelString("SssSssSss"))
}
