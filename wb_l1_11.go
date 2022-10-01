// Реализовать пересечение двух неупорядоченных множеств

package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

// реализация пересечения множеств любых comparable-значений с использованием дженериков
func findInter[T constraints.Ordered](a1, a2 []T) []T {
	// inter - слайс пересекающихся значений
	inter := make([]T, 0)
	// m - временная мапа для поиска пересекающихся значений, где "key" - значение из множества,
	// "value" - количество одинаковых значений
	m := make(map[T]int)

	// внесение всех значений и их количества из первого множества в мапу
	for _, v := range a1 {
		m[v] += 1
	}
	// сравнение существующих в мапе значений со значениями второго множества
	for _, v := range a2 {
		if m[v] != 0 {
			// если значение из второго множества уже существует в мапе
			inter = append(inter, v)
		}
		m[v] -= 1
	}

	return inter
}

func main() {
	i1 := []int{6, 5, 8, 6, 4, 3, 2}
	i2 := []int{10, 6, 78, 5, 21, 6, 9, 6}
	fmt.Println(findInter(i2, i1))

	f1 := []float64{6.52, 5.6, 8.96, 6.12, 4, 3.2, 2.6}
	f2 := []float64{10, 6, 78, 5.6, 21, 6.52, 9, 6.522}
	fmt.Println(findInter(f1, f2))

	s1 := []string{"cat", "cat", "dog", "cat", "tree"}
	s2 := []string{"dog", "cat", "tree", "dog", "cat", "tree", "dog"}
	fmt.Println(findInter(s1, s2))
}
