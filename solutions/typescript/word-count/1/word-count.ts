export function count(words: string): Map<string, number>{
    result = new Map();
    let word = '';
    for (let i = 0; i < words.length; i++) {
        if (words[i].match(/[\w']/)) {
            word += words[i];
        } else {
            if (word.length > 0) {
                result.set(word, result.get(word) + 1 || 1);
                word = '';
            }
        }
    }

    return result;
}
