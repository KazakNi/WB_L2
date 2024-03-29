package pattern

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

// Ответ

/*
Паттерн создаёт единый интерфейс для отвязывания бизнес-логики от сторонних библиотек и реализации сторонних классов.
Фасад особенно полезен, если  нужна только часть возможностей внешней сложной системы, он скрывает от клиента ненужный ему функционал.

Pros: ослабляет связь между сложной системой сервиса и клиентом
Cons: усложняет имплементацию приложения, риск для связывания фасада со сложной логикой подсистем
*/

func start() {
	order := Order{
		Products: []Product{
			{"iPhone"},
			{"iWatch"},
		},
		BuyerID: "customer_account_id",
		Amount:  200,
	}

	manager := PackageManager{}
	packageId, err := manager.PayAndPreparePackage(order)

	if err != nil {
		panic("Unable to create package: " + packageId)
	}

	ots := NewOrderTrackingService(os.Getenv("ORDER_TRACKING_SERVICE_TOKEN"))
	err = ots.TrackPackage(packageId)
	if err != nil {
		panic("Unable to track package: " + packageId)
	}

}

type PaymentsService struct {
	AuthToken string
}

func (ps PaymentsService) Execute(amount int, sender string) error {
	fmt.Printf("Transact $%d from `%s` to company account ID\n", amount, sender)
	return nil
}

func NewPaymentsService(token string) *PaymentsService {
	return &PaymentsService{token}
}

type FulfilmentService struct {
	AuthToken string
}

func NewFulfilmentService(token string) *FulfilmentService {
	return &FulfilmentService{token}
}

func (fs FulfilmentService) CreatePackage(products []Product) string {
	packageId := strconv.Itoa(rand.Intn(3500-2000) + 2000)
	fmt.Println("Creating new package with ID:", packageId)

	return packageId
}

type OrderTrackingService struct {
	AuthToken string
}

func NewOrderTrackingService(token string) *OrderTrackingService {
	return &OrderTrackingService{token}
}

func (ots OrderTrackingService) TrackPackage(id string) error {
	// Subscribing via webhooks, for instance
	fmt.Println("Subscribed to track package with ID:", id)
	return nil
}

type Order struct {
	Products []Product
	BuyerID  string
	Amount   int
}

type Product struct {
	Name string
}

type PackageManager struct {
}

func (sm PackageManager) PayAndPreparePackage(order Order) (string, error) {
	var err error

	paymentService := NewPaymentsService(os.Getenv("PAYMENT_SERVICE_TOKEN"))
	err = paymentService.Execute(order.Amount, order.BuyerID)
	if err != nil {
		return "", err
	}

	fulfilmentService := NewFulfilmentService(os.Getenv("FULFILMENT_SERVICE_TOKEN"))
	packageId := fulfilmentService.CreatePackage(order.Products)

	return packageId, err
}
