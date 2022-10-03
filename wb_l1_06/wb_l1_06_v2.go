//Реализовать все возможные способы остановки выполнения горутины.

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// функция, вызыванная в горутине, остановится,
// как только выполнится условие завершения контекста (в данном случае - таймаут)
func printNum2(ctx context.Context, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		case result, ok := <-ch:
			if ok {
				fmt.Println(result)
			}
		}
	}
}

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 100)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	wg.Add(2)
	go func() {
		defer wg.Done()
		defer close(ch)

		for i := 0; ; i++ {
			time.Sleep(time.Millisecond * 100)
			select {
			case <-ctx.Done():
				return
			case ch <- i:
			}
		}
	}()
	go printNum2(ctx, wg, ch)
	wg.Wait()
}
