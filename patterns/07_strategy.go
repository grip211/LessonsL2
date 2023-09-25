package main

import "fmt"

//Стратегия — это поведенческий паттерн,(который определяет схожие алгоритмы и помещает каждую из них в свою отдельную структуру,
//после чего алгоритмы могут взаимодействовать в исполняемой программе)
//выносит набор алгоритмов в собственные классы и делает их взаимозаменимыми.

//Вам нужно добраться до аэропорта. Можно доехать на автобусе, такси или велосипеде.
//Здесь вид транспорта является стратегией.
//Вы выбираете конкретную стратегию в зависимости от контекста — наличия денег или времени до отлёта.

// Плюсы:
// 1. Горячая замена алгоритмов на лету (т.е. переопределить алгоритм и у нас будет работать тот алгоритм который мы определили на лету).
// 2. Изолирует код и данные алгоритмов от остальных классов.
// 3. Уход от наследования к делегированию непосредственному алгоритму.
// 4. Реализует принцип открытости/закрытости.

// Минусы:
// 1. Усложняет программу за счёт дополнительных классов.
// 2. Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.

// Интерфейс стратегии

type Strategy interface {
	Route(startPoint int, endPoint int)
}

type Navigator struct {
	Strategy
}

func (nav *Navigator) SetStrategy(str Strategy) {
	nav.Strategy = str
}

// построение маршрута на машине

type RoadStrategy struct {
}

func (r *RoadStrategy) Route(startPoint int, endPoint int) {
	avgSpeed := 30
	trafficJam := 2
	total := endPoint - startPoint
	totalTime := total * 40 * trafficJam
	fmt.Printf("Road A:[%d] to B:[%d] Avg speed:[%d] Traffic jam:[%d] Total: [%d] Total time:[%d] min \n", startPoint,
		endPoint, avgSpeed, trafficJam, total, totalTime)
}

// построение маршрута на общественном транспорте

type PublicTransportStrategy struct {
}

func (r *PublicTransportStrategy) Route(startPoint int, endPoint int) {
	avgSpeed := 40
	total := endPoint - startPoint
	totalTime := total * 40
	fmt.Printf("Public Transport A:[%d] to B:[%d] Avg speed:[%d]  Total: [%d] Total time:[%d] min \n", startPoint,
		endPoint, avgSpeed, total, totalTime)
}

// построение маршрута пешком

type WalkStrategy struct {
}

func (r *WalkStrategy) Route(startPoint int, endPoint int) {
	avgSpeed := 4
	total := endPoint - startPoint
	totalTime := total * 60
	fmt.Printf("Walk  A:[%d] to B:[%d] Avg speed:[%d] Total: [%d] Total time:[%d] min \n", startPoint,
		endPoint, avgSpeed, total, totalTime)
}

var (
	start      = 10
	end        = 100
	strategies = []Strategy{
		&PublicTransportStrategy{},
		&RoadStrategy{},
		&WalkStrategy{},
	}
)

func main() {
	nav := Navigator{}
	for _, strategy := range strategies {
		nav.SetStrategy(strategy)
		nav.Route(start, end)
	}

}
