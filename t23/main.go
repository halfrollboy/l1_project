package main

import (
	"fmt"
	"strings"
)

//Задание 23
//Удалить i-ый элемент из слайса.

func delItem(i int, slice []string) []string {
	//Поскольку slice передаётся по ссылке требуется создавать копию
	sl := make([]string, len(slice))
	//тут создаём копию в новую переменную
	copy(sl, slice)
	arr := append(sl[:i], sl[i+1:]...)
	return arr
}

func main() {
	//Задаём вхожную троку чтобы не вводить
	input := "Привет мой новый код NewCode!"
	//Создаём слайс
	splitMass := strings.Split(input, " ")
	//Проверяем подряд слова которые удаляться
	for i := range splitMass {
		fmt.Println("With %d delete item = ", i, delItem(i, splitMass))
	}
}
