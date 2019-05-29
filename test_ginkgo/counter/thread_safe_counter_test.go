package counter_test

import (
	. "github.com/dhiller/go_playground/test_ginkgo/counter"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TestSubject", func() {

	var (
		testSubject ThreadSafeCounter
	)

	BeforeEach(func() {
		testSubject = ThreadSafeCounter{}
	})

	Context("Basic counter functions", func() {

		It("is initially zero", func() {
			Expect(testSubject.Value()).To(Equal(0))
		})

		It("increases the counter", func() {
			testSubject.Inc()
			Expect(testSubject.Value()).To(Equal(1))
		})

		It("increases and decreases the counter", func() {
			testSubject.Inc()
			testSubject.Dec()
			Expect(testSubject.Value()).To(Equal(0))
		})

	})

	Context("Async behavior", func() {

		It("increases the counter", func() {
			go testSubject.Inc()
			Eventually(func() int {
				return testSubject.Value()
			}).Should(Equal(1))
		})

		It("increases and decreases the counter", func() {
			testSubject.Inc()
			go func() {
				testSubject.Dec()
			}()
			Eventually(func() int {
				return testSubject.Value()
			}).Should(Equal(0))
		})

	})

})
