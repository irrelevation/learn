package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func main() {
	a := []int{0, 2, 1}
	fmt.Println(Blubb(a))
	fmt.Println(a)
}

func Blubb[T constraints.Ordered](slice []T) []T {
	fmt.Printf("sorting: %v", slice)
	fmt.Println()
	for i := 0; i < len(slice)-1; i++ {
		fmt.Printf("i: %v", i)
		fmt.Println()
		if slice[i] > slice[i+1] {
			fmt.Printf("before: %v", slice)
			fmt.Println()

			slice[i], slice[i+1] = slice[i+1], slice[i]
			fmt.Printf("after: %v", slice)
			fmt.Println()

		}
	}

	return slice
}
