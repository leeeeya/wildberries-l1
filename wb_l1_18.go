package main

import (
	"fmt"
	"sync"
)

func increment(wg *sync.WaitGroup, mx *sync.Mutex, i *int) {
	defer wg.Done()
	for *i < 100 {
		mx.Lock()
		*i++
		mx.Unlock()
	}
}

func main() {
	mx := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	i := 0
	wg.Add(10)
	for j := 0; j < 10; j++ {
		go increment(wg, mx, &i)
	}
	wg.Wait()
	fmt.Println(i)
}