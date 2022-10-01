// Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться.

package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

// отправляет значения в канал ch, пока не отменится контекст
func sendVal(ctx context.Context, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	defer close(ch)

	for i := 0; ; i++ {
		time.Sleep(time.Millisecond * 100) // Sleep для менее быстрого перебора чисел
		select {
		// по истечении таймаута канал ctx.Done() закроется и горутина завершится
		case <-ctx.Done():
			return
		// отправка чисел в канал ch
		case ch <- i:

		}
	}

}

// читает и распечатывает значения из канала ch, пока не отменится контекст
func printValue(ctx context.Context, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	for {
		select {
		// по истечении таймаута горутина завершится и выведется сообщение об отмене контекста
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		// чтение и вывод значений
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

	// считывание количества секунд, спустя которое программа должна завершиться
	if _, err := fmt.Scan(&sec); err != nil {
		log.Fatalln(err)
	}
	fmt.Println()

	// определение контекста, который заставить main-рутину завершиться по заданному таймауту
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(sec))
	// cancel() отменит контекст после того, как он завершит свою работу
	defer cancel()

	wg.Add(2)
	go sendVal(ctx, wg, ch)
	go printValue(ctx, wg, ch)
	wg.Wait()
}
