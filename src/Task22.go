package main

import (
	"fmt"
	"math/big"
)

/* Задание: Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20. */

func Multiply(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b) // Умножить
}

func Divide(a, b *big.Int) *big.Int {
	return new(big.Int).Div(a, b) // Делить
}

func Substract(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b) // Вычесть
}

func Add(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b) // Сложить
}

func main() {
	a := big.NewInt(1 << 61)
	b := big.NewInt(1 << 59)

	fmt.Println(Multiply(a, b))
	fmt.Println(Divide(a, b))
	fmt.Println(Substract(a, b))
	fmt.Println(Add(a, b))
}
