package api

import (
	"context"
	desc "github.com/ozoncp/ocp-calendar-api/pkg/ocp-calendar-api"
	"github.com/rs/zerolog/log"
)

func (a *OcpCalendarApi) ListCalendarsV1(
	ctx context.Context,
	req *desc.ListCalendarRequestV1,
) (*desc.ListCalendarResponseV1, error) {
	var response desc.ListCalendarResponseV1

	calendars, err := a.repo.ListCalendars(ctx, req.Limit, req.Offset, req.UserId, req.Type)

	for _, calendar := range calendars {
		response.Items = append(response.Items, &desc.DescribeCalendarResponseV1{
			Id: calendar.Id,
			UserId: calendar.UserId,
			Type: calendar.Type,
			Link: calendar.Link,
		})
	}

	log.Info().Msgf("List calendars attempt.")

	return &response, err
}
