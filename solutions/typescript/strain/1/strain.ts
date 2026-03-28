export function keep<T>(input: T[], predicate: (element: T) => boolean): T[] {
  let result = [];
  for (let i = 0; i < input.length; i++) {
    if (predicate(input[i])) {
      result.push(input[i]);
    }
  }
  result = result.filter((element) => element !== undefined);
  return result;
}

export function discard<T>(input: T[], predicate: (element: T) => boolean): T[] {
  let result = [];
  for (let i = 0; i < input.length; i++) {
    if (!predicate(input[i])) {
      result.push(input[i]);
    }
  }
  result = result.filter((element) => element !== undefined);
  return result;
}
