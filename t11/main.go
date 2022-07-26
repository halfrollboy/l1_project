package main

import (
	"fmt"
)

/*
	Задание 11
	Реализовать пересечение двух неупорядоченных множеств
*/

func getIntersection(a []string, b []string) []string {
	m := make(map[string]int)

	//Пишем в мапу если нет - создаём, если есть добавляем 1
	for _, k := range a {
		if _, ok := m[k]; !ok {
			m[k] = 1
		}
	}
	//Если находим в мапе то прибавляем 1
	for _, k := range b {
		if _, ok := m[k]; ok {
			m[k] += 1
		}
	}
	//Инициализируем
	result := []string{}
	//Проходимся и проверяем если в мапе больше 2
	for k, v := range m {
		if v == 2 {
			result = append(result, k)
		}
	}
	return result
}

func main() {
	//Задаём изначальные массивы
	a := []string{"a", "b", "c", "d"}
	b := []string{"c", "d", "e", "f"}
	//Получаем пересечение
	fmt.Println(getIntersection(a, b))
}
