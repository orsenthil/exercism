package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	var out map[string]int
	out = make(map[string]int)

	for k, v := range in {
		for _, letter := range v {
			letter := strings.ToLower(letter)
			out[letter] = k
		}
	}

	return out
}
