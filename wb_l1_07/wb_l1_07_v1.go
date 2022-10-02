// Реализовать конкурентную запись данных в map.

package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	// мьютекс для избежания гонки данных, так как мапа не потокобезопасна
	mx := &sync.Mutex{}
	ch := make(chan int)
	m := make(map[int]string)

	wg.Add(10)

	// отправка в канал данных для дальнейшей записи в мапу
	go func() {
		defer close(ch)

		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	// запись данных в мапу с помощью нескольких воркеров
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()

			mx.Lock()
			v := <-ch
			m[v] = fmt.Sprintf("value%d", v)
			mx.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(m)
}
