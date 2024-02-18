package main

import (
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	var s = []byte("a  b  c   d    e     f")
	if len(os.Args) > 1 {
		s = []byte(os.Args[1])
	}

	fmt.Printf("Original: %v, %[1]q\n", s)
	RemoveAdjacentDuplicateSpaces(&s)
	fmt.Printf("Deduped:  %v, %[1]q\n", s)
}

// RemoveAdjacentDuplicateSpaces removes adjacent duplicates from a slice of strings
// in place.
func RemoveAdjacentDuplicateSpaces(s *[]byte) {
	if len(*s) <= 1 {
		return
	}
	c, offset := utf8.DecodeRune(*s)

	wasSpace := unicode.IsSpace(c)
	for i := offset; i < len(*s); {
		r, size := utf8.DecodeRune((*s)[i:])
		isSpace := unicode.IsSpace(r)
		if !(isSpace && wasSpace) {
			copy((*s)[offset:], (*s)[i:i+size])
			offset += size
		}
		wasSpace = isSpace
		i += size
	}
	*s = (*s)[:offset]
}
