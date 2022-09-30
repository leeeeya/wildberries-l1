package main

import "fmt"

func main() {
	in := make(chan int, 100)
	out := make(chan int, 100)

	ar := make([]int, 0)
	for i := 0; i < 100; i++ {
		ar = append(ar, i)
	}

	go func() {
		for _, v := range ar {
			in <- v
		}
		close(in)
	}()

	go func() {
		for get := range in {
			out <- get * get
		}
		close(out)
	}()

	for pr := range out {
		fmt.Println(pr)
	}
}
