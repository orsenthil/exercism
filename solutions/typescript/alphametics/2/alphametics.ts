type Letters = { [key: string]: number };

function sumWords(dictionary: Letters, words: string[]): number {
  let sum = 0;
  for (let word of [...words]) {
    sum += Number(word.split('').map(item => dictionary[item]).join(''));
    }
    return sum;
}

function checkPuzzle(dictionary:Letters, wl: string[], wr: string[]): boolean {
  for (let word of wl) {
    if (dictionary[word[0]] === 0) {
      return false;
    }
  }
  for (let word of wr) {
    if (dictionary[word[0]] === 0) {
      return false;
    }
  }

  if (sumWords(dictionary, wl) === sumWords(dictionary, wr)) {
    return true;
  }
  return false;
}

function generatePermutations(wl:string[], wr:string[], dictionary:Letters, idx: number = 0, list: number[] =[0, 1, 2, 3, 4, 5, 6, 7, 8, 9]): Letters[] {
  let output: Letters[] = [];

  for (let np in list) {
    let key = Object.keys(dictionary)[idx];
    dictionary[key] = list[np];
    if (idx < Object.keys(dictionary).length - 1) {
      output = generatePermutations(wl, wr, dictionary, idx + 1, list.filter(item => item !== list[np]));

      if (output.length > 0) {
        break
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
  let wordsLeft: string[] = sides[0].split('+');
  let wordsRight: string[] = sides[1].split('+');
  let letters = puzzle.replace(/[^A-Z]/g,'');
  letters.split('').forEach(letter => {dictionary[letter] = 0});

  let output: Letters[] = [];
  output = generatePermutations(wordsLeft, wordsRight, dictionary);

  return output[0];
}