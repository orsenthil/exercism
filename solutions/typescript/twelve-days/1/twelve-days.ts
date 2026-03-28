function getVerse(day: number): string {
  let dayNames = [
    "first",
    "second",
    "third",
    "fourth",
    "fifth",
    "sixth",
    "seventh",
    "eighth",
    "ninth",
    "tenth",
    "eleventh",
    "twelfth",
  ];

  let gifts = [
    "a Partridge in a Pear Tree",
    "two Turtle Doves",
    "three French Hens",
    "four Calling Birds",
    "five Gold Rings",
    "six Geese-a-Laying",
    "seven Swans-a-Swimming",
    "eight Maids-a-Milking",
    "nine Ladies Dancing",
    "ten Lords-a-Leaping",
    "eleven Pipers Piping",
    "twelve Drummers Drumming",
  ];

  let verse = `On the ${dayNames[day - 1]} day of Christmas my true love gave to me: `;
  for (let i = day - 1; i >= 0; i--) {
    if (i === 0 && day > 1) {
      verse += `and ${gifts[i]}`;
    } else {
      verse += `${gifts[i]}`;
    }
    if (i > 0) {
      verse += ", ";
    }
  }
  return verse;
}

export function recite(start: number, end: number): string {
  let result: string[] = [];
  for (let i = start; i <= end; i++) {
    result.push(getVerse(i));
  }
  if (start == end) {
    result[result.length - 1] = result[result.length - 1] + ".\n";
  } else {
    for (let i = 0; i < result.length; i++) {
      result[i] = result[i] + ".\n";
    }
  }

  let lyrics = "".concat(...result);
  console.log(lyrics);
  return lyrics;
}
