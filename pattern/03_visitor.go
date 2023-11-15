package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

// Применимость: Используется, когда необходимо выполнить операцию над набором объектов, но не нужно добавлять эту операцию в класс каждого объекта.
// Плюсы: Разделение алгоритма от структуры объектов, добавление новых операций без изменения классов объектов.
// Минусы: Усложнение структуры программы.
// Пример в Go: Обход дерева или структуры данных, где каждый узел может быть посещен определенным посетителем.

// Интерфейс элемента
type Element interface {
	Accept(visitor Visitor)
}

// Конкретные элементы
type ConcreteElementA struct{}

func (c *ConcreteElementA) Accept(visitor Visitor) {
	visitor.VisitConcreteElementA(c)
}

type ConcreteElementB struct{}

func (c *ConcreteElementB) Accept(visitor Visitor) {
	visitor.VisitConcreteElementB(c)
}

// Интерфейс посетителя
type Visitor interface {
	VisitConcreteElementA(element *ConcreteElementA)
	VisitConcreteElementB(element *ConcreteElementB)
}

// Конкретные посетители
type ConcreteVisitor1 struct{}

func (c *ConcreteVisitor1) VisitConcreteElementA(element *ConcreteElementA) {
	fmt.Println("Посетитель 1 посетил ConcreteElementA")
}

func (c *ConcreteVisitor1) VisitConcreteElementB(element *ConcreteElementB) {
	fmt.Println("Посетитель 1 посетил ConcreteElementB")
}

type ConcreteVisitor2 struct{}

func (c *ConcreteVisitor2) VisitConcreteElementA(element *ConcreteElementA) {
	fmt.Println("Посетитель 2 посетил ConcreteElementA")
}

func (c *ConcreteVisitor2) VisitConcreteElementB(element *ConcreteElementB) {
	fmt.Println("Посетитель 2 посетил ConcreteElementB")
}

// Структура, содержащая элементы
type ObjectStructure struct {
	elements []Element
}

func (o *ObjectStructure) Attach(element Element) {
	o.elements = append(o.elements, element)
}

func (o *ObjectStructure) Accept(visitor Visitor) {
	for _, e := range o.elements {
		e.Accept(visitor)
	}
}

func main() {
	objectStructure := ObjectStructure{}

	objectStructure.Attach(&ConcreteElementA{})
	objectStructure.Attach(&ConcreteElementB{})

	visitor1 := &ConcreteVisitor1{}
	objectStructure.Accept(visitor1)

	visitor2 := &ConcreteVisitor2{}
	objectStructure.Accept(visitor2)
}
