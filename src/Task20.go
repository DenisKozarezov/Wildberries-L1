package main

import (
	"fmt"
	"strings"
)

/* Задание: Разработать программу, которая переворачивает слова в строке. Пример: «snow dog sun — sun dog snow». */

func ReverseWords(str string) []string {
	words := strings.Split(str, " ") // Разбить строку на слова при помощи разделителя с символом пробела

	// Перевернуть массив со словами
	left := 0
	right := len(words) - 1
	for left < right {
		words[left], words[right] = words[right], words[left]
		left++
		right--
	}
	return words
}

func main() {
	fmt.Println(ReverseWords("snow dog sun"))
}
