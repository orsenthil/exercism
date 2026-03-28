export class Rational {
  numerator: number
  denominator: number

  constructor(numerator: number, denominator: number) {
    this.numerator = numerator
    this.denominator = denominator
  }

  add(other: Rational) {
    const numerator = this.numerator * other.denominator + other.numerator * this.denominator
    const denominator = this.denominator * other.denominator
    if (numerator === 0) {
        return new Rational(0, 1)
    }
    return new Rational(numerator, denominator)
  }

  sub(other: Rational) {
    const numerator = this.numerator * other.denominator - other.numerator * this.denominator
    const denominator = this.denominator * other.denominator
    if (numerator === 0) {
        return new Rational(0, 1)
    }
    return new Rational(numerator, denominator)
  }

  mul(other: Rational) {
      const numerator = this.numerator * other.numerator
      const denominator = this.denominator * other.denominator
      if (numerator === 0) {
          return new Rational(0, 1)
      }
      if (denominator < 0) {
          return new Rational(-numerator, -denominator)
      }
      return new Rational(numerator, denominator).reduce()
  }

  div(other: Rational) {
      const numerator = this.numerator * other.denominator
      const denominator = this.denominator * other.numerator
      return new Rational(numerator, denominator).reduce()
  }

  abs() {
    const numerator = Math.abs(this.numerator)
    const denominator = Math.abs(this.denominator)
    return new Rational(numerator, denominator).reduce()
  }

  exprational(exp: number) {
    if (exp === 0) {
        return new Rational(1, 1)
    }
    if (exp < 0) {
        const numerator = Math.pow(this.denominator, Math.abs(exp))
        const denominator = Math.pow(this.numerator, Math.abs(exp))
        if (denominator < 0) {
            return new Rational(-numerator, -denominator)
        }
        return new Rational(numerator, denominator)
    }
    const numerator = Math.pow(this.numerator, exp)
    const denominator = Math.pow(this.denominator, exp)
    if (numerator === 0) {
        return new Rational(0, 1)
    }
    if (denominator < 0) {
        return new Rational(-numerator, -denominator)
    }
    return new Rational(numerator, denominator)
  }

  expreal(exp: number) {
      // raise the real number to the power of the rational number
      const numerator = Math.pow(exp, this.numerator)
      return Math.pow(numerator, 1 / this.denominator)
  }

  reduce() {
      const gcd = (a: number, b: number): number => b === 0 ? a: gcd(b, a % b)
      const divisor = gcd(this.numerator, this.denominator)
      const numerator = this.numerator / divisor
      const denominator = this.denominator / divisor
      if (denominator < 0) {
          return new Rational(-numerator, -denominator)
      }
      if (numerator === 0) {
          return new Rational(0, 1)
      }

      return new Rational(numerator, denominator)

  }
}
