package utils

import (
	"github.com/ozoncp/ocp-calendar-api/internal/entities"
	"testing"
	"time"
)

func TestSplitSlice(t *testing.T) {
	chunkSize := 3
	calendarSlice := []entities.Calendar{{
		Id:        0,
		UserId:    0,
		Name:      "",
		Events:    nil,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}, {
		Id:        0,
		UserId:    0,
		Name:      "",
		Events:    nil,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}, {
		Id:        0,
		UserId:    0,
		Name:      "",
		Events:    nil,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}, {
		Id:        0,
		UserId:    0,
		Name:      "",
		Events:    nil,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
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
