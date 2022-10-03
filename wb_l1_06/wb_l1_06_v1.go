//Реализовать все возможные способы остановки выполнения горутины.

package main

import (
	"fmt"
	"sync"
)

// функция, вызыванная в горутине, остановится,
// как только канал, из которого она читает, будет закрыт
func printNum(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	for {
		select {
		case result, ok := <-ch:
			if ok {
				fmt.Println(result)
			}
			if !ok {
				fmt.Println("channel is closed")
				return
			}
		default:
		}
	}
}

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 100)

	wg.Add(2)
	go func() {
		defer wg.Done()
		defer close(ch)

		for i := 0; i < 100; i++ {
			ch <- i
		}
	}()
	go printNum(wg, ch)
	wg.Wait()
}
