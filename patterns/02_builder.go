package main

import "fmt"

const (
	AsusCollectorType = "asus"
	HPCollectorType   = "hp"
)

type Collector interface {
	SetCore()
	SetBrand()
	SetMemory()
	SetMonitor()
	SetGraphicCard()
	GetComputer() Computer
}

func GetCollector(collectorType string) Collector {
	switch collectorType {
	default:
		return nil
	case AsusCollectorType:
		return &AsusCollector{}
	case HPCollectorType:
		return &HpCollector{}

	}
}

type AsusCollector struct {
	Core        int
	Brand       string
	Memory      int
	Monitor     int
	GraphicCard int
}

func (collector *AsusCollector) SetCore() {

	collector.Core = 4
}

func (collector *AsusCollector) SetBrand() {
	collector.Brand = "Asus"
}

func (collector *AsusCollector) SetMemory() {
	collector.Memory = 6
}

func (collector *AsusCollector) SetMonitor() {
	collector.Monitor = 1
}

func (collector *AsusCollector) SetGraphicCard() {
	collector.GraphicCard = 1
}

func (collector *AsusCollector) GetComputer() Computer {
	return Computer{
		Core:        collector.Core,
		Brand:       collector.Brand,
		Memory:      collector.Memory,
		Monitor:     collector.Monitor,
		GraphicCard: collector.GraphicCard,
	}
}

type HpCollector struct {
	Core        int
	Brand       string
	Memory      int
	Monitor     int
	GraphicCard int
}

func (collector *HpCollector) SetCore() {

	collector.Core = 4
}

func (collector *HpCollector) SetBrand() {
	collector.Brand = "Hp"
}

func (collector *HpCollector) SetMemory() {
	collector.Memory = 16
}

func (collector *HpCollector) SetMonitor() {
	collector.Monitor = 2
}

func (collector *HpCollector) SetGraphicCard() {
	collector.GraphicCard = 1
}

func (collector *HpCollector) GetComputer() Computer {
	return Computer{
		Core:        collector.Core,
		Brand:       collector.Brand,
		Memory:      collector.Memory,
		Monitor:     collector.Monitor,
		GraphicCard: collector.GraphicCard,
	}
}

type Computer struct {
	Core        int
	Brand       string
	Memory      int
	Monitor     int
	GraphicCard int
}

func (pc *Computer) Print() {
	fmt.Printf("%s Core:[%d] Mem:[%d] GraphicCard:[%d] Monitor:[%d]\n", pc.Brand, pc.Core, pc.Memory, pc.Monitor, pc.GraphicCard)
}

type Factory struct {
	Collector Collector
}

func NewFactory(collector Collector) *Factory {
	return &Factory{Collector: collector}
}

func (factory *Factory) SetCollector(collector Collector) {
	factory.Collector = collector
}

func (factory *Factory) CreateComputer() Computer {
	factory.Collector.SetCore()
	factory.Collector.SetMemory()
	factory.Collector.SetBrand()
	factory.Collector.SetGraphicCard()
	factory.Collector.SetMonitor()

	return factory.Collector.GetComputer()
}

func main() {
	asusCollector := GetCollector("asus")
	hpCollector := GetCollector("hp")

	factory := NewFactory(asusCollector)
	asusComputer := factory.CreateComputer()
	asusComputer.Print()

	factory.SetCollector(hpCollector)
	hpComputer := factory.CreateComputer()
	hpComputer.Print()

	factory.SetCollector(asusCollector)
	pc := factory.CreateComputer()
	pc.Print()
}
