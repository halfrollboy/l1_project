package main

import (
	"fmt"
)

// Задача 17
// Требуется реализовать бинраный поиск

// Функция реализующая бинарный поиск
// При неудаче возвращает -1 при удаче значение
func binarySearch(arr []int, x int) int {
	i := 0
	j := len(arr)
	for i != j {
		//ищем среднее
		var sred = (i + j) / 2
		//если равно то возвращаем
		if x == arr[sred] {
			return sred
		}
		//если меньше то берём отрезок
		if x < arr[sred] {
			j = sred
		} else {
			i = sred + 1
		}
	}
	return -1
}

func main() {
	items := []int{2, 3, 5, 7, 11, 13, 17}
	//Замеряем на различных данных
	fmt.Println(binarySearch(items, 1))
	// -1
	fmt.Println(binarySearch(items, 7))
	// 3
	fmt.Println(binarySearch(items, 19))
	// -1
}
