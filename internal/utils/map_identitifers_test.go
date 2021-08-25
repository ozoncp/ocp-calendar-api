package utils

import (
	"github.com/ozoncp/ocp-calendar-api/internal/app/models"
	"testing"
)

func TestMapIdentifiers(t *testing.T) {
	items := []models.Calendar{{
		Id:     1,
		UserId: 1,
		Type:   1,
		Link:   "https://some.io/event/1",
	}, {
		Id:     2,
		UserId: 2,
		Type:   2,
		Link:   "https://some.io/event/2",
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
