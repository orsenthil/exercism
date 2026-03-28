#include "space_age.h"

float age(planet_t planet, int64_t seconds) {
    switch (planet) {
        case MERCURY:
            return 0.2408467 * ((31.69 * 1000000000) / seconds);
        case VENUS:
            return 0.61519726 * ((31.69 * 1000000000) / seconds);
        case EARTH:
            return ((31.69 * 1000000000) / seconds);
        case MARS:
            return 1.8808158 * ((31.69 * 1000000000) / seconds);
        case JUPITER:
            return 11.862615 * ((31.69 * 1000000000) / seconds);
        case SATURN:
            return 29.447498 * ((31.69 * 1000000000) / seconds);
        case URANUS:
            return 84.016846 * ((31.69 * 1000000000) / seconds);
        case NEPTUNE:
            return 164.79132 * ((31.69 * 1000000000) / seconds);
    }

    return 0.0;
}


