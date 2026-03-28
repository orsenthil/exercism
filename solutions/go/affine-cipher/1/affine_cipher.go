package affinecipher

import (
	"errors"
	"strings"
)

func Encode(text string, a, b int) (string, error) {
	if gcd(a, 26) != 1 {
		return "", errors.New("a and 26 are not coprime")
	}
	var result string
	for _, char := range text {
		if char >= '0' && char <= '9' {
			result += string(char)
			continue
		}
		if char >= 'A' && char <= 'Z' {
			char = char + 32
		}
		if char < 'a' || char > 'z' {
			continue
		}
		encoded := (a*int(char-'a') + b) % 26
		result += string(rune('a' + encoded))
	}
	var output string

	for i := 0; i <= len(result); i += 5 {
		if i+5 > len(result) {
			output += result[i:]
		} else {
			output += result[i : i+5]
		}
		output += " "
	}

	return strings.TrimSpace(output), nil
}

func Decode(text string, a, b int) (string, error) {
	if gcd(a, 26) != 1 {
		return "", errors.New("a and 26 are not coprime")
	}
	var result string
	for _, char := range text {
		if char >= '0' && char <= '9' {
			result += string(char)
			continue
		}
		if char < 'a' || char > 'z' {
			continue
		}
		decoded := (modInverse(a, 26) * (int(char-'a') - b)) % 26
		if decoded < 0 {
			decoded += 26
		}
		result += string(rune('a' + decoded))
	}
	return strings.TrimSpace(result), nil
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func modInverse(a, m int) int {
	for i := 0; i < m; i++ {
		if (a*i)%m == 1 {
			return i
		}
	}
	return -1
}