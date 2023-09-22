package main

import (
	"errors"
	"fmt"
	"time"
)

//Фасад — это структурный паттерн, который предоставляет простой (но урезанный) интерфейс
//к сложной системе объектов, библиотеке или фреймворку.
//(упростить интерфейс пользователь разработчик который работает на нашем сервисе, реализовал интерфейс как можно проще.
//скрываем от пользователя и даем ему мимнум функционала. от поведения сложной подсистемы.)

// Применимость: Когда вам нужно представить простой или урезанный интерфейс к сложной подсистеме.
//				Когда вы хотите разложить подсистему на отдельные слои.

// Плюсы Фасада: Изолирует клиентов от компонентов сложной подсистемы.
// Минусы: Фасад рискует стать божественным объектом.
//(Будем привязаны к этому объекту и все последующие функции будут проходить через эту функцию)

var (
	bank = Bank{
		Name:  "Банк",
		Cards: []Card{},
	}
	card1 = Card{
		Name:    "CRD-1",
		Balance: 200,
		Bank:    &bank,
	}
	card2 = Card{
		Name:    "CRD-1",
		Balance: 10,
		Bank:    &bank,
	}
	user = User{
		Name: "Покупатель-1",
		Card: &card1,
	}
	user2 = User{
		Name: "Покупатель-2",
		Card: &card2,
	}
	prod = Product{
		Name:  "Сыр",
		Price: 180,
	}

	shop = Shop{
		Name: "SHOP",
		Products: []Product{
			prod,
		},
	}
)

var Cards []Card

func main() {
	Cards = append(Cards, card1, card2)
	println("[БАНК] Выпуск карт")

	fmt.Printf("[%s]", user.Name)
	err := shop.Sell(user, prod.Name)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("[%s]\n", user2.Name)
	err = shop.Sell(user2, prod.Name)
	if err != nil {
		println(err.Error())
		return
	}
}

type Bank struct {
	Name  string
	Cards []Card
}

func (bank Bank) CheckBalance(cardNumber string) error {
	println(fmt.Sprintf("[Карта] Запрос в банк для проверки остатка"), cardNumber)
	time.Sleep(time.Millisecond * 300)

	for _, card := range bank.Cards {
		if card.Name != cardNumber {
			continue
		}
		if card.Balance <= 0 {
			return errors.New("[Банк] Недостаточно средств!")
		}
	}
	println("[Банк] Остаток положительный!")
	return nil

}

type Card struct {
	Name    string
	Balance float64
	Bank    *Bank
}

func (card Card) CheckBalance() error {
	println("[Карта] Запрос в банк для проверки остатка")
	time.Sleep(time.Millisecond * 800)
	return card.Bank.CheckBalance(card.Name)

}

type Product struct {
	Name  string
	Price float64
}

type Shop struct {
	Name     string
	Products []Product
}

// здесь реализован патерн фассад, который взаимодействует со всеми другими сервисами
//(банк, карты, ассортиментыб магазин, клиент)

func (shop Shop) Sell(user User, product string) error {
	println("[Магазин] Запрос пользователю, для получению остатка по карте")
	time.Sleep(time.Millisecond * 500)
	err := user.Card.CheckBalance()
	if err != nil {
		return err
	}
	fmt.Printf("[Магазин] Проверка - может ли [%s] купить товар! \n", user.Name)
	time.Sleep(time.Millisecond * 500)

	// Здесь реализовано сам фасад
	for _, prod := range shop.Products {
		if prod.Name != product {
			continue
		}
		if prod.Price > user.GetBalance() {
			return errors.New("[Магазин] Недостаточно средств для покупки товара!")

		}
		fmt.Printf("[Магазин] Товар [%s] - куплен!\n", prod.Name)
	}
	return nil
}

type User struct {
	Name string
	Card *Card
}

func (user User) GetBalance() float64 {
	return user.Card.Balance
}
