export class ComplexNumber {
  realPart : number
  imaginaryPart: number

  constructor(real: number, imaginary: number) {
    this.realPart = real
    this.imaginaryPart = imaginary
  }

  public get real(): number {
    return this.realPart
  }

  public get imag(): number {
    return this.imaginaryPart
  }

  public add(other: ComplexNumber): ComplexNumber{
    return new ComplexNumber(this.realPart + other.realPart, this.imaginaryPart + other.imaginaryPart)
  }

  public sub(other: ComplexNumber): ComplexNumber{
    return new ComplexNumber(this.realPart - other.realPart, this.imaginaryPart - other.imaginaryPart)
  }

  public div(other: ComplexNumber): ComplexNumber{
    /** 
     * Dividing a complex number a + i * b by another c + i * d gives: (a + i * b) / (c + i * d) = (a * c + b * d)/(c^2 + d^2) + (b * c - a * d)/(c^2 + d^2) * i
     * 
     */
    return new ComplexNumber((this.realPart * other.realPart + this.imaginaryPart * other.imaginaryPart) / (other.realPart ** 2 + other.imaginaryPart ** 2),
                              (this.imaginaryPart * other.realPart - this.realPart * other.imaginaryPart) / (other.realPart ** 2 + other.imaginaryPart ** 2))
  }

  public mul(other: ComplexNumber): ComplexNumber{
    return new ComplexNumber(this.realPart * other.realPart - this.imaginaryPart * other.imaginaryPart, 
                             this.realPart * other.imaginaryPart + this.imaginaryPart * other.realPart)
  }

  public get abs(): number {
    return Math.sqrt(this.realPart ** 2 + this.imaginaryPart ** 2)
  }

  public get conj(): ComplexNumber{
    if (this.imaginaryPart == 0) {
      return new ComplexNumber(this.realPart, 0)
    }

    return new ComplexNumber(this.realPart, -this.imaginaryPart)
  }

  public get exp(): ComplexNumber {
    /**
     * Raising e to a complex exponent can be expressed as e^(a + i * b) = e^a * e^(i * b), the last term of which is given by Euler's formula e^(i * b) = cos(b) + i * sin(b).
     */
    let realPart = Math.exp(this.realPart)
    return new ComplexNumber(Math.cos(this.imaginaryPart) * realPart, Math.sin(this.imaginaryPart) * realPart)
  }
}
