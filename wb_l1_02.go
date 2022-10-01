// Конкурентно рассчитать значения квадратов чисел

package main

import (
	"fmt"
	"sync"
)

func main() {
	ar := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}

	// сообщаем wg, сколько рутин надо ожидать
	wg.Add(len(ar))
	for _, i := range ar {
		// каждое число обсчитывается в отдельной горутине и сразу выводится в stdout
		go func(n int) {
			// рутина сообщает о завершении своей работы
			defer wg.Done()

			fmt.Println(n * n)
		}(i)
	}
	//main-рутина ждёт завершения всех дочерних горутин
	wg.Wait()
}
