package flusher

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/ozoncp/ocp-calendar-api/internal/app/models"
	"github.com/ozoncp/ocp-calendar-api/internal/repo"
	"github.com/ozoncp/ocp-calendar-api/internal/utils"
	"github.com/rs/zerolog/log"
)

type Flusher interface {
	Flush(calendars []models.Calendar) []models.Calendar
}

type flusher struct {
	chunkSize    int
	calendarRepo repo.Repo
	ctx          context.Context
	parentSpan   opentracing.Span
}

func (f *flusher) Flush(calendars []models.Calendar) []models.Calendar {
	for batchId, calendarsChunk := range utils.SplitSlice(calendars, f.chunkSize) {
		if err := f.flushChunk(f.ctx, batchId, calendarsChunk); err != nil {
			log.Error().Msg(err.Error())
		}
	}

	return make([]models.Calendar, 0)
}

func (f *flusher) flushChunk(ctx context.Context, batchId int, calendars []models.Calendar) error {
	span := opentracing.StartSpan(
		"FlushCalendarsChunk",
		opentracing.ChildOf(f.parentSpan.Context()))
	span.SetTag("BatchId", batchId)
	span.SetTag("size", len(calendars))
	defer span.Finish()

	if err := f.calendarRepo.AddCalendars(ctx, calendars); err != nil {
		return err
	}

	return nil
}

func NewFlusher(
	ctx context.Context,
	parentSpan opentracing.Span,
	chunkSize int,
	repo repo.Repo,
) Flusher {
	return &flusher{
		chunkSize:    chunkSize,
		calendarRepo: repo,
		ctx:          ctx,
		parentSpan:   parentSpan,
	}
}
