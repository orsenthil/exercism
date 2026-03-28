package encode

import (
	"strconv"
	"strings"
)

func writeEncode(output *strings.Builder, letter byte, count int){
	if count > 1 {
		output.WriteString(strconv.Itoa(count))
	}
	output.WriteByte(letter)
}

func writeDecode(output *strings.Builder, letter byte, count int){
	if count == 0 {
		output.WriteByte(letter) 
	 } else {
		for i := 0; i < count; i++ {
			output.WriteByte(letter)
		}
	 }
}

func isDigit(letter byte) bool {
	test := letter - '0'
	return test >= 0 && test <= 9
}

func appendDigit(num int, letter byte) int {
	return num * 10 + int(letter - '0')
}

// RuLengthEncode takes a string and returns a string with the run-length encoding of the input string.
func RunLengthEncode(input string) string {
	strlen := len(input)
	if strlen < 3 {
		return input
	}
	var output strings.Builder
	var count = 0
	for here, next := 0, 1; next < strlen; here, next = here+1, next+1 {
		count += 1
		if input[here] != input[next] {
			writeEncode(&output, input[here], count)
			count = 0
		}
	}
	count += 1
	writeEncode(&output, input[strlen-1], count)
	return output.String()
}

func RunLengthDecode(input string) string {
	strlen := len(input)
	if strlen < 3 {
		return input
	}
	var output strings.Builder
	var count = 0
	for here := 0; here < strlen; here++ {
		if isDigit((input[here])) {
			count = appendDigit(count, input[here])
		} else {
			writeDecode(&output, input[here], count)
			count = 0
		}
	}
	return output.String()
}
