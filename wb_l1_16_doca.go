// Реализовать быструю сортировку массива встроенными методами языка

package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

// todo dodelat' docu
func checkSorted[T constraints.Ordered](ar []T) bool {
	for i := 0; i < len(ar)-1; i++ {
		if ar[i] > ar[i+1] {
			return false
		}
	}
	return true
}

func Quicksort[T constraints.Ordered](ar []T) {
	if len(ar) <= 1 || checkSorted(ar) == true {
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

func sortAndPrint[T constraints.Ordered](ar []T) {
	if len(ar) == 0 {
		fmt.Println("Empty array")
		return
	}
	if checkSorted(ar) == true {
		fmt.Println("Array is already sorted:", ar)
		return
	}
	fmt.Println("Initial array:", ar)
	Quicksort(ar)
	fmt.Println("Sorted array:", ar)
	fmt.Println()
}

func main() {

	emptyArr := make([]int, 0)
	sortAndPrint(emptyArr)
	fmt.Println()

	sortedArr := []int{1, 2, 3, 4, 5, 6}
	sortAndPrint(sortedArr)
	fmt.Println()

	intArr := []int{2, 5, 8, 1, -1}
	sortAndPrint(intArr)
	fmt.Println()

	strArr := []string{"a", "D", "B", "dfghj", ""}
	sortAndPrint(strArr)
	fmt.Println()

	floatArr := []float32{3.2, 5.8, -10.6, 0, 5.0}
	sortAndPrint(floatArr)
}
