#include <stdint.h>
#include "resistor_color.h"

uint16_t color_code(resistor_band_t color) {
    return color;
}

static resistor_band_t bands[] = { BLACK, BROWN, RED, ORANGE, YELLOW, GREEN, BLUE, VIOLET, GREY, WHITE };

const resistor_band_t *colors(void) {
    return bands;
}
