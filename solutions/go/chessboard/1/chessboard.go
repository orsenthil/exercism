package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard contains a map of eight Ranks, accessed with values from "A" to "H"
// var r1, r2, r3, r4, r5, r6, r7, r8 Rank

type Chessboard map[string]Rank

// {"A": r1, "B": r2, "C": r3, "D": r4, "E": r5, "F": r6, "G": r7, "H": r8}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	var count = 0
    if val, ok := cb[rank]; ok {
        for _, b := range val {
            if b == true {
                count += 1
            }   
        }
	}
    return count
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
    var count = 0
    for _, r := range cb {
        for i, value := range r {
            if (i + 1) == file && value == true {
                count += 1
            }
        }
    }
	return count
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
    var count = 0
    for _, rank := range cb {
        for range rank {
            count += 1
        }
    }
	return count
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
    var count = 0
    for _, rank := range cb {
        for _, cell := range rank {
            if cell == true {
                count += 1
            }
        }
    }
	return count
}