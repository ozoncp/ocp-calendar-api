package entities

import "time"

type CalendarEvent struct {
	Id uint64
	Calendar Calendar
	Title string
	Description string
	StartAt time.Time
	EndAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
