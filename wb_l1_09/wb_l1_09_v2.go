// Разработать конвейер чисел.
// Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
// после чего данные из второго канала должны выводиться в stdout

package main

import (
	"fmt"
	"sync"
)

// способ с применением WaitGroup
// данные выводятся в stdout из отдельной горутины
func main() {
	in := make(chan int, 100)
	out := make(chan int, 100)
	wg := sync.WaitGroup{}

	ar := make([]int, 0)
	for i := 0; i < 100; i++ {
		ar = append(ar, i)
	}

	wg.Add(3)
	// воркер для записи данных из массива в первый канал in
	go func() {
		defer wg.Done()

		for _, v := range ar {
			in <- v
		}
		close(in)
	}()

	// вокрер для чтения из первого канала и отправки результат операции x*2 во второй канал out
	go func() {
		defer wg.Done()
		defer close(out)

		for get := range in {
			out <- get * get
		}

	}()

	// воркер для вывода данных из канала out
	go func() {
		defer wg.Done()
		for pr := range out {
			fmt.Println(pr)
		}
	}()
	wg.Wait()
}
