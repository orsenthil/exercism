package piglatin
import "strings"

func Sentence(sentence string) string {

	// Split the sentence into words

	words := strings.Split(sentence, " ")
	var pigLatinWords []string

	for _, word := range words {
		pigLatinWords = append(pigLatinWords, ConvertToPigLatin(word))
	}

	return strings.Join(pigLatinWords, " ")

}

func ConvertToPigLatin(sentence string) string {
	// Rule 1: if a word starts with a vowel, or "xr", "yt", add "ay" to the end
	if isVowel(string(sentence[0])) || isXrOrYt(sentence) {
		return sentence + "ay"
	}

	// Rule 2: If a word begins with a consonant, move the consonant to the end of the word and add "ay"
	for i, s := range sentence {
		if isVowel(string(s)) {
			index := i
			// Rule 3: If a word starts with zero or more consonants followed by "qu", move the consonants and "qu" to the end of the word and add "ay"
			if string(s) == "u" && string(sentence[i-1]) == "q" {
				index = i + 1
			}
			return sentence[index:] + sentence[:index] + "ay"
		}
	}

	// Rule 4: If a word starts with one or more consonants followed by "y", move the consonants and "y" to the end of the word and add "ay"
	for i, s := range sentence {
		if string(s) == "y" {
			index := i
			return sentence[index:] + sentence[:index] + "ay"
		}
	}

	return sentence

}

func isXrOrYt(sentence string) bool {
	start := sentence[:2]
	if start == "xr" || start == "yt" {
		return true
	}
	return false
}

func isVowel(s string) bool {
	vowels := []string{"a", "e", "i", "o", "u"}
	for _, v := range vowels {
		if s == v {
			return true
		}
	}
	return false
}
