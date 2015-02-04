package coverage_fixture_test

import (
	. "github.com/ajuga-deps/ginkgo/integration/_fixtures/coverage_fixture"

	. "github.com/ajuga-deps/ginkgo"
	. "github.com/ajuga-deps/gomega"
)

var _ = Describe("CoverageFixture", func() {
	It("should test A", func() {
		Ω(A()).Should(Equal("A"))
	})

	It("should test B", func() {
		Ω(B()).Should(Equal("B"))
	})

	It("should test C", func() {
		Ω(C()).Should(Equal("C"))
	})

	It("should test D", func() {
		Ω(D()).Should(Equal("D"))
	})
})
