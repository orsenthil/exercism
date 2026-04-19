#include "eliuds_eggs.h"

int egg_count(unsigned int display_value) {
  int result = 0;
  int bit = 0;

  while (display_value) {
    bit = display_value & 1;
    result += bit;
    display_value = display_value >> 1;

  }

  return result;
}

