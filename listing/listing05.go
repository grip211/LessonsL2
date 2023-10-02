package main

// Какая-то ошибка со строковым содержимым
type customError struct {
	msg string
}

// Реализация метода Error() интерфейса error
func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	// test возвращает nil и customError имплементирует интерфейс error, так что err.data = nil, err.itab = error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}

/*
Вывод:
		error

Ответ:
В данном случае возвращаемая ошибка не равна нулю, поскольку в переменной err сохраняется информация о типе.
	Для функций, которые возвращают ошибки, рекомендуется всегда использовать тип error в своей подписи,
		а не конкретный тип, такой как *customError, чтобы обеспечить правильное создание ошибок.
*/
