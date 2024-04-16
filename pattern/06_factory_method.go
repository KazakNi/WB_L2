package pattern

import (
	"errors"
	"fmt"
)

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern

	Применяется:

	1.  Когда заранее неизвестны типы и зависимости объектов, с которыми должен работать ваш код.
	2. Когда вы хотите дать возможность пользователям расширять части вашего фреймворка или библиотеки.
	3.  Когда вы хотите экономить системные ресурсы, повторно используя уже созданные объекты, вместо порождения новых.

	Плюсы:
	- Избавляет класс от привязки к конкретным классам продуктов.
 	- Выделяет код производства продуктов в одно место, упрощая поддержку кода.
 	- Упрощает добавление новых продуктов в программу.
 	- Реализует принцип открытости/закрытости.

	Минус:
	- Может привести к созданию больших параллельных иерархий классов, так как для каждого класса продукта надо создать свой подкласс создателя.
*/

type IBakedFood interface {
	setName(name string)
	getType() string
	getName() string
	setType(kind string)
}

type BakedGood struct {
	name string
	kind string
}

func (b *BakedGood) setName(name string) {
	b.name = name
}

func (b *BakedGood) getName() string {
	return b.name
}

func (b *BakedGood) setType(kind string) {
	b.kind = kind
}

func (b *BakedGood) getType() string {
	return b.kind
}

type Pizza struct {
	BakedGood
}

func newPizza() IBakedFood {
	return &Pizza{BakedGood: BakedGood{
		name: "Pizza",
		kind: "Round baked dough with sasuage, cheese and ham",
	}}
}

type Cake struct {
	BakedGood
}

func newCake() IBakedFood {
	return &Cake{
		BakedGood: BakedGood{
			name: "Cake",
			kind: "sweet baked food",
		},
	}
}

func getBakery(typeOfBakery string) (IBakedFood, error) {
	if typeOfBakery == "Cake" {
		return newCake(), nil
	}
	if typeOfBakery == "Pizza" {
		return newPizza(), nil
	}

	return nil, errors.New("Ooops, no producer for such type")
}

func main() {
	pizza, _ := getBakery("Pizza")
	cake, _ := getBakery("Cake")

	printDetails(pizza)
	printDetails(cake)

}

func printDetails(b IBakedFood) {
	fmt.Printf("Gun: %s", b.getName())
	fmt.Println()
	fmt.Printf("Power: %d", b.getType())
	fmt.Println()
}
