package main

import (
	"fmt"
	"sync"
	"time"
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
		//defer close(in)
		for _, v := range ar {
			in <- v
		}
	}()
	go func() {
		defer wg.Done()
		get := <-in
		fmt.Println(get)
		out <- get * get
	}()
	go func() {
		defer wg.Done()
		fmt.Println(<-out)
	}()
	wg.Wait()
	time.Sleep(5 * time.Second)
}
