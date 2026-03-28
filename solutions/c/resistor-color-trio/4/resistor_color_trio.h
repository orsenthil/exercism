#ifndef RESISTOR_COLOR_TRIO_H
#define RESISTOR_COLOR_TRIO_H

typedef enum {
    BLACK, BROWN, RED, ORANGE, YELLOW, GREEN, BLUE, VIOLET, GREY, WHITE
} resistor_band_t;

typedef enum {
    OHMS, KILOOHMS
} resistor_value_unit;

typedef struct {
    int value;
    resistor_value_unit unit;
} resistor_value_t;

resistor_value_t color_code(const resistor_band_t* colors);
#endif
