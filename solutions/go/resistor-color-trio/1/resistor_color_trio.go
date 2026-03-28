package resistorcolortrio

import (
	"fmt"
	"math"
)

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
	// map of color to value
	colorsMap := map[string]int{
		"black":  0,
		"brown":  1,
		"red":    2,
		"orange": 3,
		"yellow": 4,
		"green":  5,
		"blue":   6,
		"violet": 7,
		"grey":   8,
		"white":  9,
	}

	// store the values of the colors
	var colorval []int
	colorval = make([]int, len(colors))
	for i, color := range colors {
		colorval[i] = colorsMap[color]
	}

	resistance := (colorval[0]*10 + colorval[1]) * pow(10, colorval[2])

	if resistance < pow(10, 3) {
		return fmt.Sprintf("%d ohms", resistance)
	} else if resistance < pow(10, 6) {
		return fmt.Sprintf("%d kiloohms", resistance/pow(10, 3))
	} else if resistance < pow(10, 9) {
		return fmt.Sprintf("%d megaohms", resistance/pow(10, 6))
	} else if resistance < pow(10, 12) {
		return fmt.Sprintf("%d gigaohms", resistance/pow(10, 9))
	}

	return ""
}
