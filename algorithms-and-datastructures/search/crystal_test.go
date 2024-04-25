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
			name:     "Test case 1",
			breaks:   []bool{false, false, false, false, true, true, true, true},
			expected: 4,
		},
		{
			name:     "Test case 2",
			breaks:   []bool{false, false, false, false, false, false, false, true},
			expected: 7,
		},
		{
			name:     "Test case 3",
			breaks:   []bool{true, true, true, true, true, true, true, true},
			expected: 0,
		},
		{
			name:     "Test case 4",
			breaks:   []bool{false, false, false, false, false, false, false, false},
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