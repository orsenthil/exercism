export function verse(index: number): string {
  let verse = '';
  switch (index) {
    case 0:
      verse = 'No more bottles of beer on the wall, no more bottles of beer.\n' +
        'Go to the store and buy some more, 99 bottles of beer on the wall.\n';
      break;
    case 1:
      verse = '1 bottle of beer on the wall, 1 bottle of beer.\n' +
        'Take it down and pass it around, no more bottles of beer on the wall.\n';
      break;
    case 2:
      verse = '2 bottles of beer on the wall, 2 bottles of beer.\n' +
        'Take one down and pass it around, 1 bottle of beer on the wall.\n';
      break;
    default:
      verse = `${index} bottles of beer on the wall, ${index} bottles of beer.\n` +
        `Take one down and pass it around, ${index - 1} bottles of beer on the wall.\n`;
      break;
  }

  return verse;
  
}

export function sing(
  initialBottlesCount?: number,
  takeDownCount?:number 
): string {
  let song = '';
  let bottlesCount = initialBottlesCount || 99;
  let takeDown = takeDownCount || 0;
  for (let i = bottlesCount; i >= takeDown; i -= 1) {
    song += verse(i);
    if (i !== takeDown) {
      console.log("I am here")
      song += '\n';
    }
  }
  console.log(song);

  return song;
}
