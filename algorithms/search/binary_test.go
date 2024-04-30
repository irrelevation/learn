package search

import (
	"testing"
)

func TestBinary(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int // Replace int with the type of your slice elements
		value    int   // Replace int with the type of your value
		expected int
	}{
		{
			name:     "Test middle",
			slice:    []int{1, 2, 3, 4, 5},
			value:    3,
			expected: 2,
		},
		{
			name:     "Test start",
			slice:    []int{1, 2, 3, 4, 5},
			value:    1,
			expected: 0,
		},
		{
			name:     "Test end",
			slice:    []int{1, 2, 3, 4, 5},
			value:    5,
			expected: 4,
		},
		{
			name:     "Test not found bigger",
			slice:    []int{1, 2, 3, 4, 5},
			value:    6,
			expected: -1,
		},
		{
			name:     "Test not found smaller",
			slice:    []int{1, 2, 3, 4, 5},
			value:    -6,
			expected: -1,
		},
		{
			name:     "Test not found middle",
			slice:    []int{2, 4, 6, 8, 10},
			value:    3,
			expected: -1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Binary(test.slice, test.value)
			if result != test.expected {
				t.Errorf("Expected %d, but got %d", test.expected, result)
			}
		})
	}
}