package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

// Применимость: Используется, когда нужно определить семейство алгоритмов, инкапсулировать каждый из них и делать их взаимозаменяемыми.
// Плюсы: Разделение алгоритма от контекста, возможность добавления новых стратегий без изменения контекста.
// Минусы: Увеличение числа классов.
// Пример в Go: Пакет сортировки sort в стандартной библиотеке, где можно указать свою функцию сравнения для различных стратегий сортировки.

// Интерфейс стратегии
type Strategy interface {
	ExecuteStrategy()
}

// Конкретная стратегия A
type ConcreteStrategyA struct{}

func (s *ConcreteStrategyA) ExecuteStrategy() {
	fmt.Println("Выполнение стратегии A")
}

// Конкретная стратегия B
type ConcreteStrategyB struct{}

func (s *ConcreteStrategyB) ExecuteStrategy() {
	fmt.Println("Выполнение стратегии B")
}

// Контекст
type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) Execute() {
	c.strategy.ExecuteStrategy()
}

func main() {
	context := &Context{}

	strategyA := &ConcreteStrategyA{}
	context.SetStrategy(strategyA)
	context.Execute()

	strategyB := &ConcreteStrategyB{}
	context.SetStrategy(strategyB)
	context.Execute()
}
