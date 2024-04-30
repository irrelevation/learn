package search

import (
	"math"
)

// Crystal is a function that determines the first position from wich the crystal ball breaks.
// It takes a slice of boolean values. values are false until they are true. they are not mixed.
// It returns the index of the first break found, or -1 if no breaks are found.
func Crystal(breaks []bool) int {
	step := int(math.Sqrt(float64(len(breaks))))
	i := step
	for ; i < len(breaks); i += step {
		if breaks[i] {
			break
		}
	}

	for j := i - step; j <= i && j < len(breaks); j++ {
		if breaks[j] {
			return j
		}
	}

	return -1
}
