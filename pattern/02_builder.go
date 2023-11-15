package pattern

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

// Применимость: Используется, когда создание сложного объекта должно быть разделено на несколько шагов, и клиент должен иметь возможность создавать различные варианты объекта.
// Плюсы: Упрощает процесс создания объекта, позволяет создавать различные конфигурации.
// Минусы: Может привести к созданию большого числа классов, если есть много вариантов конфигураций.
// Пример в Go: Построение HTTP-запроса с использованием net/http пакета, где можно добавлять заголовки, параметры и другие аспекты запроса пошагово.

// Продукт
type Product struct {
	Part1 string
	Part2 string
}

// Строитель
type Builder interface {
	BuildPart1()
	BuildPart2()
	GetProduct() *Product
}

// Конкретный строитель
type ConcreteBuilder struct {
	product *Product
}

func NewConcreteBuilder() *ConcreteBuilder {
	return &ConcreteBuilder{
		product: &Product{},
	}
}

func (b *ConcreteBuilder) BuildPart1() {
	b.product.Part1 = "Часть 1 построена"
}

func (b *ConcreteBuilder) BuildPart2() {
	b.product.Part2 = "Часть 2 построена"
}

func (b *ConcreteBuilder) GetProduct() *Product {
	return b.product
}

// Директор
type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

func (d *Director) Construct() *Product {
	d.builder.BuildPart1()
	d.builder.BuildPart2()
	return d.builder.GetProduct()
}

func main() {
	concreteBuilder := NewConcreteBuilder()
	director := NewDirector(concreteBuilder)
	product := director.Construct()

	fmt.Printf("Часть 1: %s\n", product.Part1)
	fmt.Printf("Часть 2: %s\n", product.Part2)
}
