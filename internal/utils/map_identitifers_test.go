package utils

import (
	"github.com/ozoncp/ocp-calendar-api/internal/entities"
	"testing"
	"time"
)

func TestMapIdentifiers(t *testing.T) {
	items := []entities.Calendar{{
		Id:        1,
		UserId:    1,
		Name:      "Calendar #1",
		Events:    nil,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}, {
		Id:        2,
		UserId:    1,
		Name:      "Calendar #2",
		Events:    nil,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}}
	mappedItems := MapIdentifiers(items)

	if len(mappedItems) != len(items) {
		t.Errorf(
			"Calendar slice size (%v) and calendar mapped size (%v) is not equal",
			len(mappedItems),
			len(items),
		)
	}

	for id, calendar := range mappedItems {
		if calendar.Id != id {
			t.Errorf("Calendar #%v (id: %v) is not mapped correctly", id, id)
		}
	}
}
