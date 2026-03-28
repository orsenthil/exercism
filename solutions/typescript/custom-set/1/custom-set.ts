export class CustomSet {
  setOfNumbers: number[] = []

  constructor(initial?: number[]) {
    this.setOfNumbers = initial || []
  }

  empty(): boolean{
    if (this.setOfNumbers.length === 0) {
      return true
    }
    return false
  }

  contains(element: number): boolean {
    if (this.setOfNumbers.includes(element)) {
      return true
    }

    return false
  }

  add(element: number): CustomSet {
    if (!this.setOfNumbers.includes(element)) {
      this.setOfNumbers.push(element)
    }
    return this
  }

  subset(other: CustomSet): boolean {
    if (this.setOfNumbers.length === 0) {
      return true
    }

    for (let i = 0; i < this.setOfNumbers.length; i++) {
      if (!other.setOfNumbers.includes(this.setOfNumbers[i])) {
        return false
      }
    }
    return true
  }

  disjoint(other: CustomSet): boolean {
    if (this.setOfNumbers.length === 0 && other.setOfNumbers.length === 0) {
      return true
    }
    if (this.setOfNumbers.length === 0 || other.setOfNumbers.length === 0) {
      return true
    }

    for (let i = 0; i < this.setOfNumbers.length; i++) {
      if (other.setOfNumbers.includes(this.setOfNumbers[i])) {
        return false
      }
    }
    return true
  }

  eql(other: CustomSet): boolean {
   if (this.setOfNumbers.length !== other.setOfNumbers.length) {
     return false
   }
   if (this.setOfNumbers.length === 0 && other.setOfNumbers.length === 0) {
     return true
   }

    for (let i = 0; i < this.setOfNumbers.length; i++) {
      if (!other.setOfNumbers.includes(this.setOfNumbers[i])) {
        return false
      }
    }
    return true
  }

  union(other: CustomSet): CustomSet {
    let unionSet = new CustomSet()
    for (let i = 0; i < this.setOfNumbers.length; i++) {
      unionSet.add(this.setOfNumbers[i])
    }
    for (let i = 0; i < other.setOfNumbers.length; i++) {
      unionSet.add(other.setOfNumbers[i])
    }
    return unionSet
  }

  intersection(other: CustomSet): CustomSet {
    let intersectionSet = new CustomSet()
    for (let i = 0; i < this.setOfNumbers.length; i++) {
      if (other.setOfNumbers.includes(this.setOfNumbers[i])) {
        intersectionSet.add(this.setOfNumbers[i])
      }
    }
    return intersectionSet
  }

  difference(other: CustomSet): CustomSet {
    let differenceSet = new CustomSet()
    for (let i = 0; i < this.setOfNumbers.length; i++) {
      if (!other.setOfNumbers.includes(this.setOfNumbers[i])) {
        differenceSet.add(this.setOfNumbers[i])
      }
    }

    return differenceSet
  }
}
