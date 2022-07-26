package main

import (
	"fmt"
	"sync"
)

/*
Задание 2

Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

type sum struct {
	mass []int
	itog int
	mu   sync.RWMutex
}

func setMass(mass []int) *sum {
	return &sum{mass: mass}
}

func (s *sum) Add(val int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	//Можно было сдеоать через Pow но с int попроще
	s.itog = s.itog + (val * val)

}

func main() {
	mass := []int{2, 4, 6, 8, 10}
	summer := setMass(mass)
	//Создаём группу ождания горутин
	var wg sync.WaitGroup

	for _, val := range mass {
		wg.Add(1)

		//На каждом цикле создаём горутину под значение
		go func(i int) {
			summer.Add(i)
			// на Done Горутина всё
			// Если не закрыть поймаем deadlock
			wg.Done()
			//Передаём в функцию иначем можем не успеть прочитать val до смены во внешней
			//области видимости
		}(val)
	}
	//Ждём пока все горутины отработают
	wg.Wait()
	fmt.Printf("Сумма квадратов %d\n", summer.itog)
}
