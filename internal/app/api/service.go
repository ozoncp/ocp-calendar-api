package api

import (
	"github.com/ozoncp/ocp-calendar-api/internal/repo"
	desc "github.com/ozoncp/ocp-calendar-api/pkg/ocp-calendar-api"
)

const NotFoundMessage = "Calendar not found"

type OcpCalendarApi struct {
	desc.UnimplementedOcpCalendarApiServer
	repo repo.Repo
}

func NewOcpCalendarApi(repo repo.Repo) desc.OcpCalendarApiServer {
	return &OcpCalendarApi{repo: repo}
}