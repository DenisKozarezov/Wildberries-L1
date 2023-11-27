package main

import (
	"fmt"
	"time"
)

/* Задание: Реализовать собственную функцию sleep. */

func Sleep(duration time.Duration) {
	<-time.After(duration)
}

func main() {
	start := time.Now()
	fmt.Println("Before sleep:", start)
	Sleep(2 * time.Second)
	fmt.Println("After sleep:", time.Now())
	fmt.Println("Elapsed time:", time.Since(start))
}
