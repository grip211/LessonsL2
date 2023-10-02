package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Функция принимает переменное кол-во интов, кидает их в канал, закрывает его и возвращает канал
func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

// Функция слияния двух каналов
func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			// Нужна проверка на закрытие канала: v, ok := <- a;
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
		// Тут нужно закрыть канал
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
}

/*
Ответ:
		1
		2
		3
		5
		4
		6
		7
		0
		0
		0
		0
		...

rand.Intn(1000)) * time.Millisecond  течение пока этот отчет идет функция сколько успевает выложить столько и идет,
	рандомным образом. И результат с каждым разом будет у нас разным.

После получения чисел, переданных в asChan,
	в канал c в main начинают бесконечно сыпаться нули из закрытых каналов a и b в merge,
		так как их состояние (закрыт или открыт) никак не проверяется в коде.
*/
