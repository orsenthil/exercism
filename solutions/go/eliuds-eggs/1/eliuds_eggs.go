package eliudseggs

func EggCount(displayValue int) int {
	count := 0
	for displayValue > 0 {
		digit := displayValue % 2
		if digit == 1 {
			count++
		}
		displayValue /= 2
	}
	return count
}
