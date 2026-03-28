package grains

import "errors"

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("not a valid chessboard number.")
	}
	board := make(map[int]uint64)
	board[1] = uint64(1)

	for i := 2; i <= number; i++ {
		board[i] = board[i-1] * uint64(2)
	}

	return board[number], nil
}

func Total() uint64 {
	var total uint64 = 0

	board := make(map[int]uint64)
	board[1] = uint64(1)

	for i := 2; i <= 64; i++ {
		board[i] = board[i-1] * uint64(2)
	}

	for i := 1; i <= 64; i++ {
		total += board[i]
	}

	return total
}
