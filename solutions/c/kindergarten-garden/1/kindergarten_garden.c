#include "kindergarten_garden.h"

plants_t plants(const char *garden, const char *student) {
    const char *it = garden;

    while (*(it++) != '\n') {
        if (*it == '\0') {
            break;
        }
    }

    plants_t p;
    int index = *student - 'A';
    p.plants[0] = garden[2 * index];
    p.plants[1] = garden[2 * index + 1];
    p.plants[2] = it[2 * index];
    p.plants[3] = it[2 * index + 1];

    return p;
}
