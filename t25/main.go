package main

import (
	"fmt"
	"log"
	"time"
)

/*
	Задание 25

	Реализовать собственную функцию sleep.
*/

func sleep(d time.Duration) {
	<-time.After(d)
}

//Считаем чтобы показать что основной поток уснул
func geter() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
	fmt.Println("Засыпаю")
	time.Sleep(2 * time.Second)
	fmt.Println("Проснулся")
}

//Засыпаем на 5 секунд времени
func CallSleep() {
	log.Println("Do something")
	sleep(5 * time.Second)
	log.Println("Something else")
}

func main() {
	go geter()
	CallSleep()
}
