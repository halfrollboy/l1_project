package main

import (
	"fmt"
	"strings"
)

/*
	Задание 15

	К каким негативным последствиям может привести данный фрагмент кода, и как
	это исправить? Приведите корректный пример реализации.
*/

var justString string

//Для удобства подсчёта возвращаю строку букв
func createHugeString(i float32) string {
	fmt.Println(i)

	return strings.Repeat("a", int(i))
}

func someFunc() {
	//Тут используется оператор побитового сдвига
	// Задавать так не особо очевидно
	// Может возникнуть ошибка из-за сдвига
	v := createHugeString(1 << 10)
	g := createHugeString(1 << 10)
	justString = v[:100]
	justString = g[:20]
}

// func someGoodFunc() {
// 	v := createGodeString)

// }

func main() {
	someFunc()
}
