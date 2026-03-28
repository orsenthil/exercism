export function isIsogram(word: string): Boolean{
  word = word.replaceAll(' ', '').replaceAll('-', '').toLowerCase();
  let wordArr = word.split('');
  let wordSet = new Set(wordArr);
  if(wordArr.length === wordSet.size){
    return true;
  }
  return false;
}
