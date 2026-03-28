//
// This is only a SKELETON file for the 'Pop Count' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export const eggCount = (displayValue: number): number => {
  let eggs = 0
  displayValue.toString(2).split('').forEach((bit) => {
    if (bit === '1') {
      eggs++
    }
  })
  return eggs
}
