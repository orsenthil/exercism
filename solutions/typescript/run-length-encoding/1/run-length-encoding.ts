export function encode(input: string): string {
  let result = '';
  let count = 1;
  for (let i = 0; i < input.length; i++) {
    if (input[i] === input[i + 1]) {
      count++;
    } else {
      if (count > 1) {
        result += count + input[i];
      } else {
        result += input[i];
      }
      count = 1;
    }
  }
  return result;
}

export function decode(input: string): string {
  let result = '';
  let count = '';
  for (let i = 0; i < input.length; i++) {
    if (input[i].match(/\d/)) {
      count += input[i];
    } else {
      if (count) {
        result += input[i].repeat(Number(count));
        count = '';
      } else {
        result += input[i];
      }
    }
  }
  return result;
}
