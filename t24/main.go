package main

import (
	"fmt"
	"math"
)

/*
	Задание 24
	Разработать программу нахождения расстояния между двумя точками, которые
	представлены в виде структуры Point с инкапсулированными параметрами x,y и
	конструктором
*/

// Задаём структуру с не публичными полями
type point struct {
	x int
	y int
}

// Создаём интерфейс для удобства работы с другими ф-циями
type PointInterface interface {
	GetX() int
	GetY() int
}

// Зададим конструктор
func (p *point) SetXY(xx int, yy int) {
	p.x = xx
	p.y = yy
}

//Получаем данные по X
func (p point) GetX() int {
	return p.x
}

//Получаем данные об Y
func (p point) GetY() int {
	return p.y
}

//Для простоты реализую обычную функцию подсчитывающую расстояние
func analizer(p1 PointInterface, p2 PointInterface) float64 {

	//По школьной формуле расчитываем расстояние между точками
	//Корень суммы квадратов разностей координат
	d := math.Sqrt(math.Pow(float64(p1.GetX()-p2.GetX()), 2) + math.Pow(float64(p1.GetY()-p2.GetY()), 2))
	return d
}

func main() {
	//собираем первую точку
	point1 := point{}
	point1.SetXY(9, 9)

	//собираем вторую точку
	point2 := point{}
	point2.SetXY(10, 10)

	//считаем расстояние между точками
	fmt.Println(analizer(point1, point2))
}
