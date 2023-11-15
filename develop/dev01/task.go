package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

const (
	onlyTimeStampNano = "15:04:05.000000000"
)

// timer - структура, с помощью которой будем получать время
type timer struct{}

// DoTimer - интерфейс, с помощью которого будем получать время разными способами
type DoTimer interface {
	getLocalTime() time.Time
	getCurrentTime() time.Time
	getQueryCurrentTime() time.Time
}

// NewTimer - конструктор для структуры timer
func NewTimer() DoTimer {
	return &timer{}
}

// Функция для получения локального времени
func (m *timer) getLocalTime() time.Time {
	return time.Now()
}

// Функция для получения текущего/точного времени
func (m *timer) getCurrentTime() time.Time {
	// Получение времени с NTP-сервера
	ntpTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при получении времени: %v\n", err)
		os.Exit(1)
	}

	return ntpTime
}

// Функция для получения точного времени, а также некоторые дополнительные данные синхронизации
func (m *timer) getQueryCurrentTime() time.Time {
	// Получение времени с NTP-сервера с помощью запроса
	response, err := ntp.Query("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при получении времени: %v\n", err)
		os.Exit(1)
	}

	// Получаем более точное время, т.к. плюсуем к текущему времени - ClockOffset (предполагаемое смещение локальных системных часов относительно часов сервера)
	// Для более точной настройки, можно применять и другие параметры запроса
	time := time.Now().Add(response.ClockOffset)

	return time
}

func main() {
	// Создаем объект структуры timer
	timer := NewTimer()

	// Вызываем готовые методы
	localTime := timer.getLocalTime()
	currentTime := timer.getCurrentTime()
	exactTime := timer.getQueryCurrentTime()

	// Вывод локального, текущего и точного времени
	fmt.Printf("Локальное время: %s\n", localTime.Format(onlyTimeStampNano))
	fmt.Printf("Текущее время с NTP: %s\n", currentTime.Format(onlyTimeStampNano))
	fmt.Printf("Точное время с NTP Query: %v\n", exactTime.Format(onlyTimeStampNano))
}
