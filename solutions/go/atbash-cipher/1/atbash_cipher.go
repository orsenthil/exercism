package atbash

import "strings"

func Atbash(s string) string {
	var result []rune

	s = strings.ToLower(s)

	var count int = 0

	for _, r := range s {
		if r == ' ' || r == ',' || r == '.' {
			continue
		}
		if r >= 'a' && r <= 'z' {
			result = append(result, 'z'-(r-'a'))
		} else {
			result = append(result, r)
		}
		count++
		if count%5 == 0 {
			result = append(result, ' ')
		}
	}
	return strings.TrimSpace(string(result))
}
