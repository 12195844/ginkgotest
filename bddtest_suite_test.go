package bddtest_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBddtest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bddtest Suite")
}
