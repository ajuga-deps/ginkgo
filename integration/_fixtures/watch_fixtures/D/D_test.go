package D_test

import (
	. "github.com/ajuga-deps/ginkgo/integration/_fixtures/watch_fixtures/C"

	. "github.com/ajuga-deps/ginkgo"
	. "github.com/ajuga-deps/gomega"
)

var _ = Describe("D", func() {
	It("should do it", func() {
		Ω(DoIt()).Should(Equal("done!"))
	})
})
