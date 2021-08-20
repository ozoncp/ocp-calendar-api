package ocp_calendar_api

import (
	"context"
	desc "github.com/ozoncp/ocp-calendar-api/pkg/ocp-calendar-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *OcpCalendarApi) DescribeCalendarV1(ctx context.Context, req *desc.DescribeCalendarRequestV1) (*emptypb.Empty, error) {
	log.Info().Msgf("Describe calendar attempt with ID: %d.", req.Id)
	return &emptypb.Empty{}, nil
}
