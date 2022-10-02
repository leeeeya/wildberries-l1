// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде

package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	i int
	*sync.Mutex
}

// NewCounter конструктор
func (c *Counter) NewCounter() {
	c.i = 0
	c.Mutex = &sync.Mutex{}
}

// Increment увеличивает счётчик до заданного числа max
func (c *Counter) Increment(max int) {
	for c.i < max {
		c.Lock()
		c.i++
		c.Unlock()
	}
}

func main() {
	wg := &sync.WaitGroup{}

	c := new(Counter)
	c.NewCounter()

	wg.Add(10)

	// запуск воркеров для конкурентной инкрементации счётчика
	for worker := 0; worker < 10; worker++ {
		go func() {
			defer wg.Done()

			c.Increment(100)
		}()
	}
	wg.Wait()
	fmt.Println(c.i)
}
