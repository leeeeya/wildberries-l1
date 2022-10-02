// Разработать программу которая устанавливает i-й бит в 1 или 0 в переменной типа int64

package main

import (
	"fmt"
	"log"
)

func inverseBit(n, i int64) int64 {
	fmt.Printf("given number as a set of bytes: %b\n", n)
	// инициализация числа, i-ый бит которого противоположен i-ому биту числа n
	var inv int64 = 1 << i

	fmt.Printf("number for inversion: %b\n", inv)
	// инверсия бита через "исключающее или"
	n = n ^ inv

	fmt.Printf("result as a set of bytes: %b\n", n)
	return n
}

func main() {
	var n, i int64
	// считывание числа (n) и номера бита, который надо изменить в n
	if _, err := fmt.Scan(&n, &i); err != nil {
		log.Fatalln(err)
	}
	fmt.Println()

	fmt.Printf("given number after inversion: %d\n", inverseBit(n, i))
}
