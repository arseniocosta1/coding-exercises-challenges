package main_test

import (
	"testing"

	"gopl-solutions/ch1/1.3"
)

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main.Echo1()
	}
}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main.Echo2()
	}
}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main.Echo3()
	}
}
