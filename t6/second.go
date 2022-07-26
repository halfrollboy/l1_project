package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func timeout(w *sync.WaitGroup, t time.Duration) bool {
	temp := make(chan int)
	go func() {
		defer close(temp)
		time.Sleep(5 * time.Second)
		w.Wait()
	}()
	select {
	case <-temp:
		return false
	case <-time.After(t):
		return true
	}
}

//В этом коде временной интервал, который будет использоваться при вызове
// time.After(), является аргументом функции timeout(), следовательно, является
// переменной величиной. Как и в прошлый раз, логика прерывания в случае превышения временного лимита реализована в блоке select. Кроме того, вызов w.Wait()
// заставит функцию timeout() бесконечно ожидать соответствующую функцию
// sync.Done(), чтобы завершить работу. Когда w.Wait() закончит работу, будет выполнена первая ветвь оператора select
func main2() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Need a time duration!")
		return
	}
	var w sync.WaitGroup
	w.Add(1)
	t, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	//Функция time.Duration() преобразует целочисленное значение в переменную
	//time.Duration, которую мы будем использовать позже.

	duration := time.Duration(int32(t)) * time.Millisecond
	fmt.Printf("Timeout period is %s\n", duration)
	if timeout(&w, duration) {
		fmt.Println("Timed out!")
	} else {
		fmt.Println("OK!")
	}
	//После завершения вызова w.Done() предшествующая ему функция timeout()
	//также завершится. Однако у второго вызова timeout() нет оператора sync.Done(),
	//которого она могла бы ожидать.
	w.Done()
	if timeout(&w, duration) {
		fmt.Println("Timed out!")
	} else {
		fmt.Println("OK!")
	}
}
