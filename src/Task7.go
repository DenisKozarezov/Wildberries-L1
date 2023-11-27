package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

/* Задание: Реализовать конкурентную запись данных в map.*/

type ThreadSafeMap struct {
	Map   map[int]string
	Mutex sync.Mutex
}

func (self *ThreadSafeMap) New() {
	self.Map = make(map[int]string)
}

func (self *ThreadSafeMap) Add(key int, value string) {
	self.Mutex.Lock()         // Закрывает критическую секцию, чтобы избежать состояния гонки
	defer self.Mutex.Unlock() // Открывает критическую секцию для остальных горутин

	self.Map[key] = value
}

func (self *ThreadSafeMap) Get(key int) (string, bool) {
	self.Mutex.Lock()         // Закрывает критическую секцию на случай, если другая горутина записывает
	defer self.Mutex.Unlock() // Открывает критическую секцию для остальных горутин

	val, ok := self.Map[key]
	return val, ok
}

func main() {
	threadSafeMap := ThreadSafeMap{}
	threadSafeMap.New() // Инициализация структуры

	go func() {
		for i := 0; i < 100; i++ {
			threadSafeMap.Add(i, "string "+strconv.Itoa(i))
		}
	}()

	go func() {
		for i := 100; i < 200; i++ {
			threadSafeMap.Add(i, "string "+strconv.Itoa(i))
		}
	}()

	go func() {
		for i := 0; i < 200; i++ {
			value, _ := threadSafeMap.Get(i)
			fmt.Println(value)
		}
	}()

	time.Sleep(100 * time.Millisecond)
}
