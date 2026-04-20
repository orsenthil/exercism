#include "resistor_color_trio.h"
#include <math.h>

resistor_value_t color_code(resistor_band_t bands[]) {
  long value = 0;
  value = (bands[0] * 10 + bands[1]) * (long)pow(10, bands[2]);

  if (value >= 1000000000) {
    value /= 1000000000;
    return (resistor_value_t){.value = value, .unit = GIGAOHMS};
  }

  if (value >= 1000000) {
    value /= 1000000;
    return (resistor_value_t){.value = value, .unit = MEGAOHMS};
  }

  if (value >= 1000) {
    value /= 1000;
    return (resistor_value_t){.value = value, .unit = KILOOHMS};
  }

  return (resistor_value_t){.value = value, .unit = OHMS};
}
