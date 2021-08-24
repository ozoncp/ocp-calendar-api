package api

import (
	"context"
	"github.com/ozoncp/ocp-calendar-api/internal/app/models"
	desc "github.com/ozoncp/ocp-calendar-api/pkg/ocp-calendar-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *OcpCalendarApi) CreateCalendarV1(
	ctx context.Context,
	req *desc.CreateCalendarRequestV1,
) (*emptypb.Empty, error) {
	calendar := models.Calendar{
		UserId: req.UserId,
		Type:   req.Type,
		Link:   req.Link,
	}
	msg, err := PrepareMessage("CreateCalendar", []models.Calendar{calendar})

	if err := req.Validate(); err != nil {
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, err.Error())
	}

	err = a.repo.AddCalendars(ctx, []models.Calendar{calendar})
	log.Info().Msg("Create calendar.")
	_, _, err = a.producer.SendMessage(msg)
	CreateCalendarCounter.WithLabelValues(SuccessLabelValue).Inc()

	return &emptypb.Empty{}, err
}
