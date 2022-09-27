// Реализовать быструю сортировку массива встроенными методами языка

package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func Quicksort[T constraints.Ordered](ar []T) {
	if len(ar) <= 1 {
		return
	}

	split := partition(ar)

	Quicksort(ar[:split])
	Quicksort(ar[split:])
}

func partition[T constraints.Ordered](ar []T) int {
	pivot := ar[len(ar)/2]

	left := 0
	right := len(ar) - 1

	for {
		for ; ar[left] < pivot; left++ {
		}
		for ; ar[right] > pivot; right-- {
		}

		if left >= right {
			return right
		}
		ar[left], ar[right] = ar[right], ar[left]
	}
}

func sortAndPrint[T constraints.Ordered](sl []T) {
	fmt.Println("Initial array:", sl)
	Quicksort(sl)
	fmt.Println("Sorted array:", sl)
}

func main() {
	intArr := []int{2, 5, 8, 1, -1}
	sortAndPrint(intArr)
	fmt.Println()

	strArr := []string{"a", "D", "B", "dfghj", ""}
	sortAndPrint(strArr)
	fmt.Println()

	floatArr := []float32{3.2, 5.8, -10.6, 0, 5.0}
	sortAndPrint(floatArr)
}
