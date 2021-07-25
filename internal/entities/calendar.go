package entities

import (
	"encoding/json"
	"time"
)

type Events map[uint64]CalendarEvent

type Calendar struct {
	Id        int64     `json:"id"`
	UserId    int64     `json:"user_id"`
	Name      string    `json:"name"`
	Events    Events    `json:"events,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (calendar Calendar) ToJson() (string, error) {
	result, err := json.MarshalIndent(calendar, "", "    ")
	return string(result), err
}

func (calendar Calendar) AddEvent(event CalendarEvent) error {
	// this is a draft
	// @todo: implementation
	return nil
}

func (calendar Calendar) DeleteEvent(eventId uint64) error {
	// this is a draft
	// @todo: implementation
	return nil
}

func (calendar Calendar) FindEvents(from time.Time, to time.Time) (Events, error) {
	// this is a draft
	// @todo: implementation
	return make(map[uint64]CalendarEvent), nil
}
