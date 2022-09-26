package main

import (
	"fmt"
	"time"
)

func mySleep(d time.Duration) {
	start := time.Now().UnixMicro()

	for {
		now := time.Now().UnixMicro()
		s := now - start
		if s < d.Microseconds() {
			time.Sleep(10 * time.Microsecond)
			continue
		}
		break
	}
}

func main() {
	fmt.Println("Start")
	mySleep(3 * time.Second)
	fmt.Println("Stop")
}
