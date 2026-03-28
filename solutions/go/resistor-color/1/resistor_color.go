package resistorcolor

// Colors should return the list of all colors.
func Colors() []string {
	var colors []string
	colors = append(colors, "black")
	colors = append(colors, "brown")
	colors = append(colors, "red")
	colors = append(colors, "orange")
	colors = append(colors, "yellow")
	colors = append(colors, "green")
	colors = append(colors, "blue")
	colors = append(colors, "violet")
	colors = append(colors, "grey")
	colors = append(colors, "white")
	return colors
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {

	var colors []string
	colors = Colors()

	// Get the index of the color from colors array
	var index int
	for i := 0; i < len(colors); i++ {
		if colors[i] == color {
			index = i
		}
	}
	return index
}
