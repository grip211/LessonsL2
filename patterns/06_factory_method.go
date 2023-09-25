package main

import "fmt"

//Фабричный метод (виртуальный конструктор)— это порождающий паттерн проектирования,
//который определяет общий интерфейс для создания объектов в суперклассе,
//позволяя подклассам изменять тип создаваемых объектов.

//плюсы:
//1. Избавляет класс от привязки к конкретным классам(типу) продуктов (объектов), ( у нас есть интрефейс поведения).
//	2. Выделяет код производства продуктов в одно место, упрощая поддержку кода.
// (избаляет нас от привязки к конкретному типу объекта,(т.к. есть интерфейс объекта),
//создаем общий конструктор создания наший объектов, используем одну реалиализацию конструктора  на все подклассы.
// 3.Упрощает добавление новых продуктов в программу.
// 4. Реализует принцип открытости/закрытости.

// Существенный минус:
//Может привести к созданию больших параллельных иерархий классов,
//так как для каждого класса продукта надо создать свой подкласс создателя.
//появляется божественный конструктор(всегда будем инициализировать объекты с помощью него).

// сздаем основной интрфейс поведения

const (
	ServerType           = "server"
	PersonalComputerType = "personal"
	NotebookType         = "notebook"
)

type Computeer interface {
	GetType() string
	PrintDetails()
}

// Создаем фабричный метод
// который инициализирует структуры и возвращает интерфейс компьютеру

func New(typeName string) Computeer {
	switch typeName {
	default:
		fmt.Printf("%s Несуществующий тип объекта!\n", typeName)
		return nil
	case ServerType:
		return NewServer()
	case PersonalComputerType:
		return NewPersonalComputer()
	case NotebookType:
		return NewNoteBook()
	}
}

type Server struct {
	Type   string
	Core   int
	Memory int
}

func NewServer() Computeer {
	return Server{
		Type:   ServerType,
		Core:   16,
		Memory: 512,
	}
}

func (pc Server) GetType() string {
	return pc.Type
}

func (pc Server) PrintDetails() {
	fmt.Printf("%s Core:[%d] Mem:[%d]\n", pc.Type, pc.Core, pc.Memory)
}

type PersonalComputer struct {
	Type    string
	Core    int
	Memory  int
	Monitor bool
}

func NewPersonalComputer() Computeer {
	return PersonalComputer{
		Type:    PersonalComputerType,
		Core:    4,
		Memory:  64,
		Monitor: true,
	}
}

func (pc PersonalComputer) GetType() string {
	return pc.Type
}

func (pc PersonalComputer) PrintDetails() {
	fmt.Printf("%s Core:[%d] Mem:[%d] Monitor:[%v]\n", pc.Type, pc.Core, pc.Memory, pc.Monitor)
}

type Notebook struct {
	Type    string
	Core    int
	Memory  int
	Monitor bool
}

func NewNoteBook() Computeer {
	return Notebook{
		Type:    NotebookType,
		Core:    8,
		Memory:  16,
		Monitor: true,
	}
}

func (pc Notebook) GetType() string {
	return pc.Type
}

func (pc Notebook) PrintDetails() {
	fmt.Printf("%s Core:[%d] Mem:[%d] Monitor:[%v]\n", pc.Type, pc.Core, pc.Memory, pc.Monitor)
}

var types = []string{PersonalComputerType, NotebookType, ServerType, "monoblock"}

func main() {
	for _, typeName := range types {
		computeer := New(typeName)
		if computeer == nil {
			continue
		}
		computeer.PrintDetails()
	}
}
