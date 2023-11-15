package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

// Применимость: Используется, когда необходимо передавать запросы последовательно по цепочке обработчиков.
// Плюсы: Уменьшение зависимости между отправителем и получателем, динамическое добавление и изменение обработчиков.
// Минусы: Нет гарантии обработки запроса.
// Пример в Go: Middleware в веб-фреймворках, где каждый обработчик представляет собой отдельное звено в цепочке обработки HTTP-запроса.

// Интерфейс обработчика
type Handler interface {
	SetNext(handler Handler)
	HandleRequest(request int)
}

// Конкретные обработчики
type ConcreteHandlerA struct {
	next Handler
}

func (h *ConcreteHandlerA) SetNext(handler Handler) {
	h.next = handler
}

func (h *ConcreteHandlerA) HandleRequest(request int) {
	if request < 10 {
		fmt.Println("ConcreteHandlerA обрабатывает запрос")
	} else if h.next != nil {
		h.next.HandleRequest(request)
	}
}

type ConcreteHandlerB struct {
	next Handler
}

func (h *ConcreteHandlerB) SetNext(handler Handler) {
	h.next = handler
}

func (h *ConcreteHandlerB) HandleRequest(request int) {
	if request >= 10 && request < 20 {
		fmt.Println("ConcreteHandlerB обрабатывает запрос")
	} else if h.next != nil {
		h.next.HandleRequest(request)
	}
}

func main() {
	handlerA := &ConcreteHandlerA{}
	handlerB := &ConcreteHandlerB{}

	handlerA.SetNext(handlerB)

	requests := []int{5, 12, 15, 25}

	for _, req := range requests {
		handlerA.HandleRequest(req)
	}
}
