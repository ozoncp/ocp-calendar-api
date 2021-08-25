package flusher

import (
	"context"
	"github.com/ozoncp/ocp-calendar-api/internal/app/models"
	"github.com/ozoncp/ocp-calendar-api/internal/repo"
	"github.com/ozoncp/ocp-calendar-api/internal/utils"
)

type Flusher interface {
	Flush(calendars []models.Calendar) []models.Calendar
}

type flusher struct {
	chunkSize    int
	calendarRepo repo.Repo
	ctx          context.Context
}

func (f *flusher) Flush(calendars []models.Calendar) []models.Calendar {
	for _, calendarsChunk := range utils.SplitSlice(calendars, f.chunkSize) {
		if err := f.calendarRepo.AddCalendars(f.ctx, calendarsChunk); err != nil {
			panic(err)
		}
	}

	return make([]models.Calendar, 0)
}

func NewFlusher(ctx context.Context, chunkSize int, repo repo.Repo) Flusher {
	return &flusher{
		chunkSize:    chunkSize,
		calendarRepo: repo,
		ctx:          ctx,
	}
}
