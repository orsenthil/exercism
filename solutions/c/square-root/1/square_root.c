#include "square_root.h"

uint16_t square_root(uint16_t n) {
    uint16_t root = 0;
    while (root * root < n) {
        root++;
    }
    return root;
}
