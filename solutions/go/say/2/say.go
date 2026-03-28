package say

import (
	"strconv"
)

var ones = [10]string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var tens = [10]string{"", "ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
var teens = [10]string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
var groups = [4]string{"", " thousand", " million", " billion"}

var addSpaceIf = map[bool]string{true: " ", false: ""}
var addDashIf = map[bool]string{true: "-", false: ""}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func valForIndex(str string, index int) (value byte, ok bool) {
	if index > -1 {
		return str[index], true
	}
	return value, false
}

func buildGroup(group string, groupNum int) string {
	phrase := ""
	for idx, place := len(group)-1, 1; idx > -1; idx, place = idx-1, place*10 {
		if group[idx] == '0' {
			continue
		}
		digit := group[idx] - '0'
		switch place {
		case 1:
			if tens, hasTens := valForIndex(group, idx-1); hasTens && tens == '1' {
				phrase = teens[digit]
			} else {
				phrase = ones[digit]
			}
		case 10:
			if digit != 1 || phrase == "" {
				phrase = tens[digit] + addDashIf[phrase != ""] + phrase
			}
		default:
			phrase = ones[digit] + " hundred" + addSpaceIf[phrase != ""] + phrase
		}
	}
	if phrase != "" {
		phrase = phrase + groups[groupNum]
	}
	return phrase
}

func Say(n int64) (string, bool) {
	if n < 0 || n > 999999999999 {
		return "", false
	}
	if n == 0 {
		return "zero", true
	}
	digits := strconv.Itoa(int(n))
	stop := len(digits)
	start := max(0, stop-3)
	phrase := ""
	for groupNum := 0; stop > 0; start, stop, groupNum = max(0, start-3), stop-3, groupNum+1 {
		group := buildGroup(digits[start:stop], groupNum)
		phrase = group + addSpaceIf[phrase != "" && group != ""] + phrase
	}
	return phrase, true
}
