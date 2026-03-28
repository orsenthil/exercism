type Tuple = [number, number];

type WordLocation = {
  start: Tuple
  end: Tuple
};

type Result = {[word: string]: WordLocation| undefined};

export class WordSearch {
  private puzzle: string[][];
  private maxRow: number;
  private maxCol: number;


  constructor(grid: string[]) {
    const puzzle = [];
    for (const word in grid) {
      puzzle.push(grid[word].split(''));
    }

    this.puzzle = puzzle;
    this.maxRow = puzzle.length;
    this.maxCol = puzzle[0].length;

  }

  private search(
    word: string,
    r: number,
    c: number,
    dr: number,
    dc: number
  ): WordLocation | undefined {
    const start: Tuple = [r, c];

    for (const ch of word) {
      if (this.puzzle[r-1]?.[c-1] !== ch) {
        return undefined;
      }
      r += dr;
      c += dc;
    }

    return {start, end: [r - dr, c - dc]};
  }

  private findOne(word: string, r: number, c: number): WordLocation | undefined {
    for (let dr = -1; dr <= 1; dr++) {
      for (let dc = -1; dc <= 1; dc++) {
        const result = this.search(word, r, c, dr, dc);
        if (result) {
          return result;
        }
       }
    }
  }

  public find(words: string[]): Result {
    const result: Result = words.reduce((acc, word) => {
      acc[word] = undefined;
      return acc;
    }, {} as Result);

    for (let r = 1; r <= this.maxRow; r++) {
      for (let c = 1; c <= this.maxCol; c++) {
        for (const word of words) {
          if (!result[word]) {
            const location = this.findOne(word, r, c);
            if (location) {
              result[word] = location;
          }
        }
      }
    }
  }
  return result;
}

}
