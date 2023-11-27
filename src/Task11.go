package main

import "fmt"

/* Задание: Реализовать пересечение двух неупорядоченных множеств. */

func intersection[T comparable](set1 []T, set2 []T) []T {
	result := []T{}
	hashmap := make(map[T]struct{})

	for i := 0; i < len(set1); i++ {
		hashmap[set1[i]] = struct{}{}
	}

	for i := 0; i < len(set2); i++ {
		if _, found := hashmap[set2[i]]; found {
			result = append(result, set2[i])
		}
	}

	return result
}

func main() {
	set1 := []int{-1, 5, 6, -7, 0, 13, -10}
	set2 := []int{-5, 6, 0, -13, 5, 1}

	fmt.Println(intersection[int](set1, set2)) // 5 0 6
}
