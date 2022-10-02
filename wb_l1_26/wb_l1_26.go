// проверить, что все символы в строке уникальные

package main

import "fmt"

func checkUnoqie(s string) bool {

	symb := make(map[rune]bool)
	for _, r := range s {
		if symb[r] == true {
			return false
		}
		symb[r] = true
	}
	return true
}

func printRes(s string) {
	if len(s) == 0 {
		fmt.Println("Empty string")
		return
	}
	fmt.Printf("%s - %t\n", s, checkUnoqie(s))
}

func main() {
	s := "abcd"
	s1 := "Aabcd"
	s2 := "abBcda"
	s3 := ""

	printRes(s)
	printRes(s1)
	printRes(s2)
	printRes(s3)

}
