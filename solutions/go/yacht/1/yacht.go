package yacht

import "sort"

func countDice(dice []int) map[int]int {
	counts := make(map[int]int)
	for _, d := range dice {
		counts[d]++
	}
	return counts
}

func isFullHouse(dice []int) bool {
	counts := countDice(dice)
	threes := 0
	for _, d := range dice {
		if counts[d] == 3 {
			threes = d
		}
	}
	for _, d := range dice {
		if threes != 0 && counts[d] == 2 && d != threes {
			return true
		}
	}
	return false
}

func total(dice []int) int {
	total := 0
	for _, d := range dice {
		total += d
	}
	return total
}

func isFourOfaKind(dice []int) (bool, int) {
	counts := countDice(dice)
	for _, d := range dice {
		if counts[d] >= 4 {
			return true, d
		}
	}
	return false, 0
}

func isYatch(dice []int) bool {
	count := countDice(dice)
	for _, d := range dice {
		if count[d] == 5 {
			return true
		}
	}
	return false
}

func isBigStraight(dice []int) bool {
	checkDice := make([]int, len(dice))
	copy(checkDice, dice)
	sort.Ints(checkDice)
	return checkDice[0] == 2 && checkDice[1] == 3 && checkDice[2] == 4 && checkDice[3] == 5 && checkDice[4] == 6
}

func isLittleStraight(dice []int) bool {
	checkDice := make([]int, len(dice))
	copy(checkDice, dice)
	sort.Ints(checkDice)
	return checkDice[0] == 1 && checkDice[1] == 2 && checkDice[2] == 3 && checkDice[3] == 4 && checkDice[4] == 5
}

func Score(dice []int, category string) int {
	counts := countDice(dice)
	switch category {
	case "ones":
		return counts[1]
	case "twos":
		return 2 * counts[2]
	case "threes":
		return 3 * counts[3]
	case "fours":
		return 4 * counts[4]
	case "fives":
		return 5 * counts[5]
	case "sixes":
		return 6 * counts[6]
	case "full house":
		if isFullHouse(dice) {
			return total(dice)
		}
	case "four of a kind":
		match, value := isFourOfaKind(dice)
		if match {
			return 4 * value
		}
	case "little straight":
		if isLittleStraight(dice) {
			return 30
		}
	case "big straight":
		if isBigStraight(dice) {
			return 30
		}
	case "choice":
		return total(dice)
	case "yacht":
		if isYatch(dice) {
			return 50
		}
	}

	return 0
}
