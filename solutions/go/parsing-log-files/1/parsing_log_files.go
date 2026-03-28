package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^(\[TRC]|\[DBG]|\[INF]|\[WRN]|\[ERR]|\[FTL])`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	// one or more in the character class
	re := regexp.MustCompile(`<[~=*-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*password.*"`)
	var count int = 0
	for _, line := range lines {
		result := re.FindStringSubmatch(line)
		if result != nil {
			count += len(result)
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	// array of strings
	output := make([]string, 0, len(lines))

	re := regexp.MustCompile(`User\s+([[:alpha:]]+\d*)`)
	var username string
	for _, line := range lines {
		sl := re.FindStringSubmatch(line)
		if len(sl) >= 2 {
			username = sl[1]
			output = append(output, fmt.Sprintf("[USR] %s %s", username, line))
		} else {
			output = append(output, line)
		}
	}
	return output
}
