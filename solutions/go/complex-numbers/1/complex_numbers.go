package complexnumbers

import "math"

// Define the Number type here.

type Number struct {
	a float64
	b float64
}

func (n Number) Real() float64 {
	return n.a
}

func (n Number) Imaginary() float64 {
	return n.b
}

func (n1 Number) Add(n2 Number) Number {
	return Number{n1.a + n2.a, n1.b + n2.b}
}

func (n1 Number) Subtract(n2 Number) Number {
	return Number{n1.a - n2.a, n1.b - n2.b}
}

func (n1 Number) Multiply(n2 Number) Number {
	return Number{n1.a*n2.a - n1.b*n2.b, n1.a*n2.b + n1.b*n2.a}
}

func (n Number) Times(factor float64) Number {
	return Number{n.a * factor, n.b * factor}
}

func (n1 Number) Divide(n2 Number) Number {
	return Number{(n1.a*n2.a + n1.b*n2.b) / (n2.a*n2.a + n2.b*n2.b), (n1.b*n2.a - n1.a*n2.b) / (n2.a*n2.a + n2.b*n2.b)}
}

func (n Number) Conjugate() Number {
	return Number{n.a, -n.b}
}

func (n Number) Abs() float64 {
	return math.Sqrt(n.a*n.a + n.b*n.b)
}

func (n Number) Exp() Number {
	// Raising e to a complex exponent can be expressed as `e^(a + i * b) = e^a * e^(i * b)`, the last term of which is given by Euler's formula `e^(i * b) = cos(b) + i * sin(b)`.

	// e^a
	ea := Number{math.Exp(n.a), 0}

	// e^(i * b)

	eb := Number{math.Cos(n.b), math.Sin(n.b)}

	return ea.Multiply(eb)

}
