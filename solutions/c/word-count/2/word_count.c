#include "word_count.h"
#include <ctype.h>
#include <string.h>

int count_words(const char *sentence, word_count_word_t *words)
{
    memset(words, 0, sizeof(word_count_word_t) * 20);
    char str[strlen(sentence)];
    strcpy(str, sentence);
    char *p = str;
    int arrcount = 0;
    int len = (int)strlen(str);

    for (int i = 0; i <= len; i++) {
        str[i] = tolower(str[i]);

        // regex not ANSI so manually walking through the string
        if (!isalnum(str[i]) && str[i] != '\'') {
            str[i] = '\0';
            if (strlen(p) > MAX_WORD_LENGTH)
                return EXCESSIVE_LENGTH_WORD;

            if (p[0] == '\'' && p[strlen(p) - 1] == '\'')
            {
                p[strlen(p) - 1] = '\0';
                p += 1;
            }

            if (strlen(p) == 0) {
                p = str + i + 1;
                continue;
            }
            int j = 0;
            for (j = 0; j < arrcount; j++) {
                if (!strcmp(words[j].text, p)) {
                    words[j].count += 1;
                    break;
                }
            }

            if (j == arrcount) {
                if (arrcount == 20)
                    return EXCESSIVE_NUMBER_OF_WORDS;
                strcpy(words[arrcount].text, p);
                words[arrcount].count += 1;
                arrcount += 1;
            }

            p = str + i + 1;
        }
    }
    return arrcount;
}
