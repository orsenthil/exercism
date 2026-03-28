package rectangles

// The coordinate system is as follows:
// (0,0) +-----> x
//       |
//       |
//       |
//       v
//       y

func findCorners(diagram []string) [][2]int {
	var corners [][2]int
	for y, line := range diagram {
		for x, char := range line {
			if char == '+' {
				corners = append(corners, [2]int{x, y})
			}
		}
	}
	return corners
}

func isHorizontalLineValid(diagram []string, y, x1, x2 int) bool {
	for x := x1 + 1; x < x2; x++ {
		if diagram[y][x] != '-' && diagram[y][x] != '+' {
			return false
		}
	}
	return true
}

func isVerticalLineValid(diagram []string, x, y1, y2 int) bool {
	for y := y1 + 1; y < y2; y++ {
		if diagram[y][x] != '|' && diagram[y][x] != '+' {
			return false
		}
	}
	return true
}

func isRectangle(diagram []string, x1, y1, x2, y2 int) bool {

	if !isHorizontalLineValid(diagram, y1, x1, x2) {
		return false
	}
	if !isHorizontalLineValid(diagram, y2, x1, x2) {
		return false
	}
	if !isVerticalLineValid(diagram, x1, y1, y2) {
		return false
	}
	if !isVerticalLineValid(diagram, x2, y1, y2) {
		return false
	}

	return true
}

func Count(diagram []string) int {
	corners := findCorners(diagram) // x, y
	count := 0

	// For each pair of corners, treating this as top-left
	for _, topLeft := range corners {
		for _, topRight := range corners {
			if topRight[1] != topLeft[1] || topRight[0] <= topLeft[0] {
				continue
			}
			for _, bottomRight := range corners {
				if bottomRight[0] != topRight[0] || bottomRight[1] <= topRight[1] {
					continue
				}
				for _, bottomLeft := range corners {
					if bottomLeft[1] != bottomRight[1] || bottomLeft[0] != topLeft[0] {
						continue
					}
					if isRectangle(diagram, topLeft[0], topLeft[1], bottomRight[0], bottomRight[1]) {
						count++
					}
				}

			}
		}

	}
	return count
}
