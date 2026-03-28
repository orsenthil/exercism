// Package ocr provides functionality for handling OCR numbers.
package ocr
import (
	"fmt"
	"strings"
)
func recognizeDigit(rows []string) (string, error) {
	if len(rows) != 4 {
		return "", fmt.Errorf("input string has %d rows", len(rows))
	}
	for i := 0; i < 4; i++ {
		if len(rows[i]) != 3 {
			return "", fmt.Errorf("row %d has %d characters", i, len(rows[i]))
		}
	}
	switch {
	case rows[3] != "   ":
		return "?", nil
	case rows[0] == " _ " && rows[1] == "| |" && rows[2] == "|_|":
		return "0", nil
	case rows[0] == "   " && rows[1] == "  |" && rows[2] == "  |":
		return "1", nil
	case rows[0] == " _ " && rows[1] == " _|" && rows[2] == "|_ ":
		return "2", nil
	case rows[0] == " _ " && rows[1] == " _|" && rows[2] == " _|":
		return "3", nil
	case rows[0] == "   " && rows[1] == "|_|" && rows[2] == "  |":
		return "4", nil
	case rows[0] == " _ " && rows[1] == "|_ " && rows[2] == " _|":
		return "5", nil
	case rows[0] == " _ " && rows[1] == "|_ " && rows[2] == "|_|":
		return "6", nil
	case rows[0] == " _ " && rows[1] == "  |" && rows[2] == "  |":
		return "7", nil
	case rows[0] == " _ " && rows[1] == "|_|" && rows[2] == "|_|":
		return "8", nil
	case rows[0] == " _ " && rows[1] == "|_|" && rows[2] == " _|":
		return "9", nil
	default:
		return "?", nil
	}
}
// Recognize parses the input string as OCR numbers and return the string representation of them.
func Recognize(input string) []string {
	input = strings.TrimPrefix(input, "\n")
	rows := strings.Split(input, "\n")
	result := make([]string, 0)
	for i := 0; i < len(rows); i += 4 {
		var number string
		for j := 0; j < len(rows[i]); j += 3 {
			digitInput := make([]string, 4, 4)
			for k := 0; k < 4; k++ {
				digitInput[k] = rows[i+k][j : j+3]
			}
			digit, err := recognizeDigit(digitInput)
			if err != nil {
				panic(err)
			}
			number += digit
		}
		result = append(result, number)
	}
	return result
}
