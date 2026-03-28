export class Triangle {
  sides: number[]
  triangleInequality: boolean
  constructor(...sides: number[]) {
    this.sides = sides
    this.triangleInequality = this.sides.every((side, i) => side < this.sides[(i + 1) % 3] + this.sides[(i + 2) % 3])
  }

  get isEquilateral() {
    if (!this.triangleInequality) return false
    return this.sides.every(side => side === this.sides[0])
  }

  get isIsosceles() {
    if (!this.triangleInequality) return false
    return this.sides[0] === this.sides[1] || this.sides[1] === this.sides[2] || this.sides[0] === this.sides[2]
  }

  get isScalene() {
    if (!this.triangleInequality) return false
    return this.sides[0] !== this.sides[1] && this.sides[1] !== this.sides[2] && this.sides[0] !== this.sides[2]
  }
}
