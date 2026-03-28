export function rotate(input: string, cipher: number): string {
  let result = '';
  for (let i = 0; i < input.length; i++) {
    let charCode = input.charCodeAt(i);
    if (charCode >= 65 && charCode <= 90) {
      result += String.fromCharCode(((charCode - 65 + cipher) % 26) + 65);
    } else if (charCode >= 97 && charCode <= 122) {
      result += String.fromCharCode(((charCode - 97 + cipher) % 26) + 97);
    } else {
      result += input[i];
    }
  }
  return result;
}