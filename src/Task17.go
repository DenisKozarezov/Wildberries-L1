package main

import (
	"fmt"
	"time"
)

/* Задание: Реализовать бинарный поиск встроенными методами языка. */

func BinarySearch(arr []int, value int) int {
	arrLen := len(arr)
	first := 0
	last := arrLen - 1
	for first <= last {
		middle := (first + last) / 2
		if arr[middle] == value {
			return middle
		} else if arr[middle] < value {
			first = middle + 1
		} else {
			last = middle - 1
		}
	}
	return -1
}

func main() {
	var arr []int
	for i := 0; i < 1_000_000; i += 5 {
		arr = append(arr, i)
	}

	start := time.Now()
	index := BinarySearch(arr, 65000)
	fmt.Printf("Elapsed time = %s\n", time.Since(start))
	fmt.Println("Index = ", index)
}
