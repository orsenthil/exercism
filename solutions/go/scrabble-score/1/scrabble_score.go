package scrabble

import "strings"

func Score(word string) int {

	var scrabbleScore = make(map[string]int)
	scrabbleScore["A"] = 1
	scrabbleScore["E"] = 1
	scrabbleScore["I"] = 1
	scrabbleScore["O"] = 1
	scrabbleScore["U"] = 1
	scrabbleScore["L"] = 1
	scrabbleScore["N"] = 1
	scrabbleScore["R"] = 1
	scrabbleScore["S"] = 1
	scrabbleScore["T"] = 1

	scrabbleScore["D"] = 2
	scrabbleScore["G"] = 2

	scrabbleScore["B"] = 3
	scrabbleScore["C"] = 3
	scrabbleScore["M"] = 3
	scrabbleScore["P"] = 3

	scrabbleScore["F"] = 4
	scrabbleScore["H"] = 4
	scrabbleScore["V"] = 4
	scrabbleScore["W"] = 4
	scrabbleScore["Y"] = 4

	scrabbleScore["K"] = 5

	scrabbleScore["J"] = 8
	scrabbleScore["X"] = 8

	scrabbleScore["Q"] = 10
	scrabbleScore["Z"] = 10

	var count = 0

	for _, c := range word {
		count += scrabbleScore[strings.ToUpper(string(c))]
	}

	return count

}
