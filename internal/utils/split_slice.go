package utils

import (
	"github.com/ozoncp/ocp-calendar-api/internal/entities"
	"math"
)

func expectedSliceCapacity(sliceLen, chunkSize int) int {
	chunkCount := int(math.Ceil(float64(sliceLen) / float64(chunkSize)))
	return chunkSize * chunkCount
}

func minCapacity(capacity1, capacity2 int) int {
	return int(math.Min(float64(capacity1), float64(capacity2)))
}

func SplitSlice(items []entities.Calendar, chunkSize int) [][]entities.Calendar {
	var targetSlice [][]entities.Calendar

	expectedCapacity := expectedSliceCapacity(len(items), chunkSize)
	sourceSliceCapacity := cap(items)

	for i := 0; i < expectedCapacity; i += chunkSize {
		chunk := items[i:minCapacity(i+chunkSize, sourceSliceCapacity)]
		targetSlice = append(targetSlice, chunk)
	}

	return targetSlice
}

