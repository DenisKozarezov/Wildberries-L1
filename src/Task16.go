package main

import (
	"fmt"
	"time"
)

/* Задание: Реализовать быструю сортировку массива (quicksort) встроенными методами языка. */

func partition[T ~int](arr []T, low int, high int) ([]T, int) {
	middle := (low + high) / 2                      // находим индекс числа находящегося в середине слайса
	arr[high], arr[middle] = arr[middle], arr[high] // меняем средний и последний элемент местами

	pivot := arr[high] // Выбираем опорный элемент (последний элемент)
	i := low           // Начинаем обработку с левого элемента массива
	for j := low; j < high; j++ {
		if arr[j] < pivot { // Если элемент меньше опорного, поменять местами
			arr[i], arr[j] = arr[j], arr[i]
			i++ // Сдвинуть индекс опорного элемента
		}
	}
	arr[i], arr[high] = arr[high], arr[i] // Заменить последний элемент с опорным
	return arr, i
}

func quicksort[T ~int](arr []T, low int, high int) {
	if low < high {
		arr, index := partition(arr, low, high)
		quicksort[T](arr, low, index-1)
		quicksort[T](arr, index+1, high)
	}
}

func main() {
	var arr []int
	for i := 100_000; i > 0; i-- {
		arr = append(arr, i)
	}

	t := time.Now()
	quicksort[int](arr, 0, len(arr)-1)
	elapsedTime := time.Since(t)
	fmt.Printf("Elapsed time = %s", elapsedTime)
}
