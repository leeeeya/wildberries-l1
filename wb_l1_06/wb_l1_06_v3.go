//Реализовать все возможные способы остановки выполнения горутины.

package main

import (
	"fmt"
	"time"
)

// горутина остановится, как только в канал stop отправится значение
func main() {
	stop := make(chan bool)
	go func() {
		i := 0
		for {
			select {
			case <-stop:
				fmt.Println("goroutine is stopped")
				return
			default:
				fmt.Println(i)
				time.Sleep(100 * time.Millisecond)
			}
			i++
		}
	}()
	time.Sleep(2 * time.Second)
	stop <- true
	close(stop)
}
