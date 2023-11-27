package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

/* Задание: Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика. */

type ThreadSafeCounter struct {
	Value int32
}

func (self *ThreadSafeCounter) Increment() {
	atomic.AddInt32(&self.Value, 1) // Использовать атомарную переменную для инкремента счетчика
}

func (self *ThreadSafeCounter) Decrement() {
	atomic.AddInt32(&self.Value, -1) // Использовать атомарную переменную для декремента счетчика
}

func main() {
	counter := ThreadSafeCounter{}

	go func() {
		for i := 0; i < 200; i++ {
			counter.Increment()
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			counter.Decrement()
		}
	}()

	time.Sleep(2 * time.Second)

	fmt.Println(counter.Value) // 100
}
