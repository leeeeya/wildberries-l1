//Реализовать бинарный поиск встроенными методами языка.

package main

import (
	"./wb_l1_16/quicksort"
	"fmt"
)

func binarySearch(haystack []int, needle int) (int, int, bool) {
	var mid, lookingFor int

	minIdx := 0
	maxIdx := len(haystack) - 1

	for minIdx <= maxIdx {
		// нахождение середины слайса
		mid = (minIdx + maxIdx) / 2
		lookingFor = haystack[mid]
		if lookingFor == needle {
			return mid, lookingFor, true
		}
		// если искомое значение больше
		if lookingFor > needle {
			maxIdx = mid - 1
		} else {
			minIdx = mid + 1
		}
	}
	return 0, 0, false
}

func searchAndPrint(haystack []int, needle int) {
	if len(haystack) == 0 {
		fmt.Println("Empty array")
		return
	}
	fmt.Println("Array:", haystack)
	fmt.Println("Looking for:", needle)
	fmt.Println()

	// бинарный поиск работает только если массив отсортирован
	quicksort.Quicksort(haystack)

	idx, value, ok := binarySearch(haystack, needle)
	if ok {
		fmt.Printf("Found: %d, index in sorted array: %d\n", value, idx)
	} else {
		fmt.Println("Nothing found")
	}
	fmt.Println()
}

func main() {
	haystack1 := []int{20, 5, 3, 54, -2, 9}
	needle1 := 21

	haystack2 := []int{1, 2, 3, 5, 7, 9, 10}
	needle2 := 10

	haystack3 := make([]int, 0)
	needle3 := 10

	searchAndPrint(haystack1, needle1)
	fmt.Println()
	searchAndPrint(haystack2, needle2)
	fmt.Println()
	searchAndPrint(haystack3, needle3)
}
