// Конкурентное рассчитать значения квадратов чисел

package main

import (
	"fmt"
	"sync"
)

func main() {
	ar := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}

	wg.Add(len(ar))
	for _, i := range ar {
		go func(n int) {
			defer wg.Done()
			fmt.Println(n * n)
		}(i)
	}
	wg.Wait()
}
