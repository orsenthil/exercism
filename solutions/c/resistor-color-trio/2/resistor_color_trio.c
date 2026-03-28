#include "resistor_color_trio.h"
#include <math.h>
#include <stdint.h>


resistor_value_t color_code(const resistor_band_t* colors)
{
    resistor_value_t resistor_value;

    uint16_t result;
    result = (colors[0] * 10 + colors[1]) * pow(10, colors[2]);

    if (result >= 1000) {
        result /= 1000;
        resistor_value.unit = KILOOHMS;
    }
    resistor_value.value = result;

    return resistor_value;
}
