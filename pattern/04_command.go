package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

// Применимость: Используется, когда нужно инкапсулировать запрос как объект, позволяя параметризовать клиентов с различными запросами, организовать очередь запросов и поддерживать отмену операций.
// Плюсы: Разделение отправителя и получателя, легкость добавления новых команд.
// Минусы: Может привести к созданию множества классов для каждой команды.
// Пример в Go: Реализация обработчиков HTTP-запросов как команд, позволяя легко добавлять новые обработчики.

// Интерфейс команды
type Command interface {
	Execute()
}

// Конкретные команды
type ConcreteCommandA struct {
	receiver *Receiver
}

func (c *ConcreteCommandA) Execute() {
	c.receiver.ActionA()
}

type ConcreteCommandB struct {
	receiver *Receiver
}

func (c *ConcreteCommandB) Execute() {
	c.receiver.ActionB()
}

// Получатель
type Receiver struct{}

func (r *Receiver) ActionA() {
	fmt.Println("Receiver: Выполнение действия A")
}

func (r *Receiver) ActionB() {
	fmt.Println("Receiver: Выполнение действия B")
}

// Инициатор команды
type Invoker struct {
	command Command
}

func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}

func main() {
	receiver := &Receiver{}

	commandA := &ConcreteCommandA{receiver}
	commandB := &ConcreteCommandB{receiver}

	invoker := &Invoker{}

	invoker.SetCommand(commandA)
	invoker.ExecuteCommand()

	invoker.SetCommand(commandB)
	invoker.ExecuteCommand()
}
