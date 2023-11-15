package model

import "time"

// Event представляет событие в календаре
type Event struct {
	UserID int       `json:"user_id"`
	Date   time.Time `json:"date"`
	Title  string    `json:"title"`
}
