type Alphametric = { [key: string]: number };

export function solve(puzzle: string): Alphametric | undefined {
  const mapping: Alphametric = {};
  const letters = [...new Set([...puzzle].filter(c => /[A-Z]/g.test(c)))];
  const exceptions: string[] = puzzle.split(' ').map(c=>/[A-Z]+/g.test(c) ? [...c][0]: '');

  const buildFormula = (quizz: string): string => {
    return quizz.split(' ').map((val: string) => {
      if (/^[+-]|==$/.test(val)) {
        return val;
      } else if (/^[A-Z]+$/.test(val)) {
        return [...val].reverse().map((c, i) => `${c} * ${Math.pow(10, i)}`).join(' + ');
      } else {
        return ''
      }
    }).join(' ');
  }

  const evaluateTest = new Function(...letters, `return ${buildFormula(puzzle)}`);
  const backtrack = (index: number, candidate: number[]): boolean => {
    if (index === letters.length) {
      const integers: number[] = Object.values(mapping);
      return evaluateTest(...integers);
    }

    const letter = letters[index];

    for (const num of candidate) {
      if (num === 0 && exceptions.includes(letter)) {
        continue;
      }

      if (mapping[letter] === undefined) {
        mapping[letter] = num;
      }

      if (backtrack(index + 1, candidate.filter(c => c !== num))) {
        return true;
      }

      delete mapping[letter];
    }
    return false
  }

  const allNumbers = Array.from({ length: 10 }, (_, i) => i);

  if (backtrack(0, allNumbers)) {
    return mapping;
  }

  return undefined
}