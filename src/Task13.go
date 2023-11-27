package main

import "fmt"

/* Задание: Поменять местами два числа без создания временной переменной. */

func Case1(value1 float32, value2 float32) (float32, float32) {
	x, y := value2, value1
	return x, y
}

func Case2(value1 float32, value2 float32) (float32, float32) {
	return value2, value1
}

func main() {
	// Инициализировать числа value1, value2
	var value1 float32 = 5
	var value2 float32 = 7

	fmt.Println(Case1(value1, value2))
	fmt.Println(Case2(value1, value2))

	// Case 3
	value1, value2 = value2, value1
	fmt.Println(value1, value2)
}
