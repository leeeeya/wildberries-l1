// реверс слов в строке

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun - just a set of words"
	fmt.Printf("Input string: %s\n", str)

	// конвертация в слайс строк
	sl := strings.Split(str, " ")
	l := len(sl)

	// реверс
	for i := 0; i < l/2; i++ {
		sl[i], sl[l-1-i] = sl[l-1-i], sl[i]
	}

	// обратная конвертация
	res := strings.Join(sl, " ")
	fmt.Printf("Output string: %s\n", res)
}
