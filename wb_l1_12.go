package main

import "fmt"

func main() {
	s := []string{"cat", "cat", "dog", "cat", "tree"}
	m := make(map[string]int)

	for _, v := range s {
		m[v] += 1
	}
	fmt.Println(m)
}
