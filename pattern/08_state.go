package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

// Применимость: Используется, когда объект должен изменять свое поведение в зависимости от внутреннего состояния.
// Плюсы: Легкость добавления новых состояний, избегание множества условий.
// Минусы: Увеличение числа классов.
// Пример в Go: Реализация конечного сценария, где каждое состояние представляет отдельный класс.

// Интерфейс состояния
type State interface {
	Handle()
}

// Конкретные состояния
type ConcreteStateA struct{}

func (s *ConcreteStateA) Handle() {
	fmt.Println("Обработка в состоянии A")
}

type ConcreteStateB struct{}

func (s *ConcreteStateB) Handle() {
	fmt.Println("Обработка в состоянии B")
}

// Контекст
type Context struct {
	state State
}

func (c *Context) SetState(state State) {
	c.state = state
}

func (c *Context) Request() {
	c.state.Handle()
}

func main() {
	context := &Context{}

	stateA := &ConcreteStateA{}
	context.SetState(stateA)
	context.Request()

	stateB := &ConcreteStateB{}
	context.SetState(stateB)
	context.Request()
}
