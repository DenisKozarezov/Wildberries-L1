package main

import (
	"fmt"
	"math/rand"
	"time"
)

/* Задание: Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров,
которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества
воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/

func WorkerFunc(workerId int, ch chan string) {
	for {
		if mes, opened := <-ch; opened {
			fmt.Printf("[WORKER %d]: %s\n", workerId, mes)
		} else {
			fmt.Printf("Stopping [WORKER %d]...\n", workerId)
			break
		}
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// Код для генерации случайной строки нужной длины
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func main() {
	// Считать желаемое количество воркеров для чтения данных из канала
	fmt.Println("Enter Go-workers count:")
	var N int
	_, err := fmt.Scanln(&N)
	if err != nil {
		panic(err)
	}

	ch := make(chan string)
	defer close(ch)

	// Запустить N воркеров для чтения данных из канала
	fmt.Printf("Running %d readers workers...", N)
	for i := 0; i < N; i++ {
		go WorkerFunc(i, ch)
	}

	// В главной горутине каждую секунду выбрасывать в канал рандомную 6-символьную строку
	fmt.Println("Running a writer worker...")
	for {
		time.Sleep(1 * time.Second)
		ch <- RandStringRunes(6)
	}
}
