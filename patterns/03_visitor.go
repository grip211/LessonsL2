package main

import "fmt"

//Посетитель — это поведенческий паттерн проектирования,
//который позволяет добавлять в программу новые операции, не изменяя классы объектов,
//над которыми эти операции могут выполняться.

//Плюсы: Упрощает добавление операций, работающих со сложными структурами объектов.
// 		Объединяет родственные операции в одном классе.
// 		Посетитель может накапливать состояние при обходе структуры элементов.

// Минусы:	Паттерн не оправдан, если иерархия элементов часто меняется.
// 			Может привести к нарушению инкапсуляции элементов.

// элемент

type Shape interface {
	getType() string
	accept(Visitor)
}

//Конкретный элемент

type Square struct {
	side int
}

func (s *Square) accept(v Visitor) {
	v.visitForSquare(s)
}

func (s *Square) getType() string {
	return "Square"
}

//Конкретный элемент

type Circle struct {
	radius int
}

func (c *Circle) accept(v Visitor) {
	v.visitForCircle(c)
}

func (c *Circle) getType() string {
	return "Circle"
}

//Конкретный элемент

type Rectangle struct {
	l int
	b int
}

func (t *Rectangle) accept(v Visitor) {
	v.visitForRectangle(t)
}

func (t *Rectangle) getType() string {
	return "rectangle"
}

//Посетитель

type Visitor interface {
	visitForSquare(*Square)
	visitForCircle(*Circle)
	visitForRectangle(*Rectangle)
}

//Конкретный посетитель

type AreaCalculator struct {
	area int
}

func (a *AreaCalculator) visitForSquare(s *Square) {
	// Вычисляем площадь квадрата.
	// Затем присвойте значение переменной экземпляра области.
	fmt.Println("Calculating area for square")
}

func (a *AreaCalculator) visitForCircle(s *Circle) {
	fmt.Println("Calculating area for circle")
}

func (a *AreaCalculator) visitForRectangle(s *Rectangle) {
	fmt.Println("Calculating area for rectangle")
}

//Конкретный посетитель

type MiddleCoordinates struct {
	x int
	y int
}

func (a *MiddleCoordinates) visitForSquare(s *Square) {
	// Вычисляем координаты средней точки квадрата.
	// Затем присвойте переменным экземпляра x и y.

	fmt.Println("Calculating middle point coordinates for square")
}

func (a *MiddleCoordinates) visitForCircle(c *Circle) {
	fmt.Println("Calculating middle point coordinates for circle")
}

func (a *MiddleCoordinates) visitForRectangle(t *Rectangle) {
	fmt.Println("Calculating middle point coordinates for rectangle")
}

// Клиентский код.

func main() {
	square := &Square{side: 2}
	circle := &Circle{radius: 3}
	rectangle := &Rectangle{l: 2, b: 3}

	areaCalculator := &AreaCalculator{}

	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	rectangle.accept(areaCalculator)

	fmt.Println()
	middleCoordinates := &MiddleCoordinates{}
	square.accept(middleCoordinates)
	circle.accept(middleCoordinates)
	rectangle.accept(middleCoordinates)
}
