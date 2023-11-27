package main

import (
	"fmt"
	"math"
)

/* Задание: Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.

Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/

func GetInterval(value float64, h int) (int, int) {
	lesserValue := math.Floor(value)
	intervalNumber := (int)(lesserValue) / 10

	var left, right int
	if value >= 0 {
		left = intervalNumber * 10
		right = left + h
	} else {
		right = intervalNumber * 10
		left = right - h
	}
	return left, right
}

const h int = 10

type Category = []float64

func main() {
	arr := Category{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	categories := make(map[int]Category)

	for _, el := range arr {
		left, _ := GetInterval(el, h)
		if _, ok := categories[left]; ok {
			categories[left] = append(categories[left], el)
		} else {
			categories[left] = Category{el}
		}
	}

	fmt.Println(categories)
}
