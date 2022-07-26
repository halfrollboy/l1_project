package main

import (
	"fmt"
	"math/big"
)

// Задание 22
/*
Разработать программу, которая перемножает, делит, складывает, вычитает две
числовых переменных a,b, значение которых > 2^20.
*/

//Если честно возможно я немного не уловил смысл задания,
//он на то чтобы вспосмнить типы или надо было именно написать программу с вводом из консоли итд
func main() {

	// The following prime values were taken from here: https://eips.ethereum.org/EIPS/eip-197
	var prime1, _ = new(big.Int).SetString("21888242871839275222246405745257275088548364400416034343698204186575808495617", 10)
	var prime2, _ = new(big.Int).SetString("21888242871839275222246405745257275088696311157297823662689037894645226208583", 10)

	var mult = new(big.Int)
	mult.Mul(prime1, prime2)
	var mins = new(big.Int)
	mins.Sub(prime1, prime2)
	var adder = new(big.Int)
	adder.Add(prime1, prime2)
	var div = new(big.Int)
	div.Div(prime1, prime2)

	fmt.Printf("Mul: %v\n", mult)
	fmt.Printf("Mins: %v\n", mins)
	fmt.Printf("addr: %v\n", adder)
	fmt.Printf("div: %v\n", div)
}
