package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return fmt.Sprintf("Welcome to the Tech Palace, %s", strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	// In this context, simple string concatenation using the '+' operator proves to be more performant
	// and requires fewer memory allocations compared to using strings.Builder. Benchmarks demonstrate this:
	// goos: darwin
	// goarch: arm64
	// pkg: techpalace
	// BenchmarkAddBorderBuilder-8             20751432                58.12 ns/op           48 B/op          2 allocs/op
	// BenchmarkAddBorderConcatenate-8         21357711                56.39 ns/op           16 B/op          1 allocs/op
	//
	// While strings.Builder is typically more efficient for handling a large number of string concatenations,
	// it is less advantageous in scenarios involving a small, fixed number of string concatenations in a single expression.
	// This behavior is attributed to the Go compiler's optimization for string concatenations using the '+' operator,
	// particularly effective when the number of concatenated strings is known at compile time which is the case here
	// For more insights into string concatenation in Go and compiler optimizations, see:
	// - Go 101 String Optimization: https://go101.org/article/string.html
	// - Golang Programs String Concatenation Guide: https://www.golangprograms.com/how-to-concatenate-strings-in-golang.html
	//
	// See https://go101.org/article/string.htm

	border := strings.Repeat("*", numStarsPerLine)
	return border + "\n" + welcomeMsg + "\n" + border
}

func AddBorderBuilder(welcomeMsg string, numStarsPerLine int) string {

	var builder strings.Builder
	builder.Grow(numStarsPerLine*2 + len(welcomeMsg) + 2)

	border := strings.Repeat("*", numStarsPerLine)

	builder.WriteString(border)
	builder.WriteString("\n")
	builder.WriteString(welcomeMsg)
	builder.WriteString("\n")
	builder.WriteString(border)

	return builder.String()
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	noStars := strings.ReplaceAll(oldMsg, "*", "")
	return strings.TrimSpace(noStars)
}
