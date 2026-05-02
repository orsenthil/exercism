#include "wordy.h"
#include <string.h>
#include <ctype.h>

bool answer(const char *question, int *result) {
    // Verify the question starts with "What is"
    if (strncmp(question, "What is", 7) != 0) {
        return false;
    }

    // Step 2. Parse the first number
    int number = 0;
    int i = 8;
    bool is_negative = false;
    bool has_digits = false;
    if (question[i] == '-') {
        is_negative = true;
        i++;
    }
    while (isdigit(question[i])) {
        has_digits = true;
        number = number * 10 + (question[i] - '0');
        i++;
    }

    if (!has_digits) {
        return false;
    }

    if (is_negative) {
        number = -number;
    }
    bool is_plus = false;
    bool is_minus = false;
    bool is_multiplied_by = false;
    bool is_divided_by = false;

    while (question[i] != '?') {
        // Skip whitespace
        while (isspace(question[i])) {
            i++;
        }
        // Parse plus or minus
        if (strncmp(question + i, "plus", 4) == 0) {
            is_plus = true;
            i += 4;
        } else if (strncmp(question + i, "minus", 5) == 0) {
            is_minus = true;
            i += 5;
        } else if (strncmp(question + i, "multiplied by", 13) == 0) {
            is_multiplied_by = true;
            i += 13;
        } else if (strncmp(question + i, "divided by", 10) == 0) {
            is_divided_by = true;
            i += 10;
        } else {
            return false;
        }
        // Skip past the operator words
        while (isspace(question[i])) {
            i++;
        }
        // Parse the next number
        bool is_negative = false;
        bool has_digits = false;
        if (question[i] == '-') {
            is_negative = true;
            i++;
        }
        int next_number = 0;
        while (isdigit(question[i])) {
            has_digits = true;
            next_number = next_number * 10 + (question[i] - '0');
            i++;
        }
        if (!has_digits) {
            return false;
        }
        if (is_negative) {
            next_number = -next_number;
        }
        if (is_plus) {
            number += next_number;
            is_plus = false;
        } else if (is_minus) {
            number -= next_number;
            is_minus = false;
        } else if (is_multiplied_by) {
            number *= next_number;
            is_multiplied_by = false;
        } else if (is_divided_by) {
            number /= next_number;
            is_divided_by = false;
        }
    }

    *result = number;
    return true;
}
