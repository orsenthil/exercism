package change



import (
	"errors"
	"sort"
)

func Change(coins []int, target int) ([]int, error) {
	if target == 0 {
		return []int{}, nil
	}
	solutions := make([][]int, 0)
	for i := len(coins) - 1; i >= 0; i-- {
		if coins[i] <= target {
			solution := solve(coins[:i+1], target)

			if solution != nil {
				solutions = append(solutions, solution)
			}

		}
	}

	if len(solutions) > 0 {
		sort.Slice(solutions, func(i, j int) bool { return len(solutions[i]) < len(solutions[j]) })
		return solutions[0], nil
	}
	return nil, errors.New("no possible solution")
}

func solve(coins []int, target int) []int {
	numOfCoins := target / coins[len(coins)-1]
	remainder := target % coins[len(coins)-1]
	coinsForRemainder := make([]int, 0)

	for remainder != 0 {
		if len(coins) == 1 {
			return nil
		}

		coinsForRemainder = solve(coins[:len(coins)-1], remainder)

		if len(coinsForRemainder) == 0 {
			if numOfCoins <= 1 {
				return nil
			}
			numOfCoins -= 1
			remainder += coins[len(coins)-1]
		} else {
			remainder = 0
		}
	}

	coinsForQuotient := make([]int, numOfCoins, numOfCoins)
	for i := 0; i < numOfCoins; i++ {
		coinsForQuotient[i] = coins[len(coins)-1]
	}

	return append(coinsForRemainder, coinsForQuotient...)
}
