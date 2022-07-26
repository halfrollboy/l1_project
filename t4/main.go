package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*
	Задание 4

	Реализовать постоянную запись данных в канал (главный поток). Реализовать
	набор из N воркеров, которые читают произвольные данные из канала и
	выводят в stdout. Необходима возможность выбора количества воркеров при
	старте.
	Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
	способ завершения работы всех воркеров.

*/

/*
В своей реализации я операюсь на контекст и выделение канала, под системный вызво закрытия
Считаю что так можно точно отловить вызов закрытия программы и смочь закрыть ресурсы
*/

func main() {
	intCh := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	//Ожидает закрытия программы
	go awaitCancel(cancel)
	//Задаём определённое кол-во горутин
	i := 50
	for j := 0; j < i; j++ {
		wg.Add(1)
		//Сделал анонимную ф-цию чтобы не передавать wg и не нарушать логику
		go func(fact int, ch chan int, contx context.Context) {
			go factorial(fact, ch, contx)
			wg.Done()
		}(7, intCh, ctx)
	}

	for {
		num, opened := <-intCh // получаем данные из потока
		if !opened {
			break // если поток закрыт, выход из цикла
		}
		fmt.Println(num)
	}
	wg.Wait()
}

// Какая-то долгая ф-ция
func factorial(n int, ch chan int, ctx context.Context) {
	defer close(ch)
	result := 1
	for i := 1; i <= n; i++ {
		select {
		case <-ctx.Done():
			log.Println("yep - закрываю канал прибираюсь")
			close(ch)
		default:
			log.Println("обычный день")
		}
		result *= i
		time.Sleep(1 * time.Second)
		ch <- result // посылаем по числу
	}
}

//Ожидание завершения на Ctrl+C
func awaitCancel(cancel context.CancelFunc) {
	grace := make(chan os.Signal, 1)
	signal.Notify(grace, syscall.SIGINT, syscall.SIGTERM)
	<-grace
	fmt.Println("Done!")
	cancel()
}
