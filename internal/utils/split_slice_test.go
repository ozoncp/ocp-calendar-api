package utils

import (
	"github.com/ozoncp/ocp-calendar-api/app/models"
	"testing"
)

func TestSplitSlice(t *testing.T) {
	chunkSize := 3
	calendarSlice := []models.Calendar{{
		Id:     0,
		UserId: 0,
		Type:   1,
		Link:   "https://some.io/event/1",
	}, {
		Id:     0,
		UserId: 0,
		Type:   1,
		Link:   "https://some.io/event/1",
	}, {
		Id:     0,
		UserId: 0,
		Type:   1,
		Link:   "https://some.io/event/1",
	}, {
		Id:     0,
		UserId: 0,
		Type:   1,
		Link:   "https://some.io/event/1",
	}}
	chunks := SplitSlice(calendarSlice, chunkSize)

	if len(chunks) != 2 {
		t.Error("Chunks length should be equal 2")
	}

	if (len(chunks[0])) != 3 {
		t.Errorf("Wrong 1st chunk size, should be equal %v, got %v", chunkSize, len(chunks[0]))
	}

	if (len(chunks[1])) != 1 {
		t.Errorf("Tail chunk should contain 1 element, got %v", len(chunks[1]))
	}
}
