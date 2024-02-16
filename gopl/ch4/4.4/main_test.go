package main

import (
	"slices"
	"testing"
)

func TestRotate(t *testing.T) {
	var tests = []struct {
		a, acopy []int
		n        int
		want     []int
	}{

		{a: []int{1, 2, 3, 4, 5}, acopy: []int{1, 2, 3, 4, 5}, n: 0, want: []int{1, 2, 3, 4, 5}},
		{a: []int{5, 4, 3, 2, 1}, acopy: []int{5, 4, 3, 2, 1}, n: 0, want: []int{5, 4, 3, 2, 1}},
		{a: []int{10, 20, 30, 40, 50}, acopy: []int{10, 20, 30, 40, 50}, n: 0, want: []int{10, 20, 30, 40, 50}},

		{a: []int{1, 2, 3, 4, 5}, acopy: []int{1, 2, 3, 4, 5}, n: 1, want: []int{5, 1, 2, 3, 4}},
		{a: []int{1, 2, 3, 4, 5}, acopy: []int{1, 2, 3, 4, 5}, n: 2, want: []int{4, 5, 1, 2, 3}},
		{a: []int{5, 4, 3, 2, 1}, acopy: []int{5, 4, 3, 2, 1}, n: 3, want: []int{3, 2, 1, 5, 4}},
		{a: []int{10, 20, 30, 40, 50}, acopy: []int{10, 20, 30, 40, 50}, n: 4, want: []int{20, 30, 40, 50, 10}},

		{a: []int{1, 2, 3, 4, 5}, acopy: []int{1, 2, 3, 4, 5}, n: -2, want: []int{3, 4, 5, 1, 2}},
		{a: []int{5, 4, 3, 2, 1}, acopy: []int{5, 4, 3, 2, 1}, n: -3, want: []int{2, 1, 5, 4, 3}},
		{a: []int{10, 20, 30, 40, 50}, acopy: []int{10, 20, 30, 40, 50}, n: -4, want: []int{50, 10, 20, 30, 40}},

		// n greater than the length of the slice
		{a: []int{1, 2, 3, 4, 5}, acopy: []int{1, 2, 3, 4, 5}, n: 7, want: []int{4, 5, 1, 2, 3}},
		{a: []int{5, 4, 3, 2, 1}, acopy: []int{5, 4, 3, 2, 1}, n: 6, want: []int{1, 5, 4, 3, 2}},

		// n multiple of the length of the slice
		{a: []int{10, 20, 30, 40, 50}, acopy: []int{10, 20, 30, 40, 50}, n: 5, want: []int{10, 20, 30, 40, 50}},
		{a: []int{10, 20, 30, 40, 50}, acopy: []int{10, 20, 30, 40, 50}, n: 10, want: []int{10, 20, 30, 40, 50}},
		// n multiple negative of the length of the slice
		{a: []int{10, 20, 30, 40, 50}, acopy: []int{10, 20, 30, 40, 50}, n: -5, want: []int{10, 20, 30, 40, 50}},
		{a: []int{10, 20, 30, 40, 50}, acopy: []int{10, 20, 30, 40, 50}, n: -10, want: []int{10, 20, 30, 40, 50}},
	}

	for _, test := range tests {
		rotate(test.a, test.n)
		if !slices.Equal(test.a, test.want) {
			t.Errorf("reverse(%v, %d) = %v, want %v", test.acopy, test.n, test.a, test.want)
		}
	}
}
