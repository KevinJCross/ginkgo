package addruntimes_fixture_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
)

var _ = Describe("ProgressFixture", func() {
	BeforeEach(func() {
		_, _ = fmt.Fprintln(GinkgoWriter, ">outer before<")
	})

	JustBeforeEach(func() {
		_, _ = fmt.Fprintln(GinkgoWriter, ">outer just before<")
	})

	AfterEach(func() {
		_, _ = fmt.Fprintln(GinkgoWriter, ">outer after<")
	})

	Context("Inner Context", func() {
		BeforeEach(func() {
			_, _ = fmt.Fprintln(GinkgoWriter, ">inner before<")
		})

		JustBeforeEach(func() {
			_, _ = fmt.Fprintln(GinkgoWriter, ">inner just before<")
		})

		AfterEach(func() {
			_, _ = fmt.Fprintln(GinkgoWriter, ">inner after<")
		})

		When("Inner When", func() {
			BeforeEach(func() {
				_, _ = fmt.Fprintln(GinkgoWriter, ">inner before<")
			})

			It("should emit progress as it goes", func() {
				_, _ = fmt.Fprintln(GinkgoWriter, ">it<")
			})
		})
	})

	Specify("should emit progress as it goes", func() {
		_, _ = fmt.Fprintln(GinkgoWriter, ">specify<")
	})
})
