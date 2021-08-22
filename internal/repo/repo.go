package repo

import (
	"context"
	"github.com/ozoncp/ocp-calendar-api/internal/app/models"
)

type Repo interface {
	AddCalendars(ctx context.Context, calendars []models.Calendar) error
	ListCalendars(ctx context.Context, limit, offset, userId, calendarType uint64) ([]models.Calendar, error)
	DescribeCalendar(ctx context.Context, calendarId uint64) (models.Calendar, error)
	RemoveCalendar(ctx context.Context, calendarId uint64) error
}
