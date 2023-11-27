package main

import "fmt"

// Задание: Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.

// Возвращает новую строку. Копирует символы.
// Returns a new allocated string. Bad for huge strings, can't handle Cyrilic alphabet.
func ReverseString1(str string) string {
	var result string
	for i := len(str) - 1; i >= 0; i-- {
		result += string(str[i])
	}
	return result
}

// Не аллоцирует новую строку. Использует лишь два индексатора по 4 байт и указатель на слайс.
// No allocated string in return. Just 2 integer indexers and pointer to slice.
func ReverseString2(str []rune) {
	left := 0
	right := len(str) - 1
	for left < right {
		str[left], str[right] = str[right], str[left]
		left++
		right--
	}
}

func main() {
	var str string = "главрыба"
	fmt.Println(ReverseString1(str))

	var runes []rune = []rune("главрыба")
	ReverseString2(runes)
	fmt.Println(string(runes))
}
