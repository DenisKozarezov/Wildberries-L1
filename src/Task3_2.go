package main

import (
	"fmt"
	"sync"
)

/* Задание: Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений. */

func PartialBlock(subArr []int, from int, to int, ch chan int) {
	for i := 0; i < len(subArr); i++ {
		ch <- subArr[i] * subArr[i]
	}
}

func main() {
	// Исходный массив
	arr := [5]int{2, 4, 6, 8, 10}
	arrLen := len(arr)
	var sum int = 0

	// Разбить массив на одинаковые части (N), которые будут обрабатываться параллельно
	var ParallelCount int = 2
	wg := sync.WaitGroup{}
	wg.Add(ParallelCount)

	// Канал для получения значения квадрата
	ch := make(chan int, ParallelCount)

	// Создать N горутин для конкурентного выполнения
	part := arrLen / ParallelCount
	for i := 0; i < ParallelCount; i++ {

		// Рассчитать индексы каждой отдельной части массива
		from := i * part
		to := from + part
		if i == ParallelCount-1 {
			to++
		}
		subArr := arr[from:to]

		// Пусть каждая часть считается в отдельной горутине
		go func(from int, to int) {
			defer wg.Done()
			PartialBlock(subArr, from, to, ch)
		}(from, to)
	}

	// Дополнительная горутина, которая слушает канал и суммирует квадраты
	go func() {
		for {
			if mes, opened := <-ch; opened {
				sum += mes
			} else {
				return
			}
		}
	}()

	// Дождаться, пока основные горутины закончат работу
	wg.Wait()
	close(ch)

	fmt.Println(arr)
	fmt.Println("Sum of squares:", sum)
}
