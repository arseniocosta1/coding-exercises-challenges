package logs

import (
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	runeToApp := map[rune]string{
		'❗': "recommendation",
		'🔍': "search",
		'☀': "weather",
	}
	for _, v := range log {
		if _, ok := runeToApp[v]; ok {
			return runeToApp[v]
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.Replace(log, string(oldRune), string(newRune), -1)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
