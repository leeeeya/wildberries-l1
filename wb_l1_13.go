// Поменять местами два числа без создания временной переменной.

package main

import "fmt"

// свап через кортеж
func simpleSwap(a, b int) {

	fmt.Printf("Before swap: %d, %d\n", a, b)
	b, a = a, b
	fmt.Printf("After swap: %d, %d\n", a, b)
}

//свап через сумму
func swapBySum(a, b int) {
	fmt.Printf("Before swap: %d, %d\n", a, b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("After swap: %d, %d\n", a, b)
}

func main() {
	a, b := 3, 5
	simpleSwap(a, b)
	fmt.Println()
	swapBySum(a, b)
}
