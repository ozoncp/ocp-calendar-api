package api

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/ozoncp/ocp-calendar-api/internal/app/models"
	"github.com/ozoncp/ocp-calendar-api/internal/flusher"
	desc "github.com/ozoncp/ocp-calendar-api/pkg/ocp-calendar-api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *OcpCalendarApi) MultiCreateCalendarV1(
	ctx context.Context,
	req *desc.MultiCreateCalendarRequestV1,
) (*emptypb.Empty, error) {
	calendars := getCalendarsFromRequest(req)
	msg, err := PrepareMessage("MultiCreateCalendars", calendars)

	if err != nil {
		return &emptypb.Empty{}, err
	}

	if err := req.Validate(); err != nil {
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span := opentracing.GlobalTracer().StartSpan("MultiCreateCalendarV1")
	span.SetTag("Id", req.Id)
	span.SetTag("size", len(req.Items))
	defer span.Finish()

	flusher.NewFlusher(
		ctx,
		span,
		2,
		a.repo,
	).Flush(getCalendarsFromRequest(req))

	_, _, err = a.producer.SendMessage(msg)

	if err != nil {
		return &emptypb.Empty{}, err
	}

	CreateCalendarsCounter.WithLabelValues(SuccessLabelValue).Inc()

	return &emptypb.Empty{}, nil
}

// Читает в запросе `items` и возвращает как слайс `models.Calendar`
func getCalendarsFromRequest(req *desc.MultiCreateCalendarRequestV1) []models.Calendar {
	var calendars []models.Calendar

	for _, requestCalendar := range req.Items {
		calendars = append(calendars, models.Calendar{
			UserId: requestCalendar.UserId,
			Type:   requestCalendar.Type,
			Link:   requestCalendar.Link,
		})
	}

	return calendars
}
