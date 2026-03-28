export const toRoman = (integer: number): string => {
  // convert the integer to a roman numeral

  let roman = '';
  let num = integer;

  const romanNumList = {
    M: 1000,
    CM: 900,
    D: 500,
    CD: 400,
    C: 100,
    XC: 90,
    L: 50,
    XL: 40,
    X: 10,
    IX: 9,
    V: 5,
    IV: 4,
    I: 1
  };

  for (let key in romanNumList) {
    let value = romanNumList[key as keyof typeof romanNumList];
    roman += key.repeat(Math.floor(num / value));
    num = num % value;
  }

  return roman;
}
