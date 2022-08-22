package qnaming_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/saitofun/qlib/util/qnaming"
)

func TestNaming(t *testing.T) {
	name := "i_am_a_10_years_senior"

	NewWithT(t).Expect(qnaming.LowerCamelCase(name)).To(Equal("iAmA10YearsSenior"))
	NewWithT(t).Expect(qnaming.LowerSnakeCase(name)).To(Equal("i_am_a_10_years_senior"))
	NewWithT(t).Expect(qnaming.UpperCamelCase(name)).To(Equal("IAmA10YearsSenior"))
	NewWithT(t).Expect(qnaming.UpperSnakeCase(name)).To(Equal("I_AM_A_10_YEARS_SENIOR"))

	NewWithT(t).Expect(qnaming.UpperCamelCase("OrgID")).To(Equal("OrgID"))
}
