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
	fmt.Println(v, m[v])
	mx.Unlock()

}

func sendValue(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch)

	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func main() {
	m := make(map[int]string)
	wg := &sync.WaitGroup{}
	mx := &sync.Mutex{}
	ch := make(chan int)

	wg.Add(11)
	go sendValue(ch, wg)
	for i := 0; i < 10; i++ {
		go setValue(wg, m, ch, mx)
	}
	wg.Wait()
	fmt.Println(m)
}
