package sort

import (
	"reflect"
	"testing"
)

func TestBubble(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Test case 1",
			input:    []int{4, 2, 7, 1, 5},
			expected: []int{1, 2, 4, 5, 7},
		},
		{
			name:     "Test case 2",
			input:    []int{9, 3, 6, 2, 8},
			expected: []int{2, 3, 6, 8, 9},
		},
		{
			name:     "Test case 3",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Test case 4",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			Bubble(test.input)
			if !reflect.DeepEqual(test.input, test.expected) {
				t.Errorf("Expected %v, but got %v", test.expected, test.input)
			}
		})
	}
}
