package dominoes

// Define the Domino type here.

type Domino [2]int

func rotateDomino(d Domino) Domino {
	return Domino{d[1], d[0]}
}

func MakeChain(input []Domino) (chain []Domino, ok bool) {
	chain = make([]Domino, 0, len(input))

	if len(input) == 0 {
		return chain, true
	} else if (len(input) == 1) && (input[0][0] == input[0][1]) {
		chain = append(chain, input[0])
		return input, true
	}
	usedDominos, _, ok := getChain(input, []int{0}, input[0][1], input[0][0])

	if ok {
		for i := 0; i < len(usedDominos); i++ {
			if len(chain) > 0 && chain[len(chain)-1][1] != input[usedDominos[i]][0] {
				chain = append(chain, Domino{input[usedDominos[i]][1], input[usedDominos[i]][0]})
			} else {
				chain = append(chain, input[usedDominos[i]])
			}

		}
	}
	return chain, ok
}

func getChain(input []Domino, usedDominos []int, numberToChain int, beginingNumber int) ([]int, int, bool) {
	for i := 1; i < len(input); i++ {
		if isDominoUsed(usedDominos, i) {
			continue
		}
		if nextNumberToCheck, ok := isCandidate(input[i], numberToChain); ok {
			if len(usedDominos) == len(input)-1 {
				if nextNumberToCheck == beginingNumber {
					return append(usedDominos, i), nextNumberToCheck, true
				}
				return nil, -1, false
			}

			chain, nextNumberToChain, ok := getChain(input, append(usedDominos, i), nextNumberToCheck, beginingNumber)

			if len(chain) > 0 {
				return chain, nextNumberToChain, ok
			}
		}
	}

	return nil, -1, false
}

func isCandidate(domino Domino, numberToChain int) (int, bool) {
	if domino[0] == numberToChain {
		return domino[1], true
	} else if domino[1] == numberToChain {
		return domino[0], true
	}
	return -1, false
}

func isDominoUsed(dominoToCheck []int, usedDominoes int) bool {
	for i := 0; i < len(dominoToCheck); i++ {
		if dominoToCheck[i] == usedDominoes {
			return true
		}
	}
	return false
}
