package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

// Применимость: Используется, когда нужно предоставить простой интерфейс к сложной подсистеме. Фасад позволяет скрыть сложность и уменьшить зависимости клиента от подсистемы.
// Плюсы: Упрощает интерфейс для клиента, уменьшает связанность кода.
// Минусы: Может не предоставлять полного контроля над подсистемой.
// Пример в Go: Если есть сложная система, например, библиотека для работы с базой данных, можно создать фасад для предоставления простого API для основных операций.

// Сложная подсистема
type SubsystemA struct {
}

func (s *SubsystemA) OperationA() {
	fmt.Println("ПодсистемаA: Операция A")
}

type SubsystemB struct {
}

func (s *SubsystemB) OperationB() {
	fmt.Println("ПодсистемаB: Операция B")
}

// Фасад, скрывающий сложность подсистемы
type Facade struct {
	subsystemA *SubsystemA
	subsystemB *SubsystemB
}

func NewFacade() *Facade {
	return &Facade{
		subsystemA: &SubsystemA{},
		subsystemB: &SubsystemB{},
	}
}

func (f *Facade) Operation() {
	fmt.Println("Фасад: Операция")
	f.subsystemA.OperationA()
	f.subsystemB.OperationB()
}

func main() {
	facade := NewFacade()
	facade.Operation()
}
