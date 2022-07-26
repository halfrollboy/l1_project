package main

import (
	"fmt"
	"strings"
)

/*
	Задание 20
	Разработать программу, которая переворачивает слова в строке.
	Пример: «snow dog sun — sun dog snow».
*/

func main() {
	//Задаём вхожную троку чтобы не вводить
	input := "Привет мой новый код NewCode!"
	//Делим строку на []string
	splitMass := strings.Split(input, " ")

	//Выполняем переставление внутри строки
	for i, j := 0, len(splitMass)-1; i < j; i, j = i+1, j-1 {
		splitMass[i], splitMass[j] = splitMass[j], splitMass[i]
	}

	fmt.Println(splitMass)
}
