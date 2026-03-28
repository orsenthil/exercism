
const english1 = [
  "zero",
  "one",
  "two",
  "three",
  "four",
  "five",
  "six",
  "seven",
  "eight",
  "nine",
  "ten",
  "eleven",
  "twelve",
  "thirteen",
  "fourteen",
  "fifteen",
  "sixteen",
  "seventeen",
  "eighteen",
  "nineteen"
];
const english2 = [
  "twenty",
  "thirty",
  "forty",
  "fifty",
  "sixty",
  "seventy",
  "eighty",
  "ninety"
];
const english3 = [
  "thousand",
  "million",
  "billion"
]

// Create an array of digits from a number
function digits(n: number): number[] {
  const d = Math.floor(n / 10);
  const r = n % 10;
  if (d === 0) return [r];
  return digits(d).concat(r);
}

// Group digits into groups of 3
function group(digits: number[]): number[][] {
  if (digits.length <= 3) return [digits];
  return group(digits.slice(0, digits.length - 3)).concat([digits.slice(-3)]);
}

// handle one or two digits
function basic(digits: number[]): string {
  const len = digits.length;
  if (len === 1) return english1[digits[0]];
  if (digits[0] === 0) return digits[1] === 0 ? "" : english1[digits[1]];
  if (digits[0] === 1) return english1[digits[1] + 10];

  const tens = english2[digits[0] - 2];
  const ones = digits[1] === 0 ? "" : "-" + english1[digits[1]];

  return tens + ones;
}

// handle three digits or less
function triplet(digits: number[]): string {
  if (digits.length < 3) return basic(digits);
  const hundreds = digits[0] === 0 ? "" : english1[digits[0]] + " hundred ";
  return `${hundreds}${basic(digits.slice(-2))}`;
}

export function sayInEnglish(n : number): string {
  if (n < 0 || n > 999999999999) {
    throw new Error('Number must be between 0 and 999,999,999,999.')
  }

  const triplets = group(digits(n));
  const mag = triplets.length - 2;

  return triplets.reduce<string[]>((acc, digits, index) => {
    const amount = triplet(digits);
    const scale = mag - index < 0 ? '' : english3[mag - index];
    return amount ? acc.concat(amount, scale) : acc.concat(amount);
  }, []).filter(s => s.length > 0).join(' ').trim();
}
