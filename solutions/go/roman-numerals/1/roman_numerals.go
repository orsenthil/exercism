package romannumerals

import "errors"

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input > 3999 {
		return "", errors.New("invalid input")
	}

	// Define Roman Numeral symbols
	romanSymbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	// Define Roman Numeral values
	romanValues := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	// Convert integer to Roman Numeral
	var romanNumeral string
	for i := 0; i < len(romanValues); i++ {
		for input >= romanValues[i] {
			romanNumeral += romanSymbols[i]
			input -= romanValues[i]
		}
	}

	return romanNumeral, nil

}
