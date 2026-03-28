export class Allergies {
  allergenIndex: number
  allergens: Map<number, string>
  allergies: string[]

  constructor(allergenIndex: number) {
    this.allergies = []
    this.allergenIndex = allergenIndex
    this.allergens = new Map<number, string>([
      [128, 'cats'],
      [64, 'pollen'],
      [32, 'chocolate'],
      [16, 'tomatoes'],
      [8, 'strawberries'],
      [4, 'shellfish'],
      [2, 'peanuts'],
      [1, 'eggs']
    ])

    let helper = allergenIndex % 256
    for (const value of this.allergens.keys()) {
      if (helper >= value) {
        const allergen = this.allergens.get(value)
        if (allergen) {
          this.allergies.push(allergen)
        }
        helper -= value
      }
    }
  }

  public list(): string[] {
    return this.allergies.reverse()
  }

  public allergicTo(allergen: string): boolean {
    return this.allergies.includes(allergen)
 } 

}
