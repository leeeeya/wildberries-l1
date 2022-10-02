// Реализовать собственную функцию sleep

package main

import (
	"fmt"
	"time"
)

// работает для d >= 1 * time.Millisecond
func mySleep(d time.Duration) {
	// начало отсчёта
	start := time.Now().UnixMicro()

	for {
		// настоящее время
		now := time.Now().UnixMicro()
		s := now - start

		// пока разница между стартом отсчёта и настоящим врменем не приблизится к заданному промежутку времени d,
		// программа засыпает на короткие промежутки времени
		if s < d.Microseconds() {
			time.Sleep(10 * time.Microsecond)
			continue
		}
		break
	}
}

func main() {
	fmt.Println("Start:", time.Now().UnixMilli())
	mySleep(13 * time.Millisecond)
	fmt.Println("Stop:", time.Now().UnixMilli())
}
