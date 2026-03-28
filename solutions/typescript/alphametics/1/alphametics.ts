type Letters = {[key: string]: number};

function sumWords(dictionary: Letters, words: string[]): number {
  let sum = 0;
  for (let word of [...words]) {
    sum += Number(word.split('').map(item => dictionary[item]).join(''));
  }
  return sum;
}

function checkPuzzle(dictionary:Letters, wl:string[], wr:string[]): boolean {
  for (let word in wl) {
    if (dictionary[word[0]] === 0) {
      return false;
    }
  }
  for (let word in wr) {
    if (dictionary[word[0]] === 0) {
      return false;
    }
  }

  if (sumWords(dictionary, wl) === sumWords(dictionary, wr)) {
    return true;
  }

  return false

}

function generatePermutations(wl: string[], wr: string[], dictionary: Letters, idx: number= 0, list:number[] = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]): Letters[] {
  let output: Letters[] = [];

  for (let np in list) {
    let key = Object.keys(dictionary)[idx];
    dictionary[key] = list[np];
    if ( idx < Object.keys(dictionary).length - 1) {
      output = generatePermutations(wl, wr, dictionary, idx + 1, list.filter(item => item !== list[np]));

      if (output.length) {
        break;
      }
    } else {
      if (checkPuzzle(dictionary, wl, wr)) {
        output.push({...dictionary});
        break;
      }
    }
  }
  return output;
}

export function solve(puzzle: string): Letters {
  let dictionary: Letters = {};
  puzzle = puzzle.replace(/ /g, '');
  let sides = puzzle.split('==');
  let wl = sides[0].split('+');
  let wr = sides[1].split('+');
  let letters = new Set([...wl.join(''), ...wr.join('')]);
  for (let letter of letters) {
    dictionary[letter] = 0;
  }
  return generatePermutations(wl, wr, dictionary)[0];
}
