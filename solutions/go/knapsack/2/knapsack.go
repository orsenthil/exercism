package knapsack

type Item struct {
	Weight, Value int
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Knapsack takes in a maximum carrying capacity and a collection of items
// and returns the maximum value that can be carried by the knapsack
// given that the knapsack can only carry a maximum weight given by maximumWeight
func Knapsack(maximumWeight int, items []Item) int {
	amountOfItems := len(items)
	knapsack := make([][]int, amountOfItems+1)

	for i := range knapsack {
		knapsack[i] = make([]int, maximumWeight+1)
	}

	for i := 1; i <= amountOfItems; i++ {
		for w := 0; w <= maximumWeight; w++ {
			if items[i-1].Weight > w {
				knapsack[i][w] = knapsack[i-1][w]
			} else {
				withoutCurrent := knapsack[i-1][w]
				withCurrent := knapsack[i-1][w-items[i-1].Weight] + items[i-1].Value
				knapsack[i][w] = maxValue(withoutCurrent, withCurrent)
			}

		}
	}
	return knapsack[amountOfItems][maximumWeight]
}
