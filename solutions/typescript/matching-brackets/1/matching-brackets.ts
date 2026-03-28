export function isPaired(input: string): boolean{
  let stack = [];
  let openBrackets = ['{', '[', '('];
  let closeBrackets = ['}', ']', ')'];
  let bracketPairs: { [key: string]: string } = {
    '}': '{',
    ']': '[',
    ')': '('
  }
  while (input.length > 0){
    let char = input[0];
    input = input.substring(1);
    if (openBrackets.includes(char)){
      stack.push(char);
    } else if (closeBrackets.includes(char)){
      if (stack.length === 0 || stack[stack.length - 1] !== bracketPairs[char]){
        return false;
      }
      stack.pop();
    }
  }
  return stack.length === 0;
}
