package main

import (
	"fmt"
	"reflect"
)

/*
	Задание 14
	Разработать программу, которая в рантайме способна определить тип
	переменной: int, string, bool, channel из переменной типа interface{}.
*/

//Создаём ф-цию для определения
func getMyType(val interface{}) {
	xType := reflect.TypeOf(val)
	fmt.Println(xType) // "[]int [1 2 3]"
}

func main() {
	//Вводим любой тип перемнной
	getMyType([]int{1, 2, 3})
}
