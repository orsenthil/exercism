export class Rational {
  numerator: number
  denominator: number

  constructor(numerator: number, denominator: number) {
    this.numerator = numerator
    this.denominator = denominator
  }

  formatReal() {
      if (this.numerator === 0) {
          this.denominator = 1
      }
      if (this.denominator < 0) {
          this.numerator = -this.numerator
          this.denominator = -this.denominator
      }
      return this
  }

  add(other: Rational) {
    const numerator = this.numerator * other.denominator + other.numerator * this.denominator
    const denominator = this.denominator * other.denominator
    return new Rational(numerator, denominator).formatReal()
  }

  sub(other: Rational) {
    const numerator = this.numerator * other.denominator - other.numerator * this.denominator
    const denominator = this.denominator * other.denominator
    return new Rational(numerator, denominator).formatReal()
  }

  mul(other: Rational) {
      const numerator = this.numerator * other.numerator
      const denominator = this.denominator * other.denominator
      return new Rational(numerator, denominator).formatReal().reduce()
  }

  div(other: Rational) {
      const numerator = this.numerator * other.denominator
      const denominator = this.denominator * other.numerator
      return new Rational(numerator, denominator).reduce()
  }

  abs() {
    const numerator = Math.abs(this.numerator)
    const denominator = Math.abs(this.denominator)
    return new Rational(numerator, denominator).formatReal().reduce()
  }

  exprational(exp: number) {
    if (exp === 0) {
        return new Rational(1, 1)
    }
    if (exp < 0) {
        const numerator = Math.pow(this.denominator, Math.abs(exp))
        const denominator = Math.pow(this.numerator, Math.abs(exp))
        return new Rational(numerator, denominator).formatReal()
    }
    const numerator = Math.pow(this.numerator, exp)
    const denominator = Math.pow(this.denominator, exp)
    return new Rational(numerator, denominator).formatReal()
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
      return new Rational(numerator, denominator).formatReal()
  }
}
