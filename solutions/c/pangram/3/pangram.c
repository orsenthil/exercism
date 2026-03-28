#include "pangram.h"

bool all(const int *alphabet, size_t n)
{
    for(size_t i = 0; i < n; i++)
    {
        if (alphabet[i] < 1)
            return false;
    }
    return true;
}

char* to_lower(const char *sentence)
{
    size_t n = strlen(sentence);
    char* lower = calloc(n, sizeof(char));

    for(size_t i = 0; i < n; i++)
    {
        if (sentence[i] >= 'A' && sentence[i] <= 'Z')
            lower[i] = sentence[i] - 'A' + 'a';
        else
            lower[i] = sentence[i];
    }
    return lower;
}

bool is_pangram(const char *sentence)
{
    if (sentence == NULL) {
        return false;
    }

    int* alphabet = calloc(26, sizeof(int));
    char* lowercase = to_lower(sentence);

    for (int i = 0; lowercase[i] != '\0'; i++) {
        alphabet[lowercase[i] - 'a'] += 1;
    }

    return all(alphabet, 26);
}
