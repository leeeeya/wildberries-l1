// реверс строки

package main

import "fmt"

func main() {
	str := "Mouse собака мәче 狐"

	fmt.Printf("Input string: %s\n", str)
	n := 0
	runes := make([]rune, len(str))

	// конвертация в руны
	for _, sign := range str {
		runes[n] = sign
		n++
	}
	// реверс
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	// обратная конвертация в строку
	res := string(runes)
	fmt.Printf("Output string: %s\n", res)
}
