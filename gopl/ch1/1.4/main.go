package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type LineInfo struct {
	Count     int
	Filenames []string
}

func main() {
	counts := make(map[string]LineInfo)
	for _, filename := range os.Args[1:] {
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			li := counts[line]
			li.Count++

			if li.Count > 1 && !slices.Contains(li.Filenames, filename) {
				li.Filenames = append(li.Filenames, filename)
			}
			counts[line] = li // update map
		}
	}
	for line, li := range counts {
		if li.Count > 1 {
			fmt.Printf("%-25s count: %-4d line: %s\n", strings.Join(li.Filenames, ", "), li.Count, line)
		}
	}
}
