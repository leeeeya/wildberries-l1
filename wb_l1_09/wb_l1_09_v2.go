package main

import (
	"fmt"
	"sync"
)

func main() {
	in := make(chan int, 100)
	out := make(chan int, 100)
	wg := sync.WaitGroup{}

	ar := make([]int, 0)
	for i := 0; i < 100; i++ {
		ar = append(ar, i)
	}

	wg.Add(3)
	go func() {
		defer wg.Done()

		for _, v := range ar {
			in <- v
		}
		close(in)
	}()

	go func() {
		defer wg.Done()
		defer close(out)

		for get := range in {
			out <- get * get
		}

	}()

	go func() {
		defer wg.Done()
		for pr := range out {
			fmt.Println(pr)
		}
	}()
	wg.Wait()
}
