#include "two_fer.h"
#include <stdio.h>

#define BUFSIZE 100

void two_fer(char *buffer, const char *name) {
  if (name == NULL) {
    snprintf(buffer, BUFSIZE, "One for %s, one for me.", "you");
  } else {
    snprintf(buffer, BUFSIZE, "One for %s, one for me.", name);
  }
 }

