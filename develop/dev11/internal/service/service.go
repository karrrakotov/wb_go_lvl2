package service

import (
	"time"

	"wb_go_lvl2/dev11/internal/model"
)

type Calendar interface {
	Create(event model.Event) error
	Update(event model.Event) error
	Delete(event model.Event) error
	EventsForDay(date time.Time) (events []model.Event, err error)
	EventsForWeek(date time.Time) (events []model.Event, err error)
	EventsForMonth(date time.Time) (events []model.Event, err error)
}
