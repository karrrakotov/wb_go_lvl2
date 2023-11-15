package main

import (
	"fmt"
	"time"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

func or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		// Если нет входных каналов, создаем и возвращаем закрытый канал.
		c := make(chan interface{})
		close(c)
		return c
	case 1:
		// Если есть только один входной канал, возвращаем его.
		return channels[0]
	}

	// Используем горутину для объединения каналов.
	orChannel := make(chan interface{})
	go func() {
		defer close(orChannel)

		select {
		case <-channels[0]:
			// Если первый канал закрыт, закрываем orChannel.
			return
		case <-channels[1]:
			// Если второй канал закрыт, закрываем orChannel.
			return
		case <-or(channels[2:]...):
			// Рекурсивно вызываем or для оставшихся каналов.
			return
		}
	}()

	return orChannel
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(10*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("Закрылся через: %v\n", time.Since(start))
}
