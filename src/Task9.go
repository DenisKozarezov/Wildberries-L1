package main

import (
	"fmt"
	"time"
)

/* Задание: Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.*/

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	arr := [...]int{1, 3, 5, 7, 9, 11}

	go func() {
		fmt.Println(arr)
		for _, el := range arr {
			fmt.Println("[CHANNEL 1]:", el)
			ch1 <- el // Записать в первый канал число из массива
		}
		fmt.Println("Channel 1 is closed")
		close(ch1) // Если поток данных иссяк, закрыть первый канал
	}()

	go func() {
		for {
			if el, opened := <-ch1; opened { // Чтение из 1-го канала
				fmt.Println("[CHANNEL 2]:", el)
				ch2 <- el * 2 // Записать во второй канал x * 2
			} else {
				fmt.Println("Channel 2 is closed")
				close(ch2) // Если поток данных иссяк, закрыть второй канал
				return
			}
		}
	}()

	for i := 0; i < len(arr); i++ {
		fmt.Println("[MAIN]: x * 2 =", <-ch2) // Чтение из 2-го канала
	}

	time.Sleep(100 * time.Millisecond)
}
