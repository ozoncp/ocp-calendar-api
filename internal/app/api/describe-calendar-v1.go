package api

import (
	"context"
	"database/sql"
	"errors"
	desc "github.com/ozoncp/ocp-calendar-api/pkg/ocp-calendar-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *OcpCalendarApi) DescribeCalendarV1(
	ctx context.Context,
	req *desc.DescribeCalendarRequestV1,
) (*desc.DescribeCalendarResponseV1, error) {
	calendar, err := a.repo.DescribeCalendar(ctx, req.Id)
	log.Info().Msgf("Describe calendar attempt with ID: %d.", req.Id)

	if errors.Is(err, sql.ErrNoRows) {
		err = status.Error(codes.NotFound, NotFoundMessage)
		log.Error().Msgf(err.Error())

		return nil, err
	}

	return &desc.DescribeCalendarResponseV1{
		Id: calendar.Id,
		UserId: calendar.UserId,
		Type: calendar.Type,
		Link: calendar.Link,
	}, err
}
