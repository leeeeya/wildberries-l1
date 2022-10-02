package main

import (
	"./quicksort"
	"fmt"
)

// сортировка для множеств значений разных comparable-типов данных, реализация через дженерики
func main() {

	emptyArr := make([]int, 0)
	quicksort.SortAndPrint(emptyArr)
	fmt.Println()

	sortedArr := []int{1, 2, 3, 4, 5, 6}
	quicksort.SortAndPrint(sortedArr)
	fmt.Println()

	intArr := []int{2, -1, 0, 5, 8, 9, 10, 11}
	quicksort.SortAndPrint(intArr)
	fmt.Println()

	strArr := []string{"a", "D", "B", "dfghj", ""}
	quicksort.SortAndPrint(strArr)
	fmt.Println()

	floatArr := []float32{3.2, 5.8, -10.6, 0, 5.0}
	quicksort.SortAndPrint(floatArr)
}
