package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	var freq Frequency

	freq = make(Frequency)

	// Words can be separated by any form of punctuation (e.g. ":", "!", or "?") or whitespace (e.g. "\t", "\n", or " ").
	reg := regexp.MustCompile(`[\s+,:!?\t\n]`)
	output := reg.Split(phrase, -1)
	for i := 0; i < len(output); i++ {
		output[i] = strings.TrimSpace(output[i])
		output[i] = strings.ToLower(output[i])
		output[i] = regexp.MustCompile(`^'|'$`).ReplaceAllString(output[i], "")
		output[i] = regexp.MustCompile(`\.`).ReplaceAllString(output[i], "")
		output[i] = regexp.MustCompile(`[&@\$%\^]`).ReplaceAllString(output[i], "")

		if len(output[i]) == 0 {
			continue
		}
		freq[output[i]]++
	}
	return freq
}
