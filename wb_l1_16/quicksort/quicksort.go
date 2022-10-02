// Реализовать быструю сортировку массива (quicksort) встроенными методами языка

package quicksort

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

// проверка на то, что массив отсортирован
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

	// рекурсивное деление массива на два подмассива и сортировка каждого из них
	Quicksort(ar[:split])
	Quicksort(ar[split:])
}

// разделение массива на две части по опорной точке
func partition[T constraints.Ordered](ar []T) int {
	// нахождение опорной точки ("середины" массива)
	mid := ar[len(ar)/2]

	left := 0
	right := len(ar) - 1

	for {
		// все значения меньше mid оказываются слева
		for ; ar[left] < mid; left++ {
		}
		// все значения больше mid оказываются справа
		for ; ar[right] > mid; right-- {
		}

		if left >= right {
			return right
		}
		// если слева оказалось значение больше mid, а справа - меньше, они меняются местами
		ar[left], ar[right] = ar[right], ar[left]
	}
}

func SortAndPrint[T constraints.Ordered](ar []T) {
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
