package api

import (
	"context"
	"github.com/ozoncp/ocp-calendar-api/internal/app/models"
	desc "github.com/ozoncp/ocp-calendar-api/pkg/ocp-calendar-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *OcpCalendarApi) RemoveCalendarV1(
	ctx context.Context,
	req *desc.RemoveCalendarRequestV1,
) (*emptypb.Empty, error) {
	calendar := models.Calendar{Id: req.Id}
	msg, err := PrepareMessage("RemoveCalendar", []models.Calendar{calendar})

	if err == nil {
		err = a.repo.RemoveCalendar(ctx, req.Id)
		if err == nil {
			log.Info().Msgf("calendar ID: %d deleted", req.Id)
			_, _, err = a.producer.SendMessage(msg)
			RemoveCalendarCounter.WithLabelValues(SuccessLabelValue).Inc()
		}
	}

	return &emptypb.Empty{}, err
}
