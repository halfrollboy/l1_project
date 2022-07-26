package main

import (
	"fmt"
	"strings"
)

/*
	Задание 26
	Разработать программу, которая проверяет, что все символы в строке
	уникальные (true — если уникальные, false etc). Функция проверки должна быть
	регистронезависимой.
*/

//Функция проверки символов на повторения
func checkCount(str string) bool {
	//Сразу в нижний регистр всё чтобы не зависеть от него
	str = strings.ToLower(str)
	//Пробегаемся циклом и считаем сколько повторений
	for _, val := range str {
		if strings.Count(str, string(val)) > 1 {
			return false
		}
	}
	return true
}

func main() {
	//Входные данные
	input := []string{
		"Мой новый код",
		"vjq67",
	}
	//Пробегаемся по элементам слайса и отправляем на проверку
	for _, str := range input {
		fmt.Println("Строка ", str, checkCount(str))
	}
}
