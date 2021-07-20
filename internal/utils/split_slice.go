package utils

import (
	"math"
)

func expectedSliceCapacity(sliceLen, chunkSize int) int {
	chunkCount := int(math.Ceil(float64(sliceLen) / float64(chunkSize)))
	return chunkSize * chunkCount
}

func minCapacity(capacity1, capacity2 int) int {
	return int(math.Min(float64(capacity1), float64(capacity2)))
}

func SplitSlice(sourceSlice []int, chunkSize int) [][]int {
	var targetSlice [][]int

	expectedCapacity := expectedSliceCapacity(len(sourceSlice), chunkSize)
	sourceSliceCapacity := cap(sourceSlice)

	for i := 0; i < expectedCapacity; i += chunkSize {
		chunk := sourceSlice[i:minCapacity(i+chunkSize, sourceSliceCapacity)]
		targetSlice = append(targetSlice, chunk)
	}

	return targetSlice
}
