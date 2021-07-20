package main

import (
	"fmt"
	"github.com/ozoncp/ocp-calendar-api/internal/utils"
)

func main() {
	fmt.Println("1. Chunked:\t", utils.SplitSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11}, 3))
	fmt.Println("2. Swapped:\t", utils.SwapKeyValue(map[int]int{
		0: 1,
		1: 2,
		2: 3,
	}))
	fmt.Println("3. Filtered:\t", utils.PickUnique([]int{1, 2, 3, 5}, map[int]int{1: 1, 3: 1}))
}
