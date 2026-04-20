#include "gigasecond.h"

void gigasecond(time_t input, char *output, size_t size) {
  time_t result =
      input + 1000000000L; // Add one gigasecond (1,000,000,000 seconds)
  struct tm *t = gmtime(&result);
  strftime(output, size, "%Y-%m-%d %H:%M:%S", t);
}
