package main

import "fmt"

/* Задание: Удалить i-ый элемент из слайса. */

func RemoveFromSlice[T any](arr []T, index int) []T {
	return append(arr[0:index], arr[index+1:]...)
}

func main() {
	arr := []int{0, 3, 5, 9, 11, 21, 25}
	fmt.Println(RemoveFromSlice[int](arr, 3))
}
