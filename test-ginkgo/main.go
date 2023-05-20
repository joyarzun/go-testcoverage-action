package main

import (
	"fmt"
)

type calculatorI interface {
	add(a, b int) int
}

type Calculator struct {
	calculatorI // Embeds the calculator interface
}

func (m *Calculator) add(a, b int) int { return a + b }

func summarize(calc calculatorI, values ...int) int {
	var result int

	for _, val := range values {
		result = calc.add(result, val)
	}

	return result
}

func main() {
	calc := new(Calculator)
	result := summarize(calc, 10, 12, 9, 11)
	fmt.Println(result) // Prints 42
}
