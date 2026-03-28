#include "space_age.h"

float age(planet_t planet, int64_t seconds) {
    switch (planet) {
        case MERCURY:
            return  (seconds / (31557600)) / 0.2408467;
        case VENUS:
            return  (seconds / (31557600)) / 0.61519726;
        case EARTH:
            return (seconds / (31557600));
        case MARS:
            return (seconds / (31557600)) / 1.8808158;
        case JUPITER:
            return  (seconds / (31557600)) / 11.862615;
        case SATURN:
            return (seconds / (31557600)) / 29.447498;
        case URANUS:
            return (seconds / (31557600)) / 84.016846;
        case NEPTUNE:
            return (seconds / (31557600)) / 164.79132;
        default:
            return -1;
    }

    return 0.0;
}


