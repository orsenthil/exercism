#include "pangram.h"
#include <ctype.h>
#include <stdlib.h>
#include <string.h>

bool is_pangram(const char *sentence) {
  if (sentence == NULL) {
    return false;
  }

  size_t len = strlen(sentence);
  int seen = 0;
  char c;

  for (size_t i = 0; i < len; i++) {
    c = sentence[i];
    if (c >= 'A' && c <= 'Z') {
      c = tolower(c);
    }

    if (c >= 'a' && c <= 'z') {
      seen |= 1 << (c - 'a');
    }
  }

  return seen == ((1 << 26) - 1);
}
