package main

import "fmt"

/* Задание: Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской
структуры Human (аналог наследования). */

type Human struct {
	Action
}

type Action struct {
}

func (self *Action) Method1() {
	fmt.Println("Method1()")
}

func (self *Action) Method2() {
	fmt.Println("Method2()")
}

func (self *Action) Method3() {
	fmt.Println("Method3()")
}

func main() {
	human := Human{}
	human.Method1()
	human.Method2()
	human.Method3()
}
