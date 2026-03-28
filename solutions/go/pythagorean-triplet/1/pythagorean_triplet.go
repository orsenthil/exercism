package pythagorean

import "math"

type Triplet [3]int

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) []Triplet {
	var result []Triplet
	for a := min; a <= max; a++ {
		for b := a + 1; b <= max; b++ {
			c := int(math.Sqrt(float64(a*a + b*b)))
			if c <= max && a*a+b*b == c*c {
				result = append(result, Triplet{a, b, c})
			}

		}
	}
	return result
}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) []Triplet {
	var result []Triplet
	for a := 1; a < p/3; a++ {
		for b := a + 1; b < p/2; b++ {
			c := p - a - b
			if a*a+b*b == c*c {
				result = append(result, Triplet{a, b, c})
			}
		}
	}
	return result
}
