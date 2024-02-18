package main

import (
	"reflect"
	"testing"
)

func TestRemoveAdjacentDuplicateSpaces(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected []byte
	}{
		{
			name:     "No duplicates",
			input:    []byte("Hello World"),
			expected: []byte("Hello World"),
		},
		{
			name:     "Adjacent duplicates",
			input:    []byte("Hello   World"),
			expected: []byte("Hello World"),
		},
		{
			name:     "Multiple adjacent duplicates",
			input:    []byte("Hello   World   !"),
			expected: []byte("Hello World !"),
		},
		{
			name:     "Spaces at start",
			input:    []byte("   Hello World"),
			expected: []byte(" Hello World"),
		},
		{
			name:     "Spaces at end",
			input:    []byte("Hello World   "),
			expected: []byte("Hello World "),
		},
		{
			name:     "Newline characters",
			input:    []byte("Hello\n\nWorld"),
			expected: []byte("Hello\nWorld"),
		},
		{
			name:     "Tab characters",
			input:    []byte("Hello\t\tWorld"),
			expected: []byte("Hello\tWorld"),
		},
		{
			name:     "Mixed spaces and tabs",
			input:    []byte("Hello   \tWorld"),
			expected: []byte("Hello World"),
		},
		{
			name:     "Form feed characters",
			input:    []byte("Hello\f\fWorld"),
			expected: []byte("Hello\fWorld"),
		},
		{
			name:     "Carriage return characters",
			input:    []byte("Hello\r\rWorld"),
			expected: []byte("Hello\rWorld"),
		},
		{
			name:     "Vertical tab characters",
			input:    []byte("Hello\v\vWorld"),
			expected: []byte("Hello\vWorld"),
		},
		{
			name:     "NEL (NEXT LINE) character",
			input:    []byte("Hello\xc2\x85\xc2\x85World"),
			expected: []byte("Hello\xc2\x85World"),
		},
		{
			name: "NBSP (non-breaking space) character",
			// Note that NSB in ISO-8859-1 is \xA0 but in UTF-8 is \xC2\xA0
			// I could have used the \u00A0 escape sequence, but I wanted to show the UTF-8 encoding
			// of the NBSP character.
			input:    []byte("\xC2\xA0\xC2\xA0"),
			expected: []byte("\xC2\xA0"),
		},
		{
			name:     "All IsSpace characters",
			input:    []byte("Hello \t\n\f\xC2\xA0\xc2\x85\r\v World"),
			expected: []byte("Hello World"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Make a copy of the input to avoid modifying the original slice
			input := make([]byte, len(test.input))
			copy(input, test.input)

			RemoveAdjacentDuplicateSpaces(&input)

			if !reflect.DeepEqual(input, test.expected) {
				t.Errorf("Unexpected result. Got: %s, Expected: %s", input, test.expected)
			}
		})
	}
}

func BenchmarkRemoveAdjacentDuplicateSpaces(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := []byte("Hello\u00A0\u00A0World")
		RemoveAdjacentDuplicateSpaces(&s)
	}
}
