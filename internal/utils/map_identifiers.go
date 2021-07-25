package utils

import "github.com/ozoncp/ocp-calendar-api/internal/entities"

func MapIdentifiers(items []entities.Calendar) map[uint64]entities.Calendar {
	mappedEntities := map[uint64]entities.Calendar{}

	for _, calendar := range items {
		mappedEntities[calendar.Id] = calendar
	}

	return mappedEntities
}
