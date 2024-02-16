package main

import (
	"crypto/sha256"
	"testing"
)

// -- Benchmarks --

func BenchmarkCountDiffBits(b *testing.B) {
	h1 := sha256.Sum256([]byte("x"))
	h2 := sha256.Sum256([]byte("X"))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		countDiffBits(&h1, &h2)
	}
}

func BenchmarkCountDiffBitsWithPrecomputedTable(b *testing.B) {
	h1 := sha256.Sum256([]byte("x"))
	h2 := sha256.Sum256([]byte("X"))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		countDiffBitsWithPrecomputedTable(&h1, &h2)
	}
}

//❯ go test -v -bench=. --benchmem
//goos: darwin
//goarch: arm64
//pkg: gopl-solutions/ch4/4.1
//BenchmarkCountDiffBits
//BenchmarkCountDiffBits-8                        14192522                81.98 ns/op            0 B/op          0 allocs/op
//BenchmarkCountDiffBitsWithPrecomputedTable
//BenchmarkCountDiffBitsWithPrecomputedTable-8    100000000               11.44 ns/op            0 B/op          0 allocs/op
//PASS
//ok      gopl-solutions/ch4/4.1  3.341s
