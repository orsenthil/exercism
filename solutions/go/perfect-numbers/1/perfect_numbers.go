package perfect

import "errors"

// Define the Classification type here.

type Classification int

const (
	ClassificationDeficient Classification = iota
	ClassificationPerfect
	ClassificationAbundant
)

var ErrOnlyPositive = errors.New("only positive numbers are allowed")

func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return 0, ErrOnlyPositive
	}

	aliquotSum := getSum(getFactors(n))

	switch {
	case aliquotSum == n:
		return ClassificationPerfect, nil
	case aliquotSum < n:
		return ClassificationDeficient, nil
	default:
		return ClassificationAbundant, nil
	}
}

func getFactors(n int64) []int64 {
	var factors []int64
	for i := int64(1); i < n; i++ {
		if n%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}

func getSum(factors []int64) int64 {
	var sum int64
	for _, factor := range factors {
		sum += factor
	}
	return sum
}
