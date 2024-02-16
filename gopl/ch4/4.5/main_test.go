package main

import (
	"reflect"
	"testing"
)

func TestRemoveAdjacentDupes(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "No duplicates",
			input:    []string{"a", "b", "c"},
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "Adjacent duplicates",
			input:    []string{"a", "a", "b", "b", "c", "c"},
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "Non-adjacent duplicates",
			input:    []string{"a", "b", "a", "c", "b", "c"},
			expected: []string{"a", "b", "a", "c", "b", "c"},
		},
		{
			name:     "Empty slice",
			input:    []string{},
			expected: []string{},
		},
		{
			name:     "Single element",
			input:    []string{"a"},
			expected: []string{"a"},
		},
		{
			name:     "All duplicates",
			input:    []string{"a", "a", "a", "a"},
			expected: []string{"a"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Make a copy of the input slice to compare with the expected output
			inputCopy := make([]string, len(test.input))
			copy(inputCopy, test.input)

			removeAdjacentDupes(&inputCopy)

			if !reflect.DeepEqual(inputCopy, test.expected) {
				t.Errorf("unexpected result, got: %v, want: %v", inputCopy, test.expected)
			}
		})
	}
}
