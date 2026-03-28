package wordy

import (
	"strconv"
	"strings"
)

func Answer(question string) (int, bool) {
	if question[len(question)-1] != '?' {
		return 0, false
	}

	question = question[:len(question)-1]

	words := strings.Split(question, " ")
	if len(words) < 3 {
		return 0, false
	}
	if words[0] != "What" || words[1] != "is" {
		return 0, false
	}
	words = words[2:]
	if len(words) < 1 {
		return 0, false
	}
	result, ok := parseNumber(words[0])
	if !ok {
		return 0, false
	}
	words = words[1:]
	for len(words) > 0 {
		if len(words) < 2 {
			return 0, false
		}
		operator := words[0]
		words = words[1:]

		switch operator {
		case "plus":
			number, ok := parseNumber(words[0])
			if !ok {
				return 0, false
			}
			words = words[1:]
			result += number
		case "minus":
			number, ok := parseNumber(words[0])
			if !ok {
				return 0, false
			}
			words = words[1:]
			result -= number
		case "multiplied":
			if len(words) < 1 || words[0] != "by" {
				return 0, false
			}
			words = words[1:]
			number, ok := parseNumber(words[0])
			if !ok {
				return 0, false
			}
			words = words[1:]
			result *= number
		case "divided":
			if len(words) < 1 || words[0] != "by" {
				return 0, false
			}
			words = words[1:]
			number, ok := parseNumber(words[0])
			if !ok {
				return 0, false
			}
			words = words[1:]
			result /= number
		default:
			return 0, false
		}
	}
	return result, true
}

func parseNumber(word string) (int, bool) {
	ans, err := strconv.Atoi(word)
	if err != nil {
		return 0, false
	}
	return ans, true
}
