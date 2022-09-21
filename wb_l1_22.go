package main

import (
	"fmt"
	"math/big"
)

func main() {
	a, _ := new(big.Int).SetString("123456789012345678934567", 10)
	b, _ := new(big.Int).SetString("444445678987000097", 10)

	// сложение
	sum := a.Add(a, b)
	fmt.Println(sum)
	fmt.Println()

	// вычитание
	sub := a.Sub(a, b)
	fmt.Println(sub)
	fmt.Println()

	// деление
	div := a.Div(a, b)
	fmt.Println(div)
	fmt.Println()

	// умножение
	mult := a.Mul(a, b)
	fmt.Println(mult)
	fmt.Println()

}
