package utils

func PickUnique(from []int, compare map[int]int) []int {
	var uniqueValues []int

	for _, value := range from {
		if _, found := compare[value]; !found {
			uniqueValues = append(uniqueValues, value)
		}
	}

	return uniqueValues
}
