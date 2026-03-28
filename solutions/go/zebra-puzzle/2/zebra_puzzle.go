// Package zebra provides functionality for solving the zebra puzzle.
package zebra
const (
	red, green, ivory, yellow, blue                         = "red", "green", "ivory", "yellow", "blue"
	englishman, spaniard, ukrainian, norwegian, japanese    = "Englishman", "Spaniard", "Ukrainian", "Norwegian", "Japanese"
	dog, snails, fox, horse, zebra                          = "dog", "snails", "fox", "horse", "zebra"
	coffee, tea, milk, orangeJuice, water                   = "coffee", "tea", "milk", "orange juice", "water"
	oldGold, kools, chesterfields, luckyStrike, parliaments = "Old Gold", "Kools", "Chesterfields", "Lucky Strike", "Parliaments"
)
var colors = []string{red, green, ivory, yellow, blue}
var residents = []string{englishman, spaniard, ukrainian, norwegian, japanese}
var pets = []string{dog, snails, fox, horse, zebra}
var beverages = []string{coffee, tea, milk, orangeJuice, water}
var cigarettes = []string{oldGold, kools, chesterfields, luckyStrike, parliaments}
// Solution represents the solution to the zebra puzzle.
type Solution struct {
	DrinksWater string
	OwnsZebra   string
}
// SolvePuzzle solves the zebra puzzle.
func SolvePuzzle() Solution {
	candidates := getPermutations([]int{0, 1, 2, 3, 4})
	for _, res := range candidates {
		if residents[res[0]] != norwegian {
			continue // Condition #10
		}
	bevLoop:
		for _, bev := range candidates {
			if beverages[bev[2]] != milk {
				continue // Condition #9
			}
			for idx, elem := range bev {
				if beverages[elem] == tea && residents[res[idx]] != ukrainian {
					continue bevLoop // Condition #5
				}
			}
		colLoop:
			for _, col := range candidates {
				for idx, elem := range col {
					if colors[elem] == red && residents[res[idx]] != englishman {
						continue colLoop // Condition #2
					} else if colors[elem] == green && beverages[bev[idx]] != coffee {
						continue colLoop // Condition #4
					} else if colors[elem] == green && (idx == 0 || colors[col[idx-1]] != ivory) {
						continue colLoop // Condition #6
					} else if colors[elem] == blue {
						if (idx == 0 && residents[res[idx+1]] != norwegian) ||
							((0 < idx && idx < 4) && residents[res[idx-1]] != norwegian && residents[res[idx+1]] != norwegian) ||
							idx == 4 && residents[res[idx-1]] != norwegian {
							continue colLoop // Condition #15
						}
					}
				}
			petLoop:
				for _, pet := range candidates {
					for idx, elem := range pet {
						if pets[elem] == dog && residents[res[idx]] != spaniard {
							continue petLoop // Condition #3
						}
					}
				cigLoop:
					for _, cig := range candidates {
						for idx, elem := range cig {
							if cigarettes[elem] == oldGold && pets[pet[idx]] != snails {
								continue cigLoop // Condition #7
							} else if cigarettes[elem] == kools && colors[col[idx]] != yellow {
								continue cigLoop // Condition #8
							} else if cigarettes[elem] == chesterfields {
								if (idx == 0 && pets[pet[idx+1]] != fox) ||
									((0 < idx && idx < 4) && pets[pet[idx-1]] != fox && pets[pet[idx+1]] != fox) ||
									idx == 4 && pets[pet[idx-1]] != fox {
									continue cigLoop // Condition #11
								}
							} else if cigarettes[elem] == kools {
								if (idx == 0 && pets[pet[idx+1]] != horse) ||
									((0 < idx && idx < 4) && pets[pet[idx-1]] != horse && pets[pet[idx+1]] != horse) ||
									idx == 4 && pets[pet[idx-1]] != horse {
									continue cigLoop // Condition #12
								}
							} else if cigarettes[elem] == luckyStrike && beverages[bev[idx]] != orangeJuice {
								continue cigLoop // Condition #13
							} else if cigarettes[elem] == parliaments && residents[res[idx]] != japanese {
								continue cigLoop // Condition #14
							}
						}
						// if we are here, it means all conditions are met.
						water, zebra := getResidentDrinkingWater(res, bev), getResidentOwningZebra(res, pet)
						return Solution{DrinksWater: residents[water], OwnsZebra: residents[zebra]}
					}
				}
			}
		}
	}
	panic("could not solve puzzle")
}
func getPermutations(input []int) [][]int {
	if len(input) == 1 {
		return [][]int{[]int{input[0]}}
	}
	result := make([][]int, 0)
	for i := 0; i < len(input); i++ {
		tmpInput := make([]int, len(input)) // Need to prevent the original underlying array gets modified
		copy(tmpInput, input)
		permu := getPermutations(append(tmpInput[0:i], tmpInput[i+1:len(input)]...))
		for j := 0; j < len(permu); j++ {
			tmpResult := []int{input[i]}
			tmpResult = append(tmpResult, permu[j]...)
			result = append(result, tmpResult)
		}
	}
	return result
}
func getResidentDrinkingWater(res, bev []int) int {
	for i := 0; i < len(bev); i++ {
		if beverages[bev[i]] == water {
			return res[i]
		}
	}
	panic("nobody drinks water")
}
func getResidentOwningZebra(res, pet []int) int {
	for i := 0; i < len(pet); i++ {
		if pets[pet[i]] == zebra {
			return res[i]
		}
	}
	panic("nobody owns zebra")
}
