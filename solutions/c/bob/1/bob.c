#include "bob.h"
#include <stdbool.h>
#include <string.h>
#include <ctype.h>

char *hey_bob(char *greeting) {
    size_t len = strlen(greeting);

    bool is_silence = true;
    for (size_t i = 0; i < len; i++) {
        if (!isspace(greeting[i])) {
            is_silence = false;
            break;
        }
    }
    if (is_silence || len == 0) {
        return "Fine. Be that way!";
    }

    // Check if the greeting is shouting
    bool is_shouting = false;
    for (size_t i = 0; i < len; i++) {
        if (isalpha(greeting[i])) {
            if (islower(greeting[i])) {
                is_shouting = false;
                break;
            }
            is_shouting = true;
        }
    }
    bool is_question = false;
    // strip trailing whitespace
    size_t last_non_whitespace = len - 1;
    for (size_t i = len - 1; i > 0; i--) {
        if (!isspace(greeting[i])) {
            last_non_whitespace = i;
            break;
        }
    }
    if (greeting[last_non_whitespace] == '?') {
        is_question = true;
    }

    if (is_shouting && is_question) {
        return "Calm down, I know what I'm doing!";
    } else if (is_shouting) {
        return "Whoa, chill out!";
    } else if (is_question) {
        return "Sure.";
    } else {
        return "Whatever.";
    }
}
