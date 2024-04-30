package search

import (
	"testing"
)

func TestCrystal(t *testing.T) {
	tests := []struct {
		name     string
		breaks   []bool
		expected int
	}{
		{
			name:     "Test case 1a",
			breaks:   []bool{false, false, false, true, true, true, true},
			expected: 3,
		},
		{
			name:     "Test case 1b",
			breaks:   []bool{false, false, false, false, true, true, true, true},
			expected: 4,
		},
		{
			name:     "Test case 1c",
			breaks:   []bool{false, false, false, false, true, true, true},
			expected: 4,
		},
		{
			name:     "Test case 2a",
			breaks:   []bool{false, false, false, false, false, false, false, true},
			expected: 7,
		},
		{
			name:     "Test case 2b",
			breaks:   []bool{false, false, false, false, false, false, true},
			expected: 6,
		},
		{
			name:     "Test case 4a",
			breaks:   []bool{true, true, true, true, true, true, true, true},
			expected: 0,
		},
		{
			name:     "Test case 4b",
			breaks:   []bool{true, true, true, true, true, true, true},
			expected: 0,
		},
		{
			name:     "Test case 5a",
			breaks:   []bool{false, false, false, false, false, false, false, false},
			expected: -1,
		},
		{
			name:     "Test case 5b",
			breaks:   []bool{false, false, false, false, false, false, false},
			expected: -1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Crystal(test.breaks)
			if result != test.expected {
				t.Errorf("Expected %d, but got %d", test.expected, result)
			}
		})
	}
}
