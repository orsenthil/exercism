export function proverb(...args: string[]) {
  let result = '';
  for (let i = 0; i < args.length; i++) {
    if (i === 0) {
      result += `For want of a ${args[i]} the ${args[i + 1]} was lost.\n`;
    } else if (i === args.length - 1) {
      result += `And all for the want of a ${args[0]}.`;
    } else {
      result += `For want of a ${args[i]} the ${args[i + 1]} was lost.\n`;
    }
  }
  return result
}
