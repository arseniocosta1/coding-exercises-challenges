package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	var reader io.Reader = os.Stdin
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	freq := make(map[string]int)

	for scanner.Scan() {
		freq[scanner.Text()]++
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	for word, count := range freq {
		fmt.Printf("%-30s %d\n", word, count)
	}
}
