package isbn

import (
	"fmt"
	"strings"
)

func IsValidISBN(isbn string) bool {

	isbn = strings.ReplaceAll(isbn, "-", "")

	if len(isbn) != 10 {
		return false
	}

	// Determine if the string is a valid ISBN number

	multiple := 10
	var result int = 0

	for _, v := range isbn {
		if multiple == 1 && v == 'X' {
			result += 10 * multiple
			continue
		}
		if v >= '0' && v <= '9' {
			result += (int(v-'0') * multiple)
			multiple--
		} else {
			return false
		}
	}

	fmt.Println(result)

	if result == 0 {
		return false
	}

	if result%11 == 0 {
		return true
	}

	return false
}
