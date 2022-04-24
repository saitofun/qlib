package qnaming_test

import (
	"testing"

	"github.com/saitofun/qlib/util/qnaming"
	"github.com/stretchr/testify/require"
)

func TestNaming(t *testing.T) {
	tr := require.New(t)

	name := "i_am_a_10_years_senior"

	tr.Equal(qnaming.LowerCamelCase(name), "iAmA10YearsSenior")
	tr.Equal(qnaming.LowerSnakeCase(name), name)
	tr.Equal(qnaming.UpperCamelCase(name), "IAmA10YearsSenior")
	tr.Equal(qnaming.UpperSnakeCase(name), "I_AM_A_10_YEARS_SENIOR")
}
