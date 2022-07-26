package main

import (
	"fmt"
	"sync"
)

/*
	Задание 7

	Реализовать конкурентную запись данных в map.
*/

func main() {
	//Существует прекрасный sync.Map который избавляет нас от cach condition
	var counter sync.Map

	//Создаём группу горутин
	var wg sync.WaitGroup

	//Стартуем 50 горутин
	for i := 0; i < 50; i++ {
		wg.Add(1)

		//Чтобы подтвердить что наш counter работает просто выполним какие-то действия
		go func() {
			if vv, ok := counter.LoadOrStore(1, "c"); ok {
				fmt.Println(vv)
			}
			if vv, ok := counter.LoadOrStore(2, "c"); !ok {
				fmt.Println(vv)
			}
			counter.Store(1, "a")

			// Загрузить метод, получить значение
			if v, ok := counter.Load(1); ok {
				fmt.Println(v)
			}

			wg.Done()
		}()
	}
	//Ждём группу горутин
	wg.Wait()
}
