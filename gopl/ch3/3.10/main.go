package main

import (
	"bytes"
	"os"
)

func comma(s string) string {
	buffer := bytes.Buffer{}

	i := len(s) % 3
	if i == 0 {
		i = 3
	}

	buffer.WriteString(s[:i])

	for ; i < len(s); i += 3 {
		buffer.WriteString("," + s[i:i+3])
	}

	return buffer.String()
}

func main() {
	s := "12345"

	if len(os.Args) > 1 {
		s = os.Args[1]
	}

	println(s, ":", comma(s))
}
