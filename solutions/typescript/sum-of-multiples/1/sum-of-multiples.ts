export function sum(basevalues: number[], level: number) {
  let result = 0;
  for (let i = 1; i < level; i++) {
    for (let j = 0; j < basevalues.length; j++) {
      if (i % basevalues[j] === 0) {
        result += i;
        break;
      }
    }
  }

  return result;
}
