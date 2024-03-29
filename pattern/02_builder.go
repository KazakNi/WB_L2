package pattern

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/* Паттерн используется в случаях, когда необходимо создавать класс с нетривиальной структурой,
   "строитель" изолирует процесс создания экземпляра класса от его представления

Pros: Отдельный класс для последовательного создания сложного класса, отвязывает реализацию конструктора от представления класса
	  Дает более тонкий контроль над процессом конструирования
Cons:  Усложнение кодовой базы
*/

type Vehicle struct {
	engineType string
	modelType  string
	bodyType   string
}

type VehicleBuilder interface {
	setEngineType()
	setModelType()
	setBodyType()
	getVehicle() Vehicle
}

func getBuilder(builderType string) VehicleBuilder {
	if builderType == "car" {
		return newCarBuilder()
	}

	if builderType == "truck" {
		return newTruckBuilder()
	}
	return nil
}

type CarBuilder struct {
	engineType string
	modelType  string
	bodyType   string
}

func newCarBuilder() *CarBuilder {
	return &CarBuilder{}
}

func (b *CarBuilder) setBodyType() {
	b.bodyType = "saloon"
}

func (b *CarBuilder) setEngineType() {
	b.engineType = "petroleum"
}

func (b *CarBuilder) setModelType() {
	b.modelType = "family"
}

func (b *CarBuilder) getVehicle() Vehicle {
	return Vehicle{bodyType: b.bodyType,
		modelType:  b.modelType,
		engineType: b.engineType,
	}
}

type TruckBuilder struct {
	engineType string
	modelType  string
	bodyType   string
}

func newTruckBuilder() *TruckBuilder {
	return &TruckBuilder{}
}

func (b *TruckBuilder) setBodyType() {
	b.bodyType = "trailer"
}

func (b *TruckBuilder) setEngineType() {
	b.engineType = "disel"
}

func (b *TruckBuilder) setModelType() {
	b.modelType = "van"
}

func (b *TruckBuilder) getVehicle() Vehicle {
	return Vehicle{bodyType: b.bodyType,
		modelType:  b.modelType,
		engineType: b.engineType,
	}
}

type Director struct {
	builder VehicleBuilder
}

func newDirector(b VehicleBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b VehicleBuilder) {
	d.builder = b
}

func (d *Director) buildVehicle() Vehicle {
	d.builder.setBodyType()
	d.builder.setEngineType()
	d.builder.setModelType()
	return d.builder.getVehicle()
}

func main() {
	carBuilder := getBuilder("saloom")
	truckBuilder := getBuilder("truck")

	director := newDirector(carBuilder)
	carVehicle := director.buildVehicle()

	fmt.Printf("Normal House Door Type: %s\n", carVehicle.modelType)
	fmt.Printf("Normal House Window Type: %s\n", carVehicle.engineType)
	fmt.Printf("Normal House Num Floor: %d\n", carVehicle.bodyType)

	director.setBuilder(truckBuilder)
	truckVehicle := director.buildVehicle()
}
