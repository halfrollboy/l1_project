package main

import "fmt"

/*
	Задание 12

	Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
	собственное множество.
*/

type mySet struct {
	mass []string
	set  map[string]int
}

//Типа конструктор
func createMySet(mass []string) *mySet {
	return &mySet{mass: mass}
}

//Получить нащ SET
func (m *mySet) GetSet() []string {
	var keys []string
	m.set = make(map[string]int)
	for _, item := range m.mass {
		// Проверяем есть ли такое значение
		if _, found := m.set[item]; found {
			m.set[item] = m.set[item] + 1
		} else {
			m.set[item] = 1
		}
	}
	for key, _ := range m.set {
		keys = append(keys, key)
	}
	fmt.Println("Статистика вашего сета: ")
	fmt.Println(m.set)
	return keys
}

//Отладочная ф-ция для проверки
func (m *mySet) GetStat() map[string]int {
	return m.set
}

//Cоздать самый лучий SET из вашего массива
func createSet(mass []string) []string {
	setter := createMySet(mass)
	return setter.GetSet()
}

func main() {
	//Задаём изначальный массив
	massiv := []string{"cat", "cat", "dog", "cat", "tree"}

	set := createSet(massiv)
	fmt.Println("Ваш сет: ")
	fmt.Println(set)
}
