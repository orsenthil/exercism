package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	// slice of string
	var result []string
	// loop through candidates
	for _, candidate := range candidates {
		// check if subject is an anagram of candidate
		if isAnagram(subject, candidate) {
			// if so, append to result
			result = append(result, candidate)
		}
	}
	return result
}

// sortRunes is a type that implements the sort.Interface interface
type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func isAnagram(subject string, candidate string) bool {
	// convert subject to lowercase
	subject = strings.ToLower(subject)
	// convert candidate to lowercase
	candidate = strings.ToLower(candidate)

	if subject == candidate {
		return false
	}
	if len(subject) != len(candidate) {
		return false
	}
	// sort subject
	subject = SortString(subject)
	// sort candidate
	candidate = SortString(candidate)

	return subject == candidate
}
