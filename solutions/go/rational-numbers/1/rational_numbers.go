package rationalnumbers

import "math"

type Rational struct {
	numerator, denominator int
}

// Reduce simplifies a Rational, eg changing Rational{4, 2} into Rational{2, 1}.
func (r Rational) Reduce() Rational {
	gcd := gcd(abs(r.numerator), abs(r.denominator))
	if r.denominator < 0  && r.numerator < 0 {
		return Rational{abs(r.numerator / gcd), abs(r.denominator / gcd)}
	}
	if r.denominator < 0 {
		return Rational{-abs(r.numerator / gcd), abs(r.denominator / gcd)}
	}
	return Rational{r.numerator / gcd, r.denominator / gcd}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func (r Rational) Add(s Rational) Rational {
	return Rational{r.numerator * s.denominator + s.numerator * r.denominator, r.denominator * s.denominator}.Reduce()
}

func (r Rational) Sub(s Rational) Rational {
	return Rational{r.numerator * s.denominator - s.numerator * r.denominator, r.denominator * s.denominator}.Reduce()
}

func (r Rational) Mul(s Rational) Rational {
	return Rational{r.numerator * s.numerator, r.denominator * s.denominator}.Reduce()
}

func (r Rational) Div(s Rational) Rational {
	return Rational{r.numerator * s.denominator, r.denominator * s.numerator}.Reduce()
}

func (r Rational) Abs() Rational {
	return Rational{abs(r.numerator), abs(r.denominator)}.Reduce()
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// Compute r ^ power, a rational raised to an int exponent.
func (r Rational) Exprational(power int) Rational {
	if power < 0 {
		return Rational{int(math.Pow(float64(r.denominator), float64(-power))), int(math.Pow(float64(r.numerator), float64(-power)))}.Reduce()
	}
	if power == 0 {
		return Rational{1, 1}
	}
	return Rational{int(math.Pow(float64(r.numerator), float64(power))), int(math.Pow(float64(r.denominator), float64(power)))}.Reduce()
}

// Compute base ^ r, an int raised to a rational.
func (r Rational) Expreal(base int) float64 {
	if r.denominator == 0 {
		return 0
	}
	if r.numerator == 0 {
		return 1
	}
	if r.denominator < 0 {
		return math.Pow(float64(base), float64(-r.numerator)/float64(-r.denominator))
	}
	return math.Pow(float64(base), float64(r.numerator)/float64(r.denominator))
}
