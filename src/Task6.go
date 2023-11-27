package main

import (
	"context"
	"fmt"
	"time"
)

/* Задание: Реализовать все возможные способы остановки выполнения горутины. */

func Case1() {
	ch := make(chan int)

	go func(ch chan int) {
		var n int = 25
		for {
			select {
			case ch <- n: // При чтении из канала
				n++
				fmt.Println("Increment")
			case <-ch: // При записи в канал
				fmt.Println("1-ая горутина завершила работу")
				close(ch)
				return
			}
		}
	}(ch)

	time.Sleep(2 * time.Second)

	ch <- 0 // Триггер для закрытия канала и остановки горутины

	time.Sleep(1 * time.Second)
}

func Case2() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done(): // Была вызвана функция cancel()
				fmt.Println("2-ая горутина завершила работу")
				return
			}
		}
	}(ctx)

	time.Sleep(2 * time.Second)

	cancel() // Остановка горутины в явном виде

	time.Sleep(1 * time.Second)
}

func main() {
	Case1()

	Case2()
}
