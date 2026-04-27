#include "run_length_encoding.h"
#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <ctype.h>

char *encode(const char *text) {
    char *encoded = malloc((4 * strlen(text)) + 1);
    size_t encoded_index = 0;
    size_t text_index = 0;
    while (text[text_index] != '\0') {
        char current_char = text[text_index];
        size_t count = 1;
        while (text[text_index + 1] == current_char) {
            count++;
            text_index++;
        }
        if (count > 1) {
            char buf[12];
            int len = sprintf(buf, "%zd", count);
            for (int i = 0; i < len; i++) {
                encoded[encoded_index++] = buf[i];
            }
          }
        encoded[encoded_index++] = current_char;
        text_index++;
    }   
    encoded[encoded_index] = '\0';
    return encoded;
}

char *decode(const char *data) {
    char *decoded = malloc((100 * strlen(data) + 1));
    size_t decoded_index = 0;
    size_t data_index = 0;
    while (data[data_index] != '\0') {
        size_t count = 0;
        while (isdigit(data[data_index])) {
            count = count * 10 + (data[data_index] - '0');
            data_index++;
        }
        char current_char = data[data_index];
        if (count == 0) {
            count = 1;
        }
        for (size_t i = 0; i < count; i++) {
            decoded[decoded_index++] = current_char;
        }
        data_index++;
    }
    decoded[decoded_index] = '\0';
    return decoded;
}
