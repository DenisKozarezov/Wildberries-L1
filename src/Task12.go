package main

import "fmt"

/* Задание: Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество. */

func FindUniqueSequence(arr []string) []string {
	hashmap := make(map[string]struct{})
	result := []string{}

	for i := 0; i < len(arr); i++ {
		if _, found := hashmap[arr[i]]; !found {
			hashmap[arr[i]] = struct{}{}
			result = append(result, arr[i])
		}
	}

	return result
}

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(FindUniqueSequence(strings))
}
