#include "allergies.h"

bool is_allergic_to(allergen_t allergen, int score) {
    return score & (1 << allergen);
}

allergen_list_t get_allergens(int score) {
    allergen_list_t result = (allergen_list_t){0, {false, false, false, false, false, false, false, false}};

    for (int i = 0; i < ALLERGEN_COUNT; i++) {
        if (is_allergic_to(i, score)) {
            result.allergens[i] = true;
            result.count++;
        }
    }
    return result;
}
