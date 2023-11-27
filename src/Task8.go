package main

import "fmt"

/* Задание: Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0. */

func SetBit(value int64, index int64, bit bool) int64 {
	mask := (int64)(1) << index
	if bit {
		return value | mask
	} else {
		return value &^ mask
	}
}

func main() {
	var x int64 = 128
	// 1000 0000 = 128
	// 1000 0001 = 129

	x = SetBit(x, 0, true)
	fmt.Println(x) // 129

	x = 15
	// 0000 1111 = 15
	// 0000 1001 = 9

	x = SetBit(x, 1, false)
	x = SetBit(x, 2, false)
	fmt.Println(x) // 9
}
