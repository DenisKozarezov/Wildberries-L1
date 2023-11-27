package main

import (
	"fmt"
	"reflect"
)

/* Задание: Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}. */

func GetUnderlyingType(obj any) reflect.Type {
	return reflect.TypeOf(obj)
}

func main() {
	fmt.Println(GetUnderlyingType(true))
	fmt.Println(GetUnderlyingType(5))

	ch := make(chan int)
	fmt.Println(GetUnderlyingType(ch))
}
