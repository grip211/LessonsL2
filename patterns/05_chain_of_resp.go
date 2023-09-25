package main

import "fmt"

//Цепочка обязанностей — это поведенческий паттерн,
//позволяющий передавать запрос по цепочке потенциальных обработчиков, пока один из них не обработает запрос.

// Плюсы:
// 1. Уменьшает зависимость между клиентом и обработчиками (т.е. каждый обработчик независимо выполняет свою роль и свою логику внутреннюю).
// 	(изменять эту логику тоже можно будет от того общее у нас обработка или нет)
//	2.	 Реализует принцип единственной обязанности(каждый сервис выполняет свою роль).
// 	3.	Реализует принцип открытости/закрытости.

// Минусы: Запрос может остаться не обаботанным (логика была нарушена и предварительная задача не была обработана).

// Интерфейс сервиса

type Service interface {
	Execute(*Data)
	SetNext(service Service)
}

// Структура данных которая переходит от одного сервиса к другому

type Data struct {
	GetSource    bool // выполнилось ли прием данных
	UpdateSource bool // ставит отметку кто обработал данные
}

// создаем устройство
// сервис получения данных

type Deviice struct {
	Name string
	Next Service
}

func (device *Deviice) Execute(data *Data) {
	if data.GetSource {
		fmt.Printf("Data from device[%s] already get.\n", device.Name)
		device.Next.Execute(data)
		return
	}
	fmt.Printf("Get data from device[%s] already get.\n", device.Name)
	data.GetSource = true
	device.Next.Execute(data)
}

func (device *Deviice) SetNext(svc Service) {
	device.Next = svc
}

// сервис обработки данных

type UpdateDataService struct {
	Name string
	Next Service
}

func (upd *UpdateDataService) Execute(data *Data) {
	if data.UpdateSource {
		fmt.Printf("Data in service [%s] is already update.\n", upd.Name)
		upd.Next.Execute(data)
		return
	}
	fmt.Printf("Get data from device[%s] already get.\n", upd.Name)
	data.UpdateSource = true
	upd.Next.Execute(data)
}

func (upd *UpdateDataService) SetNext(svc Service) {
	upd.Next = svc
}

// Сервис сохранения данных

type DataService struct {
	Next Service
}

func (upd *DataService) Execute(data *Data) {
	if !data.UpdateSource {
		fmt.Printf("Data not update.")
		return
	}
	fmt.Printf("Data save.")
}

func (upd *DataService) SetNext(svc Service) {
	upd.Next = svc
}

func main() {
	device := &Deviice{Name: "Device-1"}
	updateSvc := &UpdateDataService{Name: "Update-1"}
	dataSvc := &DataService{}

	device.SetNext(updateSvc)
	updateSvc.SetNext(dataSvc)

	data := &Data{}
	device.Execute(data)
}
