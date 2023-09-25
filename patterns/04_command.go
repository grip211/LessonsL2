package main

import "fmt"

//Команда — это поведенческий паттерн,
//позволяющий заворачивать запросы или простые операции в отдельные объекты.

//Плюсы: Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют.
//		Позволяет реализовать простую отмену и повтор операций.
// 		Позволяет реализовать отложенный запуск операций.
// 		Позволяет собирать сложные команды из простых.
//		Реализует принцип открытости/закрытости.

// Минусы: Усложняет код программы из-за введения множества дополнительных классов.

//Отправитель

type Button struct {
	command Command
}

func (b *Button) press() {
	b.command.execute()
}

//Интерфейс команды

type Command interface {
	execute()
}

//Конкретная команда

type OnCommand struct {
	device Device
}

func (c *OnCommand) execute() {
	c.device.on()
}

//Конкретная команда

type OffCommand struct {
	device Device
}

func (c *OffCommand) execute() {
	c.device.off()
}

//Интерфейс получателя

type Device interface {
	on()
	off()
}

//Конкретный получатель

type Tv struct {
	isRunning bool
}

func (t *Tv) on() {
	t.isRunning = true
	fmt.Println("Включение телевизора")
}

func (t *Tv) off() {
	t.isRunning = false
	fmt.Println("Выключение телевизора")
}

//Клиентский код

func main() {
	tv := &Tv{}

	onCommand := &OnCommand{
		device: tv,
	}

	offCommand := &OffCommand{
		device: tv,
	}

	onButton := &Button{
		command: onCommand,
	}
	onButton.press()

	offButton := &Button{
		command: offCommand,
	}
	offButton.press()
}
