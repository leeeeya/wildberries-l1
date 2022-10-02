// Реализовать постоянную запись данных в канал в главном потоке
// и набор из N воркеров для чтения и вывода данных из канала
// с возможностью выбора количества воркеров при старте.
// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров

package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"
	"time"
)

// читает и печатает данные из канала, завершается по сигналу Ctrl+C
func printVal(ctx context.Context, ch <-chan int, w int) {
	for {
		select {
		// после получения сигнала канал ctx.Done() закроется и горутина завершится
		case <-ctx.Done():
			return
		// пока действует контекст и в канале есть значения, рутина читает и печатает данные
		case data, ok := <-ch:
			if ok {
				fmt.Printf("Worker: %d, data: %d\n", w, data)
			}
		}
	}
}

func main() {
	// определение контекста, который заставит main-рутину завершиться по нажатию Ctrl+C
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT)

	ch := make(chan int)
	var workersNum int

	// считывание количества воркеров для вывода данных
	if _, err := fmt.Scan(&workersNum); err != nil {
		log.Fatal(err)
	}
	fmt.Println()

	for ; workersNum > 0; workersNum-- {
		go printVal(ctx, ch, workersNum)
	}

	// отправка данных в канал, пока действует контекст
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			close(ch)
			fmt.Println(ctx.Err())
			return
		default:
			ch <- i
			time.Sleep(time.Millisecond * 100)
		}
	}
}

// завершение горутин с помощью контекста - потокобезопасный способ, при котором все ресурсы возвращаются
// и утечек данных не происходит
