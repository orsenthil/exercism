package lsproduct

import (
	"errors"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	var result int64
	result = 0
	if span > len(digits) {
		return 0, errors.New("span must be smaller than string length")
	}
	if span < 0 {
		return 0, errors.New("span must be greater than zero")
	}

	for i := 0; i < len(digits)-span+1; i++ {
		var product int64
		product = 1
		for j := 0; j < span; j++ {
			if (digits[i+j] - '0') > 9 {
				return 0, errors.New("digits must only contain digits")

			}
			product *= int64(digits[i+j] - '0')
		}
		if product > result {
			result = product
		}
	}
	return result, nil
}
