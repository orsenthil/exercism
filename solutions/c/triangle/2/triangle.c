#include "triangle.h"

static bool check_violation(triangle_t sides) {
    if ((sides.a + sides.b <= sides.c) || (sides.b + sides.c <= sides.a) || (sides.a + sides.c <= sides.b)) {
        return true;
    }
    return false;
}


bool is_equilateral(triangle_t sides) {
    if (check_violation(sides)) {
        return false;
    }
    if ((sides.a == sides.b) && (sides.b == sides.c)) {
        return true;
    }
    return false;
}

bool is_isosceles(triangle_t sides) {
    if (check_violation(sides)) {
        return false;
    }
    if ((sides.a == sides.b) || (sides.a == sides.c) || (sides.b == sides.c)){
        return true;
    }
    return false;
}

bool is_scalene(triangle_t sides) {
    if (check_violation(sides)) {
        return false;
    }
    if ((sides.a != sides.b) && (sides.b != sides.c)) {
        return true;
    }
    return false;
}
