package main

import (
	"fmt"
)

/*
	Задание 16
	Реализовать быструю сортировку массива (quicksort) встроенными методами
	языка.
*/

func quick(mass []int) []int {
	if len(mass) < 2 {
		return mass
	}
	//Можно брать первый или последний элемент, отличий от random не будет
	//Но Random будет работать эффективнее в заранее упорядоченных строках
	// random := rand.Intn(len(mass)-1) + 1
	random := 0
	var men []int
	var bol []int

	for _, k := range mass {
		if k < mass[random] {
			men = append(men, k)
		} else if k > mass[random] {
			bol = append(bol, k)
		}
	}

	//Рекурсивно вызываем эту же функцию
	arr := append(append(quick(men), mass[random]), quick(bol)...)
	return arr
}

func main() {
	//Задаём изначальный массив
	mass := []int{1, 4, 56, 3, 6, 4, 34, 6, 3, 2}
	fmt.Println(quick(mass))
}
