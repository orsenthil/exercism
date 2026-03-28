export function sayInEnglish(largeNumber: number): string {
  let result = '';

  let number = largeNumber;
  let remainder = 0;
  let divisor = 0;
  let divisorName = '';

  if (number < 0 || number > 999999999999) {
    throw new Error('Number must be between 0 and 999,999,999,999.');
  }

  if (number === 0) {
    return 'zero';
  }

  const numberNames = [
    'one', 'two', 'three', 'four', 'five', 'six', 'seven', 'eight', 'nine'
  ];

  const teensNames = [
    'eleven', 'twelve', 'thirteen', 'fourteen', 'fifteen', 'sixteen',
    'seventeen', 'eighteen', 'nineteen'
  ];

  const tensNames = [
    'ten', 'twenty', 'thirty', 'forty', 'fifty', 'sixty', 'seventy',
    'eighty', 'ninety'
  ];

  const largeNumberNames = [
    'thousand', 'million', 'billion'
  ];

  let largeNumberIndex = 0;

  while (number > 0) {
    if (number >= 1000) {
      divisor = 1000;
      divisorName = largeNumberNames[largeNumberIndex];
    } else {
      divisor = 100;
      divisorName = 'hundred';
    }

    remainder = number % divisor;
    number = Math.floor(number / divisor);

    if (number > 0) {
      result = sayInEnglish(number) + ' ' + divisorName + ' ' + result;
    }

    if (remainder > 0) {
      if (result !== '') {
        result += ' ';
      }

      if (remainder < 10) {
        result += numberNames[remainder - 1];
      } else if (remainder < 20) {
        result += teensNames[remainder - 11];
      } else {
        const tens = Math.floor(remainder / 10);
        const ones = remainder % 10;

        result += tensNames[tens - 1];

        if (ones > 0) {
          result += '-' + numberNames[ones - 1];
        }
      }
    }

    largeNumberIndex++;
  }

  return result;
}
