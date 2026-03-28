package allyourbase

import (
	"fmt"
	"math"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	var total int
	var outputDigits []int

	/** Input base must be >= 2 */
	if inputBase < 2 {
		return outputDigits, fmt.Errorf("input base must be >= 2")
	}

	/** Output base must be >= 2 */
	if outputBase < 2 {
		return outputDigits, fmt.Errorf("output base must be >= 2")
	}

	/** empty inputDigits */
	if len(inputDigits) == 0 {
		outputDigits = append(outputDigits, 0)
		return outputDigits, nil
	}

	for i := len(inputDigits) - 1; i >= 0; i-- {
		idx := len(inputDigits) - 1 - i
		/** all digits must satisfy 0 <= d < input base */
		if inputDigits[idx] < 0 || inputDigits[idx] >= inputBase {
			return outputDigits, fmt.Errorf("all digits must satisfy 0 <= d < input base")
		}
		total += inputDigits[idx] * int(math.Pow(float64(inputBase), float64(i)))
	}
	fmt.Println(inputDigits)
	fmt.Printf("inputBase %d, total: %d\n", inputBase, total)

	/* total is zero */
	if total == 0 {
		outputDigits = append(outputDigits, 0)
		return outputDigits, nil
	}

	for total > 0 {
		outputDigits = append([]int{total % outputBase}, outputDigits...)
		total /= outputBase
	}

	return outputDigits, nil
}
