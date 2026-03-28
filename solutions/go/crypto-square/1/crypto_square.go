package cryptosquare

import (
	"fmt"
	"math"
	"strings"
)

func Encode(pt string) string {
	lowerString := strings.ToLower(pt)
	var strippedString string = ""
	for _, c := range lowerString {
		if c >= 'a' && c <= 'z' || c >= '0' && c <= '9' {
			strippedString += string(c)
		}
	}
	// The plaintext should be organized into a rectangle as square as possible. The size of the rectangle should be decided by the length of the message.

	// If the message is a length that creates a perfect rectangle (i.e. the number of characters in the message is a square number), use that number of columns.
	// If the message doesn't fit neatly into a rectangle, choose the number of columns that corresponds to the smallest square that is larger than the number of characters in the message.

	length := len(strippedString)
	var columns int = 1
	var rows int = 1
	for columns*columns < length {
		columns++
	}
	if columns*columns == length {
		rows = columns
	} else {
		rows = int(math.Ceil(float64(length) / float64(columns)))
	}

	// assert rows - cols <= 1
	if rows-columns > 1 {
		fmt.Println("Something went wrong")
	}
	// slice of strings
	var result []string = make([]string, rows)
	var start, end int = 0, 0

	for i := 0; i < rows; i++ {
		start = i * columns
		end = start + columns
		if end > length {
			end = length
		}
		result[i] = strippedString[start:end]
	}
	var output string = ""
	for c := 0; c < columns; c++ {
		for r := 0; r < rows; r++ {
			if c < len(result[r]) {
				output += string(result[r][c])
			} else {
				output += " "
			}
		}
		output += " "
	}

	return output[:len(output)-1]
}
