package service

import (
	"fmt"
	"time"

	"wb_go_lvl2/dev11/internal/model"
	"wb_go_lvl2/dev11/internal/service"
)

type calendar struct{}

// Эмулированная БД
var emulationDB = make(map[time.Time][]model.Event)

func (s *calendar) Create(event model.Event) error {
	// Тут должна быть реализована бизнес-логика сервиса (какие-либо манипуляции с данными, запись в бд, и т.д)
	// Так как в задании нет указаний на использование бд, то данные будут храниться в обычной мапе

	// Перед тем, как сохранить данные в БД, проверим не существует ли уже запись с пользователем на этот день
	records, ok := emulationDB[event.Date]

	// Сначала проверим, существует ли вообще такая дата
	// Если существует, то проверяем есть ли тут пользователь
	// Если ничего найдено не было, тогда сохраняем запись
	if ok {
		for i := 0; i < len(records); i++ {
			if records[i].UserID == event.UserID {
				return fmt.Errorf("запись уже существует в базе данных")
			}
		}
	}

	// Сохраним входные данные в БД
	emulationDB[event.Date] = append(emulationDB[event.Date], event)

	return nil
}

func (s *calendar) Update(event model.Event) error {
	// Ищем в условной БД событие для обновления
	// Если не находим, вернем ошибку

	// Перед тем, как обновить данные в БД, проверим существует ли уже запись с пользователем на этот день
	records, ok := emulationDB[event.Date]

	// Сначала проверим, существует ли вообще такая дата
	// Если существует, то проверяем есть ли тут пользователь
	// Если ничего найдено не было, тогда возвращаем ошибку
	if ok {
		for i := 0; i < len(records); i++ {
			// Если находим запись, то обновляем событие
			if records[i].UserID == event.UserID {
				records[i].Title = event.Title
			}
		}
	}

	return fmt.Errorf("событие не было найдено")
}

func (s *calendar) Delete(event model.Event) error {
	// Ищем в условной БД событие для удаления
	// Если не находим, вернем ошибку

	// Перед тем, как удалить данные в БД, проверим существует ли уже запись с пользователем на этот день
	records, ok := emulationDB[event.Date]

	// Сначала проверим, существует ли вообще такая дата
	// Если существует, то проверяем есть ли тут пользователь
	// Если ничего найдено не было, тогда возвращаем ошибку
	if ok {
		for i := 0; i < len(records); i++ {
			// Если находим запись, то удаляем событие
			if records[i].UserID == event.UserID {
				records[i] = records[len(records)-1]
				records = records[:len(records)-1]
			}
		}
	}

	return fmt.Errorf("событие не было найдено")
}

func (s *calendar) EventsForDay(date time.Time) (events []model.Event, err error) {
	// Ищем в условной БД события для определенной даты
	// Если не находим, вернем ошибку

	events, ok := emulationDB[date]

	// Сначала проверим, существует ли вообще такая дата
	// Если существует, то возвращаем
	if ok {
		return events, nil
	}

	return events, fmt.Errorf("событие не было найдено")
}

func (s *calendar) EventsForWeek(date time.Time) (events []model.Event, err error) {
	// Высчитываем начало и конец недели
	startOfWeek := date.AddDate(0, 0, -int(date.Weekday())+1)
	endOfWeek := startOfWeek.AddDate(0, 0, 6)

	for currentDay := startOfWeek; currentDay.Before(endOfWeek.AddDate(0, 0, 1)); currentDay = currentDay.AddDate(0, 0, 1) {
		events = append(events, emulationDB[currentDay]...)
	}

	if events == nil {
		return events, fmt.Errorf("не было найдено событий")
	}

	return events, nil
}

func (s *calendar) EventsForMonth(date time.Time) (events []model.Event, err error) {
	// Высчитываем начало и конец недели
	startOfMonth := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.UTC)
	endOfMonth := startOfMonth.AddDate(0, 1, -1)

	for currentDay := startOfMonth; currentDay.Before(endOfMonth.AddDate(0, 0, 1)); currentDay = currentDay.AddDate(0, 0, 1) {
		events = append(events, emulationDB[currentDay]...)
	}

	if events == nil {
		return events, fmt.Errorf("не было найдено событий")
	}

	return events, nil
}

func NewServiceCalendar() service.Calendar {
	return &calendar{}
}
