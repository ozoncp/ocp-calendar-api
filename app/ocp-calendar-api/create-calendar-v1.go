package ocp_calendar_api

import (
	"context"
	desc "github.com/ozoncp/ocp-calendar-api/pkg/ocp-calendar-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *OcpCalendarApi) CreateCalendarV1(ctx context.Context, req *desc.CreateCalendarRequestV1) (*emptypb.Empty, error) {
	log.Info().Msg("Create calendar attempt.")
	return &emptypb.Empty{}, nil
}
