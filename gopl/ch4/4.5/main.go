package main

import (
	"fmt"
)

func main() {
	s := []string{"a", "b"}

	fmt.Printf("Original: %v\n", s)
	removeAdjacentDupes(&s)
	fmt.Printf("Deduped:  %v\n", s)
}

// removeAdjacentDupes removes adjacent duplicates from a slice of strings
// in place.
func removeAdjacentDupes(s *[]string) {
	if len(*s) <= 1 {
		return
	}
	cur := 1
	for i := 1; i < len(*s); i++ {
		if (*s)[i] != (*s)[i-1] {
			(*s)[cur] = (*s)[i]
			cur++
		}
	}
	*s = (*s)[:cur]
}
