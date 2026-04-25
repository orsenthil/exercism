#include "anagram.h"
#include <string.h>
#include <ctype.h>
#include <strings.h>

void build_freq(const char *word, int freq[26]) {
    for (size_t i = 0; i < strlen(word); i++) {
        char c = tolower(word[i]);
        if (c >= 'a' && c <= 'z') {
            freq[c - 'a']++;
        }
    }
}

void find_anagrams(const char *subject, struct candidates *candidates) {
    int subject_freq[26] = {0};
    build_freq(subject, subject_freq);
    for (size_t i = 0; i < candidates->count; i++) {
        candidates->candidate[i].is_anagram = UNCHECKED;

        const char *candidate = candidates->candidate[i].word;
        int candidate_freq[26] = {0};
        build_freq(candidate, candidate_freq);
        if (strcasecmp(subject, candidate) == 0) {
            candidates->candidate[i].is_anagram = NOT_ANAGRAM;
        } else if (memcmp(subject_freq, candidate_freq, sizeof(subject_freq)) == 0) {
            candidates->candidate[i].is_anagram = IS_ANAGRAM;
        } else {
            candidates->candidate[i].is_anagram = NOT_ANAGRAM;
        }
    }
}
