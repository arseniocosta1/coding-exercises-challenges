package main

import (
	"testing"
)

var tests = []struct {
	s, want string
}{
	{"1", "1"},
	{"12", "12"},
	{"123", "123"},
	{"1234", "1,234"},
	{"12345", "12,345"},
	{"123456", "123,456"},
	{"1234567", "1,234,567"},
	{"12345678", "12,345,678"},
	{"123456789", "123,456,789"},
	{"1234567890", "1,234,567,890"},
	{"1.1234", "1.1234"},
	{"12.1234", "12.1234"},
	{"123.1234", "123.1234"},
	{"1234.1234", "1,234.1234"},
	{"12345.1234", "12,345.1234"},
	{"123456.1234", "123,456.1234"},
	{"1234567.1234", "1,234,567.1234"},
	{"12345678.1234", "12,345,678.1234"},
	{"123456789.1234", "123,456,789.1234"},
	{"1234567890.1234", "1,234,567,890.1234"},
	{"+123456789.1234", "+123,456,789.1234"},
	{"-1234567890.1234", "-1,234,567,890.1234"},
}

func TestComma1(t *testing.T) {
	for _, test := range tests {
		got := comma1(test.s)
		if got != test.want {
			t.Errorf("comma1(%q), got %q, want %q", test.s, got, test.want)
		}
	}
}

func TestComma2(t *testing.T) {
	for _, test := range tests {
		got := comma1(test.s)
		if got != test.want {
			t.Errorf("comma1(%q), got %q, want %q", test.s, got, test.want)
		}
	}
}

func TestComma3(t *testing.T) {
	for _, test := range tests {
		got := comma1(test.s)
		if got != test.want {
			t.Errorf("comma1(%q), got %q, want %q", test.s, got, test.want)
		}
	}
}

func BenchmarkComma1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			comma1(test.s)
		}
	}
}

func BenchmarkComma2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			comma2(test.s)
		}
	}
}

func BenchmarkComma3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			comma3(test.s)
		}
	}
}
