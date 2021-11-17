package addruntimes_fixture_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestProgressFixture(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AddRuntimes Suite")
}

var _ = BeforeSuite(func() { _, _ = fmt.Fprintln(GinkgoWriter, ">BeforeSuite<") })
var _ = AfterSuite(func() { _, _ = fmt.Fprintln(GinkgoWriter, ">AfterSuite<") })
