package popcount_test

import (
	"testing"

	"gopl.io/ch2/popcount"

	main "gopl-solutions/ch2/2.3"
)

func BitCount(x uint64) int {
	// Hacker's Delight, Figure 5-2.
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	return int(x & 0x7f)
}

func PopCountByClearing(x uint64) int {
	n := 0
	for x != 0 {
		x = x & (x - 1) // clear rightmost non-zero bit
		n++
	}
	return n
}

func PopCountByShifting(x uint64) int {
	n := 0
	for i := uint(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			n++
		}
	}
	return n
}

// -- Benchmarks --

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkBitCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BitCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByClearing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByClearing(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByShifting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByShifting(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main.PopCountLoop(0x1234567890ABCDEF)
	}
}

// $ go test -cpu=4 -bench=.
// goos: darwin
// goarch: arm64
// pkg: gopl.io/solutions/ch2/2.3
// BenchmarkPopCount-4                     1000000000               0.3130 ns/op
// BenchmarkBitCount-4                     1000000000               0.3120 ns/op
// BenchmarkPopCountByClearing-4           100000000               11.30 ns/op
// BenchmarkPopCountByShifting-4           55872855                21.73 ns/op
// BenchmarkPopCountLoop-4                 309458900                3.767 ns/op
// PASS
// ok      gopl.io/solutions/ch2/2.3       4.634s
