package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func comma1(s string) string {
	buffer := bytes.Buffer{}

	split := strings.Split(s, ".")
	integral := split[0]

	mantissa := ""

	if len(split) > 1 {
		mantissa = "." + split[1]
	}

	if len(integral) == 0 {
		return integral
	}

	i := 0
	if integral[0] == '+' || integral[0] == '-' {
		i = 1
	}

	skip := (len(integral) - i) % 3
	if skip == 0 {
		skip = 3
	}

	i += skip

	buffer.WriteString(integral[0:i])

	for ; i < len(integral); i += 3 {
		buffer.WriteString("," + integral[i:i+3])
	}

	buffer.WriteString(mantissa)

	return buffer.String()
}

func comma2(s string) string {
	if len(s) == 0 {
		return s
	}

	split := strings.Split(s, ".")
	integral, mantissa := split[0], ""
	if len(split) > 1 {
		mantissa = "." + split[1]
	}

	sign, integralWithoutSign := extractSign(integral)
	formattedIntegral := formatWithCommas(integralWithoutSign)

	var builder strings.Builder
	builder.WriteString(sign)
	builder.WriteString(formattedIntegral)
	builder.WriteString(mantissa)

	return builder.String()
}

func extractSign(s string) (string, string) {
	if len(s) > 0 && (s[0] == '+' || s[0] == '-') {
		return s[:1], s[1:]
	}
	return "", s
}

func formatWithCommas(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	start := n % 3
	if start == 0 {
		start = 3
	}

	var builder strings.Builder
	builder.WriteString(s[:start])

	for i := start; i < n; i += 3 {
		builder.WriteString(",")
		builder.WriteString(s[i : i+3])
	}

	return builder.String()
}

func comma3(s string) string {
	builder := strings.Builder{}

	split := strings.Split(s, ".")
	integral, mantissa := split[0], ""

	if len(split) > 1 {
		mantissa = "." + split[1]
	}

	if len(integral) == 0 {
		return integral
	}

	i := 0
	if integral[0] == '+' || integral[0] == '-' {
		i = 1
	}

	skip := (len(integral) - i) % 3
	if skip == 0 {
		skip = 3
	}

	i += skip

	builder.WriteString(integral[0:i])

	for ; i < len(integral); i += 3 {
		builder.WriteString("," + integral[i:i+3])
	}

	builder.WriteString(mantissa)

	return builder.String()
}

func main() {

	if len(os.Args) == 1 {
		s := "1234567890"
		println(s, ":", comma1(s))
	}

	for _, s := range os.Args[1:] {
		fmt.Printf("%s: %s\n", s, comma1(s))
	}

}
