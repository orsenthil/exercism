package prism

import "math"

type Position struct {
	x     float64
	y     float64
	angle float64
}

type Prism struct {
	id    int
	x     float64
	y     float64
	angle float64
}

func FindSequence(start Position, prisms []Prism) []int {
	current := start
	result := []int{}
	for {
		candidates := []Prism{}
		for _, prism := range prisms {
			prismPos := Position{x: prism.x, y: prism.y, angle: 0}
			if distance(current, prismPos) < 1e-9 {
				continue
			}
			if angleEquals(angleTo(current, prismPos), current.angle) {
				candidates = append(candidates, prism)
			}
		}
		if len(candidates) == 0 {
			break
		}

		closest := candidates[0]
		for _, candidate := range candidates[1:] {

			candidatePos := Position{x: candidate.x, y: candidate.y}
			closestPos := Position{x: closest.x, y: closest.y}

			if distance(current, candidatePos) < distance(current, closestPos) {
				closest = candidate
			}
		}
		result = append(result, closest.id)
		current = Position{x: closest.x, y: closest.y, angle: current.angle + closest.angle}

	}
	return result
}

func angleTo(from, to Position) float64 {
	dx := to.x - from.x
	dy := to.y - from.y
	return math.Atan2(dy, dx) * 180 / math.Pi
}

func angleEquals(a, b float64) bool {
	diff := math.Mod(a-b, 360)
	if diff > 180 {
		diff -= 360
	} else if diff < -180 {
		diff += 360
	}
	return math.Abs(diff) < 1e-2
}

func distance(a, b Position) float64 {
	return math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2))
}
