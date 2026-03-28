package darts

func Score(x, y float64) int {
	var outer = 10.0
	var middle = 5.0
	var inner = 1.0
	var score = 0
	var distance = x*x + y*y

	if distance <= inner*inner {
		score = 10
	} else if distance <= middle*middle {
		score = 5
	} else if distance <= outer*outer {
		score = 1
	}
	return score
}
