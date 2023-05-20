package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Calc", func() {
	It("should sum two numbers", func() {
		calc := new(Calculator)

		Expect(calc.add(1, 2)).To(Equal(3))
	})
})
