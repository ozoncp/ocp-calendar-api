package api

import (
	"context"
	"database/sql"
	"errors"
	"github.com/ozoncp/ocp-calendar-api/internal/app/models"
	desc "github.com/ozoncp/ocp-calendar-api/pkg/ocp-calendar-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *OcpCalendarApi) UpdateCalendarV1(
	ctx context.Context,
	req *desc.UpdateCalendarRequestV1,
) (*desc.DescribeCalendarResponseV1, error) {
	calendarUpdates := models.Calendar{
		Id:     req.Id,
		UserId: req.UserId,
		Type:   req.Type,
		Link:   req.Link,
	}
	response := &desc.DescribeCalendarResponseV1{
		Id:     req.Id,
		UserId: req.UserId,
		Type:   req.Type,
		Link:   req.Link,
	}
	msg, err := PrepareMessage("UpdateCalendar", []models.Calendar{calendarUpdates})

	if err != nil {
		return response, err
	}

	if err = req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	calendar, err := a.repo.UpdateCalendar(ctx, calendarUpdates)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = status.Error(codes.NotFound, NotFoundMessage)
			log.Error().Msgf(err.Error())
		}

		return nil, err
	}

	log.Info().Msgf("calendar ID: %v updated", calendar.Id)

	_, _, err = a.producer.SendMessage(msg)

	UpdateCalendarCounter.WithLabelValues(SuccessLabelValue).Inc()

	return response, nil
}
