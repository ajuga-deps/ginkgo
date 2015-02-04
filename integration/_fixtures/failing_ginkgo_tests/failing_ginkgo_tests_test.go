package failing_ginkgo_tests_test

import (
	. "github.com/ajuga-deps/ginkgo"
	. "github.com/ajuga-deps/ginkgo/integration/_fixtures/failing_ginkgo_tests"
	. "github.com/ajuga-deps/gomega"
)

var _ = Describe("FailingGinkgoTests", func() {
	It("should fail", func() {
		Ω(AlwaysFalse()).Should(BeTrue())
	})

	It("should pass", func() {
		Ω(AlwaysFalse()).Should(BeFalse())
	})
})
