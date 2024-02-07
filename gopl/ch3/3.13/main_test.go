package main

import "testing"

var tcs = []struct {
	s1      string
	s2      string
	expects bool
}{
	{"abc", "cba", true},
	{"abc", "abc", true},
	{"abc", "abe", false},
	{"abc", "abcd", false},
	{"abc", "ab", false},
	{"abc", "", false},
	{"listen", "silent", true},
	{"triangle", "integral", true},
}

func TestIsAnagram(t *testing.T) {
	for _, tc := range tcs {
		ret := anagram(tc.s1, tc.s2)
		if ret != tc.expects {
			t.Errorf("anagram(%q, %q), Wanted: %v, Got: %v", tc.s1, tc.s2, tc.expects, ret)
		}
	}
}

func TestIsAnagramSlicesSort(t *testing.T) {
	for _, tc := range tcs {
		ret := anagramSlicesSort(tc.s1, tc.s2)
		if ret != tc.expects {
			t.Errorf("anagramSlicesSort(%q, %q), Wanted: %v, Got: %v", tc.s1, tc.s2, tc.expects, ret)
		}
	}
}

func TestIsAnagramMap(t *testing.T) {
	for _, tc := range tcs {
		ret := anagramMap(tc.s1, tc.s2)
		if ret != tc.expects {
			t.Errorf("anagramMap(%q, %q), Wanted: %v, Got: %v", tc.s1, tc.s2, tc.expects, ret)
		}
	}
}

func BenchmarkIsAnagram(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			anagram(tc.s1, tc.s2)
		}
	}
}

func BenchmarkIsAnagramSlicesSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			anagramSlicesSort(tc.s1, tc.s2)
		}
	}
}

func BenchmarkIsAnagramMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			anagramMap(tc.s1, tc.s2)
		}
	}
}
