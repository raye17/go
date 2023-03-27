package main

import "fmt"

func main() {
	oldCalc := &oldCalculatorImpl{}
	adapter := &calculatorAdapter{oldCalc: oldCalc}
	sum := adapter.add(3, 4)
	fmt.Println(sum)
}

type oldCalculator interface {
	getSum(a, b int) int
}
type oldCalculatorImpl struct {
}

func (o *oldCalculatorImpl) getSum(a, b int) int {
	fmt.Println("this is old calc...")
	return a + b
}

type calculatorAdapter struct {
	oldCalc oldCalculator
}

type newCalculator interface {
	add(a, b int) int
}

func (c *calculatorAdapter) add(a, b int) int {
	fmt.Println("this is add...")
	sum := c.oldCalc.getSum(a, b)
	return sum
}
