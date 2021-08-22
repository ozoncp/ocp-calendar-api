package repo

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/ozoncp/ocp-calendar-api/internal/app/models"
	"github.com/ozoncp/ocp-calendar-api/internal/pkg/db"
)

var tableName = "calendars"

type CalendarRepository struct {}

func (c CalendarRepository) RemoveCalendar(ctx context.Context, calendarId uint64) error {
	_, err := sq.Delete(tableName).
		Where("Id = ?", calendarId).
		RunWith(db.GetDB(ctx)).
		PlaceholderFormat(sq.Dollar).
		ExecContext(ctx)

	return err
}

func (c CalendarRepository) AddCalendars(ctx context.Context, calendars []models.Calendar) error {
	var err error

	for _, calendar := range calendars {
		_, err := sq.Insert(tableName).
			Columns("UserId", "Type", "Link").
			Values(calendar.UserId, calendar.Type, calendar.Link).
			RunWith(db.GetDB(ctx)).
			PlaceholderFormat(sq.Dollar).
			ExecContext(ctx)

		if err != nil {
			return err
		}
	}

	return err
}

func (c CalendarRepository) ListCalendars(
	ctx context.Context,
	limit, offset, userId, calendarType uint64,
) ([]models.Calendar, error) {
	var calendars []models.Calendar

	query := sq.Select("Id", "UserId", "Type", "Link").
		From(tableName).
		OrderBy("Id")

	if limit > 0 {
		query = query.Limit(limit)
	}

	if offset > 0 {
		query = query.Offset(offset)
	}

	if userId > 0 {
		query = query.Where("UserId = ?", userId)
	}

	if calendarType > 0 {
		query = query.Where("Type = ?", calendarType)
	}

	rows, err := query.RunWith(db.GetDB(ctx)).
		PlaceholderFormat(sq.Dollar).
		QueryContext(ctx)

	for rows.Next() {
		var calendar models.Calendar

		err := rows.Scan(
			&calendar.Id,
			&calendar.UserId,
			&calendar.Type,
			&calendar.Link,
		)

		if err != nil {
			return calendars, err
		}

		calendars = append(calendars, calendar)
	}

	return calendars, err
}

func (c CalendarRepository) DescribeCalendar(ctx context.Context, calendarId uint64) (models.Calendar, error) {
	var calendar models.Calendar

	err := sq.Select("*").
		From(tableName).
		Where("Id = ?", calendarId).
		RunWith(db.GetDB(ctx)).
		PlaceholderFormat(sq.Dollar).
		QueryRowContext(ctx).
		Scan(
			&calendar.Id,
			&calendar.UserId,
			&calendar.Type,
			&calendar.Link,
		)

	return calendar, err
}

func NewCalendarRepository() Repo {
	return &CalendarRepository{}
}
