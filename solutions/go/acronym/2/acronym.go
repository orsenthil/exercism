// Package acronym provides a function to abbreviate a string.
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate abbreviates a string.
// It removes underscores, and splits by whitespaces, plus, and hyphens. It
// returns the first letter of each word in uppercase.
func Abbreviate(s string) string {
	var acronym string
	s = regexp.MustCompile(`[_]`).ReplaceAllString(s, "")

	words := regexp.MustCompile(`[\s+-]`).Split(s, -1)

	for _, word := range words {
		if len(word) == 0 {
			continue
		}
		acronym += strings.ToUpper(string(word[0]))
	}

	return acronym
}
