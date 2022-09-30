package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func senValue(ctx context.Context, wg *sync.WaitGroup, ch chan int) {
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

}

func printValue(ctx context.Context, wg *sync.WaitGroup, ch chan int) {
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
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	var sec int

	if _, err := fmt.Scan(&sec); err != nil {
		return
	}
	fmt.Println()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(sec))
	defer cancel()
	wg.Add(2)
	go senValue(ctx, wg, ch)
	go printValue(ctx, wg, ch)
	wg.Wait()
}
