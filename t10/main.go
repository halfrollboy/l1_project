package main

import (
	"fmt"
	"math"
)

/*
Задание 10

	Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0,
	15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10
	градусов. Последовательность в подмножноствах не важна.

	Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/

func getTermometr(mass []float64) map[int][]float64 {
	m := make(map[int][]float64)
	for _, val := range mass {

		//Берём общее до 100
		for i := 10; i <= 100; i = i + 10 {
			//Проверяем вхождение в диапазон квадрата
			if math.Pow(math.Mod(val, float64(i)), 2) < math.Pow(math.Mod(float64(i+10), float64(i)), 2) && math.Pow(math.Mod(val, float64(i)), 2) > math.Pow(math.Mod(float64(i), float64(i)), 2) {
				m[i] = append(m[i], val)
			}
		}
	}
	return m
}

func main() {
	mass := []float64{25.4, -27.0, 13.0, 19.0,
		15.5, 24.5, -21.0, 32.5}
	fmt.Println(getTermometr(mass))
}
