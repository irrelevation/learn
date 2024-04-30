package search

import (
	"testing"
)

func TestLinear(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int // Replace int with the type of your slice elements
		value    int   // Replace int with the type of your value
		expected int
	}{
		{
			name:     "Test case 1",
			slice:    []int{1, 2, 3, 4, 5},
			value:    3,
			expected: 2,
		},
		{
			name:     "Test case 2",
			slice:    []int{1, 2, 3, 4, 5},
			value:    6,
			expected: -1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Linear(test.slice, test.value)
			if result != test.expected {
				t.Errorf("Expected %d, but got %d", test.expected, result)
			}
		})
	}
}