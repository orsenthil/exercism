#include "word_count.h"
#include <ctype.h>
#include <string.h>

static int is_word_char(const char *s, int i) {
    char c = s[i];
    if (isalnum(c)) return 1;
    if (c == '\'' && i > 0 && isalnum(s[i-1]) && isalnum(s[i+1])) return 1;
    return 0;
}

int count_words(const char *sentence, word_count_word_t *words) {
    int word_count = 0;
    for (int i = 0; sentence[i] != '\0'; i++) {
        if (word_count >= MAX_WORDS) return EXCESSIVE_NUMBER_OF_WORDS;
        if (is_word_char(sentence, i)) {
            int word_length = 0;
            while (is_word_char(sentence, i + word_length)) {
                word_length++;
            }
            char temp[MAX_WORD_LENGTH + 1];
            strncpy(temp, sentence + i, word_length);
            temp[word_length] = '\0';
            for (int k = 0; k < word_length; k++) {
                temp[k] = tolower(temp[k]);
            }
            int found = 0;
            for (int j = 0; j < word_count; j++) {
                if (strcmp(words[j].text, temp) == 0) {
                    words[j].count++;
                    found = 1;
                    break;
                }
            }
            if (!found) {
                strncpy(words[word_count].text, temp, MAX_WORD_LENGTH + 1);
                words[word_count].count = 1;
                word_count++;
            }
            i += word_length - 1;
        }
    }
    return word_count;
}
