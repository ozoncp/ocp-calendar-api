package utils

func SwapKeyValue(sourceMap map[int]int) map[int]int {
	swappedMap := map[int]int{}

	for key, value := range sourceMap {
		swappedMap[value] = key
	}

	return swappedMap
}
