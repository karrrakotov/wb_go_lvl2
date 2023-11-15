package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

// Применимость: Используется, когда создание объекта необходимо делегировать подклассам.
// Плюсы: Позволяет расширять и изменять создание объектов без изменения кода клиента.
// Минусы: Может привести к созданию множества подклассов для каждого типа продукта.
// Пример в Go: Пакет database/sql с методами Open и Driver, где каждый драйвер предоставляет свой фабричный метод.

// Интерфейс продукта
type Product interface {
	Use() string
}

// Интерфейс фабрики
type Factory interface {
	CreateProduct() Product
}

// Конкретный продукт
type ConcreteProductA struct{}

func (p *ConcreteProductA) Use() string {
	return "Продукт A"
}

// Конкретная фабрика
type ConcreteFactoryA struct{}

func (f *ConcreteFactoryA) CreateProduct() Product {
	return &ConcreteProductA{}
}

// Конкретный продукт B
type ConcreteProductB struct{}

func (p *ConcreteProductB) Use() string {
	return "Продукт B"
}

// Конкретная фабрика B
type ConcreteFactoryB struct{}

func (f *ConcreteFactoryB) CreateProduct() Product {
	return &ConcreteProductB{}
}

func main() {
	factoryA := &ConcreteFactoryA{}
	productA := factoryA.CreateProduct()
	fmt.Println(productA.Use())

	factoryB := &ConcreteFactoryB{}
	productB := factoryB.CreateProduct()
	fmt.Println(productB.Use())
}
