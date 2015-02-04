package B_test

import (
	. "github.com/ajuga-deps/ginkgo/integration/_fixtures/watch_fixtures/B"

	. "github.com/ajuga-deps/ginkgo"
	. "github.com/ajuga-deps/gomega"
)

var _ = Describe("B", func() {
	It("should do it", func() {
		Ω(DoIt()).Should(Equal("done!"))
	})
})
