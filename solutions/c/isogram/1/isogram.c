#include "isogram.h"
#include <ctype.h>
#include <string.h>

bool is_isogram(const char phrase[]) {
  if (phrase == NULL) {
    return false;
  }
  int len = strlen(phrase);
  int seen = 0;
  int bit = 0;
  char c;

  for (int i = 0; i < len; i++) {
    c = tolower(phrase[i]);

    if (c >= 'a' && c <= 'z') {
      bit = 1 << (c - 'a');

      if (bit & seen) {
        return false;
      }
      seen |= bit;
    }
  }

  return true;
}
