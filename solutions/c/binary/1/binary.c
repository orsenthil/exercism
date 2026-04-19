#include "binary.h"
#include <string.h>


int convert(const char *input) {
  int len = strlen(input);
  int result = 0;
  int digit = 0;
  for (int i = 0; i < len; i++) {
    if (input[i] != '0' && input[i] != '1') {
      return -1;
    }
    digit = input[i] - '0';
    result = (result * 2) + digit;
    
  }
  return result;
}


