package bddtest_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	"github.com/studyzy/bddtest"
)

var _ = Describe("Bddtest", func() {
	BeforeSuite(func() {
		fmt.Println("BeforeSuite")
	})
	AfterSuite(func() {
		fmt.Println("AfterSuite")
	})
	BeforeEach(func() {
		fmt.Println("BeforeEach")
	})
	JustBeforeEach(func() {
		fmt.Println("JustBeforeEach")
	})
	JustAfterEach(func() {
		fmt.Println("JustAfterEach")
	})
	AfterEach(func() {
		fmt.Println("AfterEach")
	})
	var foo=&bddtest.Foo{}
	Describe("Describe increase foo", func() {
		fmt.Println("increase foo result:", foo.Increase())
	})
	Context("Context foo.Add", func() {
		fmt.Println("foo.Add 10",foo.Add(10))
		It("Add 20", func() {
			fmt.Println("foo.Add 20",foo.Add(20))
			By("too large", func() {
				fmt.Println("By~~~~")
			})
		})
		Specify("Add 30", func() {
			fmt.Println("foo.Add 30",foo.Add(30))
			//Fail("Add 30 fail")

		})
		When("added 40", func() {
			fmt.Println("foo.Add 40",foo.Add(40))
			It("It 40~~~", func() {
				fmt.Println("It 40~~!!!!")
			})
		})

	})
})
