package queenattack

import "errors"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func CanQueenAttack(whitePosition, blackPosition string) (bool, error) {
	if whitePosition == blackPosition {
		return false, errors.New("same square")
	}

	if whitePosition == "" || blackPosition == "" {
		return false, errors.New("empty position")
	}

	whiteCoords := make([]int, 2)
	blackCoords := make([]int, 2)

	whiteCoords[0] = int(whitePosition[0] - 'a')
	whiteCoords[1] = int(whitePosition[1] - '1')

	blackCoords[0] = int(blackPosition[0] - 'a')
	blackCoords[1] = int(blackPosition[1] - '1')

	if whiteCoords[0] < 0 || whiteCoords[0] > 7 || whiteCoords[1] < 0 || whiteCoords[1] > 7 {
		return false, errors.New("invalid position")
	}
	if blackCoords[0] < 0 || blackCoords[0] > 7 || blackCoords[1] < 0 || blackCoords[1] > 7 {
		return false, errors.New("invalid position")
	}

	dx := abs(whiteCoords[0] - blackCoords[0])
	dy := abs(whiteCoords[1] - blackCoords[1])

	if dx == 0 && dy == 0 {
		return false, nil
	}

	if dx == dy || dx == 0 || dy == 0 {
		return true, nil
	}

	return false, nil
}
