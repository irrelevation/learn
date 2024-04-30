package sort

import (
	"golang.org/x/exp/constraints"
)

func Bubble[T constraints.Ordered](slice []T) {
	for top := len(slice) - 1; top > 0; top-- {
		for i := 0; i < top; i++ {
			if slice[i] > slice[i+1] {
				slice[i], slice[i+1] = slice[i+1], slice[i]
			}
		}
	}
}
