package utils

import (
	"github.com/ozoncp/ocp-calendar-api/internal/app/models"
)

func MapIdentifiers(items []models.Calendar) map[uint64]models.Calendar {
	mappedEntities := map[uint64]models.Calendar{}

	for _, calendar := range items {
		mappedEntities[calendar.Id] = calendar
	}

	return mappedEntities
}
