package more_ginkgo_tests_test

import (
	. "github.com/ajuga-deps/ginkgo"
	. "github.com/ajuga-deps/ginkgo/integration/_fixtures/more_ginkgo_tests"
	. "github.com/ajuga-deps/gomega"
)

var _ = Describe("MoreGinkgoTests", func() {
	It("should pass", func() {
		Ω(AlwaysTrue()).Should(BeTrue())
	})

	It("should always pass", func() {
		Ω(AlwaysTrue()).Should(BeTrue())
	})
})
