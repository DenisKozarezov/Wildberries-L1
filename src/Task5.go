package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

/* Задание: Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/

func WorkerFunc(workerId int, ch chan string) {
	for {
		if mes, opened := <-ch; opened {
			fmt.Printf("[WORKER %d]: %s\n", workerId, mes)
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

	fmt.Println("Timer expiration seconds:")
	var Seconds int
	_, err = fmt.Scanln(&Seconds)

	if err != nil {
		panic(err)
	}

	ch := make(chan string)

	// Запустить N воркеров для чтения данных из канала
	fmt.Printf("Running %d readers workers...", N)
	for i := 0; i < N; i++ {
		go WorkerFunc(i, ch)
	}

	// В другой горутине каждую секунду выбрасывать в канал рандомную 6-символьную строку
	go func() {
		fmt.Println("Running a writer worker...")
		for {
			time.Sleep(1 * time.Second)
			ch <- RandStringRunes(6) // Отправить в канал сообщение
		}
	}()

	time.Sleep(time.Duration(Seconds) * time.Second) // Остановить главную горутину на определенное время
	fmt.Println("Timer is expired...")
	os.Exit(0) // Выйти из программы
}
