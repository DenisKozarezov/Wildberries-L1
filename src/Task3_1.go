package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/* Задание: Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений. */

var atomicSum int32

func PartialBlock(subArr []int, from int, to int) {
	for i := 0; i < len(subArr); i++ {
		subArr[i] *= subArr[i]
		atomic.AddInt32(&atomicSum, int32(subArr[i]))
	}
}

func main() {
	// Исходный массив
	arr := [5]int{2, 4, 6, 8, 10}
	arrLen := len(arr)

	// Разбить массив на одинаковые части (N), которые будут обрабатываться параллельно
	var ParallelCount int = 2
	wg := sync.WaitGroup{}
	wg.Add(ParallelCount)

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
			PartialBlock(subArr, from, to)
		}(from, to)
	}

	// Дождаться, пока все горутины закончат работу
	wg.Wait()

	fmt.Println(arr)
	fmt.Println("Sum of squares:", atomicSum)
}
