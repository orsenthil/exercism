package secret

import (
	"fmt"
	"strconv"
)

func Handshake(code uint) []string {
	binary := strconv.FormatInt(int64(code), 2)
	paddedBinary := fmt.Sprintf("%05s", binary)
	reversed := reverse(paddedBinary)
	handshake := []string{}
	if reversed[0] == '1' {
		handshake = append(handshake, "wink")
	}
	if reversed[1] == '1' {
		handshake = append(handshake, "double blink")
	}
	if reversed[2] == '1' {
		handshake = append(handshake, "close your eyes")
	}
	if reversed[3] == '1' {
		handshake = append(handshake, "jump")
	}
	if reversed[4] == '1' {

		handshake = reverseSlice(handshake)
	}

	return handshake
}

func reverseSlice(handshake []string) []string {
	for i, j := 0, len(handshake)-1; i < j; i, j = i+1, j-1 {
		handshake[i], handshake[j] = handshake[j], handshake[i]
	}
	return handshake
}

func reverse(binary string) string {
	runes := []rune(binary)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
