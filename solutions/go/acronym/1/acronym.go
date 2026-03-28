// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	// abbreviate the string
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
