package parsinglogfiles

import (
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(?:TRC|DBG|INF|WRN|ERR|FTL)]`)

	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)

	res := re.Split(text, -1)

	return res
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`".*(?i)password.*"`)

	count := 0
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)

	res := re.ReplaceAllString(text, "")

	return res
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(?P<username>\w+)\s`)

	res := make([]string, len(lines))

	for i, line := range lines {
		if match := re.FindStringSubmatch(line); len(match) > 1 {
			res[i] = "[USR] " + match[1] + " " + line
		} else {
			res[i] = line
		}
	}

	return res
}
