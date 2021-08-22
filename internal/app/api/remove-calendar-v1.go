package api

import (
	"context"
	desc "github.com/ozoncp/ocp-calendar-api/pkg/ocp-calendar-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *OcpCalendarApi) RemoveCalendarV1(
	ctx context.Context,
	req *desc.RemoveCalendarRequestV1,
) (*emptypb.Empty, error) {
	err := a.repo.RemoveCalendar(ctx, req.Id)
	log.Info().Msgf("Remove calendar attempt with ID: %d", req.Id)

	return &emptypb.Empty{}, err
}
