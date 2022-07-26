package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
	Задание 18

	Реализовать структуру-счетчик, которая будет инкрементироваться в
	конкурентной среде. По завершению программа должна выводить итоговое
	значение счетчика.

*/

type CounterInterface interface {
	Add()
	Read()
}

type Counter struct {
	counter     int
	counterLock sync.RWMutex
}

func getCounter() *Counter {
	return &Counter{counter: 0}
}

//Через мьютекс добавляем только один элемент
func (s *Counter) Add() {
	s.counterLock.Lock()
	s.counter++
	defer s.counterLock.Unlock()
}

//Читаем и соответсвенно ставим мьютекс на чтение
func (s *Counter) Read() int {
	s.counterLock.RLock()
	defer s.counterLock.RUnlock()
	return s.counter
}

func main() {
	//Получаем наш счётчик
	counter := getCounter()
	var ops uint64

	//Создаём группу горутин
	var wg sync.WaitGroup

	//Стартуем 50 горутин
	for i := 0; i < 50; i++ {
		wg.Add(1)
		//Чтобы подтвердить что наш Counter работает добавим atomic
		go func() {
			for c := 0; c < 1000; c++ {
				counter.Add()
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}
	//Ждём группу горутин
	wg.Wait()
	//Сравниваем данные
	fmt.Println("Counter:", counter.Read())
	fmt.Println("ops:", ops)
}
