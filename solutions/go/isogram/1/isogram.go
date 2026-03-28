package isogram

import "strings"

func IsIsogram(word string) bool {
	var s string

	for _, c := range word {
		s = strings.ToUpper(string(c))
		if s != " " && s != "-" {
			if strings.Count(strings.ToUpper(word), s) > 1 {
				return false
			}
		}
	}
	return true
}
