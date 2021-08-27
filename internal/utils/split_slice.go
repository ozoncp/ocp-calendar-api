package utils

import (
	"github.com/ozoncp/ocp-calendar-api/internal/app/models"
)

func SplitSlice(items []models.Calendar, chunkSize int) [][]models.Calendar {
	var targetSlice [][]models.Calendar
	itemsLength := len(items)

	for startIndex := 0; startIndex < itemsLength; startIndex += chunkSize {
		var chunk []models.Calendar

		chunkLength := startIndex + chunkSize

		if itemsLength > chunkLength {
			chunk = items[startIndex:chunkLength]
		} else {
			chunk = items[startIndex:]
		}

		targetSlice = append(targetSlice, chunk)
	}

	return targetSlice
}
