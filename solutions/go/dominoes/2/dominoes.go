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

func getChain(input []Domino, usedDominos []int, last int, next int) ([]int, int, bool) {
	if len(usedDominos) == len(input) {
		return usedDominos, last, true
	}

	for i := 0; i < len(input); i++ {
		if !contains(usedDominos, i) {
			if input[i][0] == next {
				usedDominos = append(usedDominos, i)
				usedDominos, last, ok := getChain(input, usedDominos, next, input[i][1])
				if ok {
					return usedDominos, last, true
				}
				usedDominos = usedDominos[:len(usedDominos)-1]
			} else if input[i][1] == next {
				usedDominos = append(usedDominos, i)
				usedDominos, last, ok := getChain(input, usedDominos, next, input[i][0])
				if ok {
					return usedDominos, last, true
				}
				usedDominos = usedDominos[:len(usedDominos)-1]
			}
		}
	}
	return usedDominos, last, false
}

func contains(dominos []int, i int) bool {
	for _, d := range dominos {
		if d == i {
			return true
		}
	}
	return false
}
