package search

import "golang.org/x/exp/constraints"

func Binary[T constraints.Ordered](slice []T, value T) int {
	lo := 0
	hi := len(slice) - 1

	for {
		m := lo + (hi - lo) / 2
		v := slice[m]

		if (value == v) {
			return m
		}

		if lo >= hi {
			return -1
		}
		
		if (value < v) {
			hi = m
		} else {
			lo = m + 1
		}
	}
}
