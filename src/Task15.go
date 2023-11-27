package main

import (
	"fmt"
	"os"
)

/* Задание: К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации. */

func createHugeString(length int, file *os.File) {
	for i := 0; i < length; i++ {
		_, err := file.WriteString("0")
		if err != nil {
			panic(fmt.Errorf("Could not write in file!", err))
		}
	}
}

var justString string

func someFunc() {
	file, err := os.Open("myfile.txt")
	if err != nil {
		panic(fmt.Errorf("Could not open file!", err))
	}

	createHugeString(1<<10, file) // Лучше записать огромные строки в файл, чем держать в оперативной памяти

	buffer := make([]byte, 100)
	bytes, err := file.Read(buffer)
	justString = string(bytes)
}

func main() {
	someFunc()
}
