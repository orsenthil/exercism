#include "triangle.h"


bool is_equilateral(triangle_t sides) {
    if ((sides.a == sides.b) && (sides.b == sides.c)) {
        return true;
    }
    return false;
}

bool is_isosceles(triangle_t sides) {
    if ((sides.a == sides.b) || (sides.a == sides.c) || (sides.b == sides.c)){
        return true;
    }
    return false;
}

bool is_scalene(triangle_t sides) {
    if ((sides.a != sides.b) && (sides.b != sides.c)) {
        return true;
    }
    return false;
}
