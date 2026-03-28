#include "wordy.h"

#include <stdio.h>
#include <string.h>
#include <assert.h>

bool answer(const char *question, int *result)
{
    assert(question);
    assert(result);

    int value = 0, charsProcessed = 0;
    size_t questionSize = strlen(question);

    if (0 == sscanf(question, "What is %d %n", result, &charsProcessed)) {
        return false;
    }

    question += charsProcessed;
    questionSize -= charsProcessed;

    while (questionSize > 1) {
        if (sscanf(question, " plus %d %n", &value, &charsProcessed)) {
            *result += value;
        } else if (sscanf(question, " minus %d %n", &value, &charsProcessed)) {
            *result -= value;
        } else if (sscanf(question, " divided by %d %n", &value, &charsProcessed)) {
            *result /= value;
        } else if (sscanf(question, " multiplied by %d %n", &value, &charsProcessed)) {
            *result *= value;
        } else {
            return false;
        }

        question += charsProcessed;
        questionSize -= charsProcessed;
    }

    if (strcmp(question, "?"))
        return false;

    return true;
}
