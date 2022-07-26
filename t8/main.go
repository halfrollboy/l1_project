package main

import "fmt"

/*
Задание 8

Установить i бит в 1 или 0
*/

func getBit(i int) {
	//Создаём маску и просто на нужный нам бит ставим 1
	var b int = 1 << i
	//Складываем маску с числом
	fmt.Println(5 ^ b)
}

func main() {
	//Для примера берём 5 = 101
	//Требуется поменять первый бит на 1, чтобы вышло 7
	//он будет 1 потому что есть ещё 0 бит
	getBit(1)
	//Проверяем что получилось 7
	fmt.Println(5 ^ 2)
}