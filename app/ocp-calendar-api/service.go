package ocp_calendar_api

import desc "github.com/ozoncp/ocp-calendar-api/pkg/ocp-calendar-api"

type OcpCalendarApi struct {
	desc.UnimplementedOcpCalendarApiServer
}

func NewOcpCalendarApi() desc.OcpCalendarApiServer {
	return &OcpCalendarApi{}
}
