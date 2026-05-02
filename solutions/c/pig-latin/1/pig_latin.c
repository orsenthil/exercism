#include "pig_latin.h"
#include <string.h>
#include <stdlib.h>

static int is_vowel(char c) {
    return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u';
}

static char *translate_word(const char *word) {
    // find the split point i such that word[0..i-1] are the leading consonants
    // then build word[i..] + word[0..i-1] + "ay"
    int i = 0;
    if (is_vowel(word[0]) || strncmp(word, "xr", 2) == 0 || strncmp(word, "yt", 2) == 0) {
        i = 0;
    } else {
        // Scan consonants until you hit a vowel
        // OR if you hit "qu" OR if you hit "y" (after at least one consonant)
        while (word[i] != '\0' && !is_vowel(word[i]) && strncmp(word + i, "qu", 2) != 0) {
            if (word[i] == 'y' && i > 0) {
                break;
            }
            i++;
        }
        if (strncmp(word + i, "qu", 2) == 0) {
            i += 2;
        }
    }
    char *result = calloc(strlen(word) + 3, sizeof(char));
    memcpy(result, word + i, strlen(word) - i); // copy the rest of the word
    memcpy(result + strlen(word) - i, word, i); // copy the leading consonants
    strcat(result, "ay");
    return result;
}

char *translate(const char *phrase) {
    // Find each word's start and end, extract it with memcpy into a temp buffer
    // Call translate_word then stich results together
    char *result = calloc(strlen(phrase) * 2 + 2, sizeof(char));
    size_t start = 0;
    size_t end = 0;
    while (phrase[end] != '\0') {
        if (phrase[end] == ' ') {
            char *word = calloc(end - start + 1, sizeof(char));
            memcpy(word, phrase + start, end - start);
            strcat(result, translate_word(word));
            free(word);
            strcat(result, " ");
            start = end + 1;
        }
        end++;
    }
    char *word = calloc(end - start + 1, sizeof(char));
    memcpy(word, phrase + start, end - start);
    strcat(result, translate_word(word));
    return result;
}
