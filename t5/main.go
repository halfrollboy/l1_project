package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

/*
	Задание 5

	Разработать программу, которая будет последовательно отправлять значения в
	канал, а с другой стороны канала — читать. По истечению N секунд программа
	должна завершаться.

*/

func main() {
	intCh := make(chan int)
	//Контекст с заданным временем
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	go factorial(7, intCh, ctx)

	for {
		num, opened := <-intCh // получаем данные из потока
		if !opened {
			break // если поток закрыт, выход из цикла
		}
		fmt.Println(num)
	}
}

func factorial(n int, ch chan int, ctx context.Context) {
	defer close(ch)
	result := 1
	for i := 1; i <= n; i++ {
		select {
		case <-ctx.Done(): //Закрываем канал по истечению времени
			log.Println("yep - закрываю канал прибираюсь")
			close(ch)
		default:
			log.Println("обычный день")
			result *= i
			ch <- result
		}
		time.Sleep(2 * time.Second)
		// посылаем по числу
	}
}
