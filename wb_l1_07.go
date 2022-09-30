package main

import (
	"fmt"
	"sync"
)

func setValue(wg *sync.WaitGroup, m map[int]string, ch <-chan int, mx *sync.Mutex) {
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
	mx := &sync.Mutex{}
	ch := make(chan int)
	defer close(ch)

	wg.Add(11)
	go sendValue(ch, wg)
	for i := 0; i < 10; i++ {
		go setValue(wg, m, ch, mx)
	}
	wg.Wait()
	fmt.Println(m)
}
