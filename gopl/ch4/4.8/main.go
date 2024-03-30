package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func charcount(reader io.Reader) error {
	counts := make(map[rune]int)    // counts of Unicode characters
	digits := make(map[rune]int)    // counts of Unicode digits
	letters := make(map[rune]int)   // counts of Unicode letters
	cats := make(map[string]int)    // counts of Unicode categories
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(reader)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		if unicode.IsLetter(r) {
			letters[r]++
		}

		if unicode.IsDigit(r) {
			digits[r]++
		}

		for c, rt := range unicode.Properties {
			if unicode.In(r, rt) {
				cats[c]++
			}
		}

		if unicode.In(r, unicode.N) {
			cats["Number"]++
		}

		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}

	const format = "%-30q\t%d\n"
	fmt.Printf("digits count\t\n")
	for c, n := range digits {
		fmt.Printf(format, c, n)
	}

	fmt.Printf("letters count\t\n")
	for c, n := range letters {
		fmt.Printf(format, c, n)
	}

	fmt.Printf("categories count\t\n")
	for c, n := range cats {
		fmt.Printf(format, c, n)

	}
	return nil
}

func main() {
	args := os.Args[1:]

	var reader io.Reader = os.Stdin

	if len(args) > 0 {
		fmt.Println(args[0])
		reader = strings.NewReader(args[0])
	}
	if err := charcount(reader); err != nil {
		fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
		os.Exit(1)
	}
}
