package dominoes

// Define the Domino type here.

func rotateDomino(d Domino) Domino {
	return Domino{d[1], d[0]}
}

func MakeChain(input []Domino) ([]Domino, bool) {
	panic("Please implement the MakeChain function")
}
