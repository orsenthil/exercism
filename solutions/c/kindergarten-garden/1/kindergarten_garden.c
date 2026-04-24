#include "kindergarten_garden.h"

#include <string.h>


plant_t get_plant_name(char plant) {
    switch (plant) {
        case 'G':
            return GRASS;
        case 'C':
            return CLOVER;
        case 'R':
            return RADISHES;
        case 'V':
            return VIOLETS;
    }
    return -1;
}


plants_t plants(const char *diagram, const char *student) {
    /* calculate the index of the student in the alphabet   */
    int index = student[0] - 'A';
    /* calculate the index of the plants in the diagram */
    int plants_index = index * 2;
    /* Find the index of \n in the diagram */
    char *newline_index = strchr(diagram, '\n');
    int row_two_index = (int)(newline_index - diagram) + 1 + plants_index;

    plants_t plants = { .plants = { get_plant_name(diagram[plants_index]), get_plant_name(diagram[plants_index + 1]), get_plant_name(diagram[row_two_index]), get_plant_name(diagram[row_two_index + 1]) } };
    return plants;
}
