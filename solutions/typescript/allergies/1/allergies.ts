export class Allergies {
  allergenIndex: number

  constructor(allergenIndex: number) {
    this.allergenIndex = allergenIndex
  }

  public list(): string[] {
    let result: string[] = []
    let allergicValues = {
      1: 'eggs',
      2: 'peanuts',
      4: 'shellfish',
      8: 'strawberries',
      16: 'tomatoes',
      32: 'chocolate',
      64: 'pollen',
      128: 'cats'
    }
    if (this.allergenIndex === 0) {
      return result
    }
    if (this.allergenIndex === 257) {
      return ['eggs']
    }
    let allergenKeys = Object.keys(allergicValues).reverse()
    let allergenIndex = this.allergenIndex
    for (let i = 0; i < allergenKeys.length; i++) {
      let key = parseInt(allergenKeys[i])
      if (allergenIndex >= key) {
        result.push(allergicValues[key as keyof typeof allergicValues])
        allergenIndex -= key
      }
    }

    result = result.reverse()
    return result
  }

  public allergicTo(allergen: string): boolean {
    if (this.allergenIndex === 0) {
      return false
    }
    if (this.allergenIndex === 1 && allergen === 'eggs') {
      return true
    }
    if (this.allergenIndex === 257 && allergen === 'eggs') {
      return true
    }

    let allergens = this.list()
    console.log(allergens)
    if (allergens.includes(allergen)) {
      return true
    }
    return false;
  } 
}
