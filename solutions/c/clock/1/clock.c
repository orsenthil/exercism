#include "clock.h"
#include <stdio.h>
#include <string.h>

static clock_t normalize(int hour, int minute) {
    int total = hour * 60 + minute;
    total = (total % (24 * 60) + (24 * 60)) % (24 * 60);
    int h = total / 60;
    int m = total % 60;
    clock_t clock;
    snprintf(clock.text, MAX_STR_LEN, "%02d:%02d", h, m);
    return clock;
}

clock_t clock_create(int hour, int minute) {
    return normalize(hour, minute);
}
clock_t clock_add(clock_t clock, int minute_add) {
    int hour = (clock.text[0] - '0') * 10 + (clock.text[1] - '0');
    int minute = (clock.text[3] - '0') * 10 + (clock.text[4] - '0');
    return normalize(hour, minute + minute_add);
}
clock_t clock_subtract(clock_t clock, int minute_subtract) {
    int hour = (clock.text[0] - '0') * 10 + (clock.text[1] - '0');
    int minute = (clock.text[3] - '0') * 10 + (clock.text[4] - '0');
    return normalize(hour, minute - minute_subtract);
}
bool clock_is_equal(clock_t a, clock_t b) {
    return strcmp(a.text, b.text) == 0;
}
