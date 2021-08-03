package repo

import "github.com/ozoncp/ocp-calendar-api/app/models"

type Repo interface {
	AddCalendars(calendar []models.Calendar) error
	ListCalendars(limit, offset uint64) ([]models.Calendar, error)
	DescribeCalendar(calendarId uint64) (models.Calendar, error)
}
