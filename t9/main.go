package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

/*
	Задание 9

	Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
	массива, во второй — результат операции x*2, после чего данные из второго
	канала должны выводиться в stdout.


*/

func main() {
	intCh := make(chan int)
	secondCh := make(chan int)
	//Контекст с заданным временем
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	go kvadrat(7, intCh, ctx)
	go printer(secondCh, ctx)
	for {
		num, opened := <-intCh // получаем данные из потока
		if !opened {
			break // если поток закрыт, выход из цикла
		}
		secondCh <- num * num
	}
}

func printer(ch chan int, ctx context.Context) {

	for {
		select {
		case <-ctx.Done(): //Закрываем канал по истечению времени
			log.Println("yep - закрываю канал прибираюсь")
			close(ch)
		case a := <-ch:
			fmt.Println("result", a)
		}
	}
}

func kvadrat(n int, ch chan int, ctx context.Context) {
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
		time.Sleep(1 * time.Second)
		// посылаем по числу
	}
}
