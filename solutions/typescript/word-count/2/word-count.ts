export function count(words: string): Map<string, number>{
    let wordMap = new Map<string, number>();

    for (let word of words.match(/\b\w+('\w+)?\b/g) || []) {
        word = word.toLowerCase();
        wordMap.set(word, (wordMap.get(word) || 0) + 1);
    }

    return wordMap;

}
