// Реализовать конкурентную запись данных в map.

package main

import (
	"fmt"
	"sync"
)

func main() {
	m := sync.Map{}
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(10)

	// отправка в канал данных для дальнейшей записи в мапу
	go func() {
		defer close(ch)

		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	// запись данных в sync.Map
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()

			m.Store(i, fmt.Sprintf("value%d", i))
		}(i)
	}
	wg.Wait()

	// вывод всех значений
	m.Range(func(k, v interface{}) bool {
		fmt.Println("key:", k, ", val:", v)
		return true
	})
}
