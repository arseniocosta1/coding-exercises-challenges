package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var pattern = regexp.MustCompile(`\$(\w+)`)

// expand replaces each substring “$foo” within s by the text returned by f("foo").
func expand(s string, f func(string) string) string {
	return pattern.ReplaceAllStringFunc(s, func(s string) string {
		return f(s[1:])
	})
}

func main() {
	var input string

	if len(os.Args) > 1 {
		input = os.Args[1]
	} else {
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("Enter a string: ")
		if !scanner.Scan() {
			fmt.Fprintf(os.Stderr, "Error reading input %v\n", scanner.Err())
			os.Exit(1)
		}

		input = scanner.Text()
	}

	res := expand(input, strings.ToUpper)

	fmt.Printf("Input: %s\n", input)
	fmt.Printf("Output: %s\n", res)
}
