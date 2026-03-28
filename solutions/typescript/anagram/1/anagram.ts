export class Anagram {
  word: string;

  constructor(input: string) {
    this.word = input;
  }

  public matches(...potentials: string[]): string[]{
    let result: string[] = [];
    potentials.forEach((potential) => {
      if (this.isAnagram(potential)) {
        result.push(potential);
      }
    });
    return result;

  }

  private isAnagram(potential: string): boolean {
    if (this.word.length !== potential.length) {
      return false;
    }
    if (this.word.toLowerCase() === potential.toLowerCase()) {
      return false;
    }
    return this.sortString(this.word) === this.sortString(potential);
  }

  private sortString(input: string): string {
    return input.toLowerCase().split('').sort().join('');
  }

}

