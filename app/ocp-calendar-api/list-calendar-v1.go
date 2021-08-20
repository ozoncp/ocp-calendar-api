package ocp_calendar_api

import (
	"context"
	desc "github.com/ozoncp/ocp-calendar-api/pkg/ocp-calendar-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *OcpCalendarApi) ListCalendarsV1(ctx context.Context, req *desc.ListCalendarRequestV1) (*emptypb.Empty, error) {
	log.Info().Msgf("List calendars attempt.")
	return &emptypb.Empty{}, nil
}
