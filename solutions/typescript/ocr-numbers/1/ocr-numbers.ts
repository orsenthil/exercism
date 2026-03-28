export function convert(input: string): string {
  let result = '';
  let lines = input.split('\n');
  let lineCount = lines.length;
  let charCount = lines[0].length;
  let charWidth = charCount / 3;
  let charHeight = lineCount / 4;
  let chars = [];
  for (let i = 0; i < charHeight; i++) {
    chars.push([]);
  }
  for (let i = 0; i < charHeight; i++) {
    for (let j = 0; j < charWidth; j++) {
      let char = '';
      for (let k = 0; k < 3; k++) {
        for (let l = 0; l < 4; l++) {
          char += lines[i * 4 + l][j * 3 + k];
        }
      }
      chars[i].push(char);
    }
  }

  let numbers = [
    ' _ | ||_|',
    '     |  |',
    ' _  _||_ ',
    ' _  _| _|',
    '   |_|  |',
    ' _ |_  _|',
    ' _ |_ |_|',
    ' _   |  |',
    ' _ |_||_|',
    ' _ |_| _|',
  ];

  for (let i = 0; i < charHeight; i++) {
    for (let j = 0; j < charWidth; j++) {
      let char = chars[i][j];
      let index = numbers.indexOf(char);
      if (index === -1) {
        result += '?';
      } else {
        result += index.toString();
      }
    }
  }

  return result;

}
