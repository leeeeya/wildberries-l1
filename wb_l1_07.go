package main

import (
	"fmt"
	"sync"
)

func setValue(m map[int]string, ch <-chan int, wg *sync.WaitGroup) {
	mx := sync.Mutex{}

	defer wg.Done()
	mx.Lock()
	v := <-ch
	m[v] = fmt.Sprintf("value%d", v)
	mx.Unlock()

}

func sendValue(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func main() {
	m := make(map[int]string)
	wg := &sync.WaitGroup{}
	ch := make(chan int, 5)

	wg.Add(10)
	go sendValue(ch, wg)
	for i := 0; i < 10; i++ {
		go func(ch <-chan int, wg *sync.WaitGroup) {
			mx := sync.Mutex{}

			defer wg.Done()
			mx.Lock()
			v := <-ch
			m[v] = fmt.Sprintf("value%d", v)
			mx.Unlock()
		}(ch, wg)
	}
	wg.Wait()
	fmt.Println(m)
}
