package allyourbase

import (
	"fmt"
	"math"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	var total int
	var outputDigits []int

	// Check if input base is valid
	if inputBase < 2 {
		return outputDigits, fmt.Errorf("input base must be >= 2")
	}

	// Check if output base is valid
	if outputBase < 2 {
		return outputDigits, fmt.Errorf("output base must be >= 2")
	}

	// If input digits are empty, return [0]
	if len(inputDigits) == 0 {
		outputDigits = append(outputDigits, 0)
		return outputDigits, nil
	}

	// Convert input digits to decimal
	for i := len(inputDigits) - 1; i >= 0; i-- {
		// Get the index of the current digit
		idx := len(inputDigits) - 1 - i
		// Check if each digit is valid
		if inputDigits[idx] < 0 || inputDigits[idx] >= inputBase {
			return outputDigits, fmt.Errorf("all digits must satisfy 0 <= d < input base")
		}
		total += inputDigits[idx] * int(math.Pow(float64(inputBase), float64(i)))
	}

	// If total is zero, return [0]
	if total == 0 {
		outputDigits = append(outputDigits, 0)
		return outputDigits, nil
	}

	// Convert decimal to output base
	for total > 0 {
		remainder := total % outputBase
		// Add the remainder to the beginning of the output digits
		// INSIGHT(senthil): This is a key insight for the golang way to append to the beginning of the slice
		outputDigits = append([]int{remainder}, outputDigits...)
		total /= outputBase
	}

	return outputDigits, nil
}
