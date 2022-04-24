package qstrings_test

import (
	"testing"

	"github.com/saitofun/qlib/util/qstrings"
	"github.com/stretchr/testify/require"
)

func Test_SplitToWords(t *testing.T) {
	sentences := []string{
		"IAmA10YearsSenior",
		"I Am A 10 Years Senior",
		". I_ Am_A_10_Years____Senior__",
		"I-~~ Am\nA\t10 Years *** Senior",
	}
	words := []string{"I", "Am", "A", "10", "Years", "Senior"}

	tr := require.New(t)

	for i := range sentences {
		tr.Equal(words, qstrings.SplitToWords(sentences[i]))
	}
}
