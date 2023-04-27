package main

import (
	"fmt"
	calc "math/calculator"
)

func main() {
	fmt.Println(calc.SumWith(112, 10))
	fmt.Println(calc.SubtractBy(112, 10))
	fmt.Println(calc.MultipleBy(112, 10))
	fmt.Println(calc.DivideBy(112, 10))
}
