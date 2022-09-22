// Найти сумму квадратов чисел с использованием конкурентных вычислений.

package main

import (
	"fmt"
	"sync"
)

func main() {
	ar := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	ch := make(chan int)
	var sum int

	wg.Add(len(ar))
	for _, i := range ar {
		go func(n int, ch chan int) {
			defer wg.Done()
			ch <- n * n
		}(i, ch)
		sum += <-ch
	}
	wg.Wait()
	fmt.Println(sum)
}
