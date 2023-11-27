package main

import (
	"fmt"
	"strings"
)

func CheckUnique(str string) bool {
	hashtable := make(map[rune]struct{})

	str = strings.ToUpper(str)
	for _, sym := range str {
		if _, found := hashtable[sym]; !found {
			hashtable[sym] = struct{}{}
		} else {
			return false
		}
	}
	return true
}

func main() {
	str1 := "abcd"
	fmt.Println(CheckUnique(str1)) // true

	str2 := "abCdefAaf"
	fmt.Println(CheckUnique(str2)) // false

	str3 := "aabcd"                // both 'a' are Latin
	fmt.Println(CheckUnique(str3)) // false

	str4 := "аabcd"                // first 'a' is Cyrilic, second 'a' is Latin
	fmt.Println(CheckUnique(str4)) // true

	str5 := "Aabcd"
	fmt.Println(CheckUnique(str5))
}
