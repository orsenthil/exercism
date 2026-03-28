// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "strings"

// Hey should have a comment documenting it.
func Hey(remark string) string {
	var allCaps = false
	var question = false

	remark = strings.TrimSpace(remark)

	if len(remark) == 0 {
		return "Fine. Be that way!"
	}

	for _, i := range remark {
		if i >= 'a' && i <= 'z' {
			allCaps = false
			break
		} else if i >= 'A' && i <= 'Z' {
			allCaps = true
		}
	}

	if strings.HasSuffix(remark, "?") {
		question = true
	}

	if question && allCaps {
		return "Calm down, I know what I'm doing!"
	} else if allCaps {
		return "Whoa, chill out!"
	} else if question {
		return "Sure."
	}

	return "Whatever."
}
