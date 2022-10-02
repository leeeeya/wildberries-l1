// Разработать конвейер чисел.
// Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
// после чего данные из второго канала должны выводиться в stdout

package main

import "fmt"

// способ без WaitGroup, для синхронизации рутин используются только каналы.
// данные выводятся в stdout из main-рутины
func main() {
	in := make(chan int, 100)
	out := make(chan int, 100)

	ar := make([]int, 0)
	for i := 0; i < 100; i++ {
		ar = append(ar, i)
	}

	// воркер для записи данных из массива в первый канал in
	go func() {
		for _, v := range ar {
			in <- v
		}
		close(in)
	}()

	// вокрер для чтения из первого канала и отправки результат операции x*2 во второй канал out
	go func() {
		for get := range in {
			out <- get * get
		}
		close(out)
	}()

	// вывод данных из канала out
	for pr := range out {
		fmt.Println(pr)
	}
}
