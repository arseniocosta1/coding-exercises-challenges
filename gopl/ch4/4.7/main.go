package main

import (
	"fmt"
)

func main() {
	//
	//s := []byte("\U0001F600")
	//s := []byte("\U0001F600\u3042\u00C5Räksmörgås")
	//s := []byte("2\U0001F6031")
	//s := []byte("👋🏿")
	s := []byte("👋\U0001F3FF Hello there!")

	fmt.Printf("Original: %v, %[1]q\n", s)
	reverse(s)
	fmt.Printf("Reverse:  %v, %q\n", s, string(s))
}

const (
	t1 = 0b00000000
	t2 = 0b11000000
	t3 = 0b11100000
	t4 = 0b11110000
)

// reverse reverses a slice of UTF-8 encoded bytes in place.
// It assumes that the input is a valid UTF-8 byte sequence;
// otherwise, the behavior is undefined, and the function may panic.
// For example, providing a byte slice like []byte{0b11000000},
// where the leading bits indicate a multi-byte UTF-8 character without the following bytes,
// will result in incorrect behavior or a panic.
// reverse does not check for skin tone modifiers.
// As shown by running the benchmark, reverse does not allocate any memory.
func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	for i := 0; i < len(s); i++ {
		if s[i]&t4 == t4 {
			s[i], s[i-1], s[i-2], s[i-3] = s[i-3], s[i-2], s[i-1], s[i]
			i += 3
			continue
		}

		if s[i]&t3 == t3 {
			s[i], s[i-2] = s[i-2], s[i]
			i += 2
			continue
		}

		if s[i]&t2 == t2 {
			s[i], s[i-1] = s[i-1], s[i]
			i++
			continue
		}
	}
}
