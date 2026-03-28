#include "resistor_color.h"

resistor_band_t expected[] = {
      BLACK, BROWN, RED, ORANGE, YELLOW,
      GREEN, BLUE, VIOLET, GREY, WHITE
};


int color_code(resistor_band_t color) {
    if (color == BLACK) {
        return 0;
    }
    if (color == BROWN) {
        return 1;
    }
    if (color == RED) {
        return 2;
    }
    if (color == ORANGE) {
        return 3;
    }

    if (color == YELLOW) {
        return 4;
    }

    if (color == GREEN) {
        return 5;
    }

    if (color == BLUE) {
        return 6;
    }

    if (color == VIOLET) {
        return 7;
    }

    if (color == GREY) {
        return 8;
    }

    if (color == WHITE) {
        return 9;
    }
    return -1;
}

resistor_band_t * colors() {

    return expected;
}