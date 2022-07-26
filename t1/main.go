/*
	Задание 1

	Дана структура Human (с произвольным набором полей и методов).
	Реализовать встраивание методов в структуре Action от родительской структуры Human
	(аналог наследования).
*/
package main

import "fmt"

// Структура родитель Human человек который умеет ходить и говорить
type human struct {
	age    int
	height int
	width  int
}

func (h human) walk() {
	fmt.Println("Идти")
}

func (h human) say() {
	fmt.Println("Говори")
}

type Action struct {
	human //встроили тип, теперь он и его поля доступны для вызова и записи
}

func main() {
	//Вызваем наследованный метод
	Action{}.say()
}
