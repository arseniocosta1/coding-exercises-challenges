package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
)

func anagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	a1 := []rune(a)
	b1 := []rune(b)

	sort.Slice(a1, func(i, j int) bool { return a1[i] < a1[j] })
	sort.Slice(b1, func(i, j int) bool { return b1[i] < b1[j] })

	return string(a1) == string(b1)
}

// slices.Sort is available since Go 1.21 (https://go.dev/doc/go1.21#slices)
func anagramSlicesSort(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	a1 := []rune(a)
	b1 := []rune(b)

	slices.Sort(a1)
	slices.Sort(b1)

	return string(a1) == string(b1)
}

func anagramMap(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	m := make(map[rune]int)

	for i := 0; i < len(a); i++ {
		m[rune(a[i])]++
		m[rune(b[i])]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}

func main() {
	if len(os.Args) < 2 {
		a := "1234567890"
		b := "0987654321"
		fmt.Printf("a: %s, b: %s: anagram: %v\n", a, b, anagram(a, b))
	}
}
