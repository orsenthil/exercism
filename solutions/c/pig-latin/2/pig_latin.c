#include "pig_latin.h"

#include <stdio.h>
#include <string.h>
#include <stdbool.h>
#include <stdlib.h>


static char VOWELS[] = "aeiou";

static inline bool is_vovel(char c)
{
    return strchr(VOWELS, c);
}

inline static bool is_consonant(char c)
{
    return !is_vovel(c);
}

inline static bool starts_with(const char* string, const char* prefix)
{
    return strstr(string, prefix) == string;
}

static size_t starts_with_consonant_cluster(const char* string)
{
    size_t cluster_size = 0;
    char* pointer = (char*) string;

    while (is_consonant(*pointer)) {
        if (*pointer == 'q' && *(pointer + 1) == 'u') {
            pointer++;
            cluster_size++;
        }

        if (*(pointer + 1) == 'y') {
            return cluster_size + 1;
        }
        pointer += 1;
        cluster_size += 1;
    }
    return cluster_size;
}

static char* apply_rules(const char* phrase)
{
    char* ret = NULL;
    size_t cluster_size;

    if (is_vovel(phrase[0]) || starts_with(phrase, "yt") || starts_with(phrase, "xr")) {
        ret = malloc(strlen(phrase) + 3);
        sprintf(ret, "%say", phrase);
        return ret;
    } else if ((cluster_size = starts_with_consonant_cluster(phrase))) {
        ret = malloc(strlen(phrase) + 3);
        sprintf(ret, "%s%.*say", phrase + cluster_size, (int)cluster_size, phrase);
    }
    return ret;
}

char *translate(const char *phrase)
{
    char* output = malloc(strlen(phrase) * 2);
    char* the_copy = malloc(strlen(phrase) + 1);
    strcpy(the_copy, phrase);

    char* prev_ptr = (char*)the_copy;
    size_t buffer_offset = 0;

    for (char* next_ptr = strchr(the_copy, ' '); next_ptr; prev_ptr = next_ptr + 1, next_ptr = strchr(next_ptr + 1, ' ')) {
        *next_ptr = 0;
        char *with_rules_applied = apply_rules(prev_ptr);
        buffer_offset += sprintf(output + buffer_offset, "%s ", with_rules_applied);

        if (with_rules_applied) {
            free(with_rules_applied);
        }
    }


    char* with_rules_applied = apply_rules(prev_ptr);
    sprintf(output + buffer_offset, "%s", with_rules_applied);

    if (with_rules_applied)
        free(with_rules_applied);

    free(the_copy);
    return output;
}
