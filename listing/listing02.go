package main

import "fmt"

// Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и порядок их вызовов.

func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}

func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}

func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}

/*
В функции test значение переменной x хранится в стеке вызывающей функции,
поэтому при вызове отложенной инкременты возвращаемое значение обновится.
В функции anotherTest значение переменной x хранится в стеке внутренней функции.
Перед возвратом значения отложенная инкремента обновит x, но на результат это уже не повлияет.
*/

/*
	Ключевое слово defer используется для отложенного вызова функции.
	При этом, место объявления одной инструкции defer в коде никак не влияет на то, когда та выполнится.

	Функция с defer всегда выполняется перед выходом из внешней функции, в которой defer объявлялась.

	defer добавляет переданную после него функцию в стэк. При возврате внешней функции, вызываются все, добавленные в стэк вызовы.
	Поскольку стэк работает по принципу LIFO (last in first out), значения стэка возвращаются в порядке от последнего к первому.

	Таким образом функции c defer будут вызываться в обратной последовательности от их объявления во внешней функции.
	Аргументы функций, перед которыми указано ключевое слово defer оцениваются немедленно.
	То есть на тот момент, когда переданы в функцию.
*/
