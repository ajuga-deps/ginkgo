package eventually_failing_test

import (
	"fmt"
	"io/ioutil"
	"time"

	. "github.com/ajuga-deps/ginkgo"
	. "github.com/ajuga-deps/gomega"
)

var _ = Describe("EventuallyFailing", func() {
	It("should fail on the third try", func() {
		time.Sleep(time.Second)
		files, err := ioutil.ReadDir(".")
		Ω(err).ShouldNot(HaveOccurred())
		Ω(len(files)).Should(BeNumerically("<", 5))
		ioutil.WriteFile(fmt.Sprintf("./%d", len(files)), []byte("foo"), 0777)
	})
})
