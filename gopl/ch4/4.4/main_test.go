package main

import (
	"slices"
	"testing"
)

func TestRotate(t *testing.T) {
	var tests = []struct {
		name     string
		a        []int
		n        int
		expected []int
	}{
		{
			name:     "No rotation",
			a:        []int{1, 2, 3, 4, 5},
			n:        0,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Positive rotation",
			a:        []int{1, 2, 3, 4, 5},
			n:        1,
			expected: []int{5, 1, 2, 3, 4},
		},
		{
			name:     "Negative rotation",
			a:        []int{1, 2, 3, 4, 5},
			n:        -2,
			expected: []int{3, 4, 5, 1, 2},
		},
		{
			name:     "Rotation greater than length",
			a:        []int{1, 2, 3, 4, 5},
			n:        7,
			expected: []int{4, 5, 1, 2, 3},
		},
		{
			name:     "Rotation less than negative length",
			a:        []int{1, 2, 3, 4, 5},
			n:        -8,
			expected: []int{4, 5, 1, 2, 3},
		},
		{
			name:     "Rotation multiple of length",
			a:        []int{1, 2, 3, 4, 5},
			n:        10,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Negative rotation multiple of length",
			a:        []int{1, 2, 3, 4, 5},
			n:        -15,
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Make a copy of the original slice since the rotate function modifies the original slice in place
			aCopy := make([]int, len(test.a))
			copy(aCopy, test.a)

			rotate(test.a, test.n)
			if !slices.Equal(test.a, test.expected) {
				t.Errorf("rotate(%v, %d) = %v, want %v", aCopy, test.n, test.a, test.expected)
			}
		})
	}
}
