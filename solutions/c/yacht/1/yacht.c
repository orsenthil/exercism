#include "yacht.h"

#define DIE_FACES 6
#define NUM_DICE 5

int score(dice_t dice, category_t category) {
    int faces[DIE_FACES] = {0};
    int score = 0;
    int non_zero = 0;

    for (int i = 0; i < NUM_DICE; i++) {
        if (faces[dice.faces[i] - 1] == 0) non_zero++;
        faces[dice.faces[i] - 1]++;
        score += dice.faces[i];
    }

    if (ONES <= category && category <= SIXES) {
        return faces[category] * (category + 1);
    } else if (category == YACHT && non_zero == 1) {
        return 50;
    } else if (category == FULL_HOUSE) {
        for (int i = 0; i < NUM_DICE; i++) {
            if (non_zero != 2 || (faces[dice.faces[i] - 1] != 2 && faces[dice.faces[i] - 1] != 3)) {
                return 0;
            }
            return score;
        }
    } else if (category == FOUR_OF_A_KIND) {
        for (int i = 0; i < NUM_DICE; i++) {
            if (faces[dice.faces[i] - 1] == 4 || faces[dice.faces[i] - 1] == 5) {
                return dice.faces[i] * 4;
            }
        }
    } else if (non_zero == 5 &&
        ((category == LITTLE_STRAIGHT && faces[DIE_FACES -1] == 0) ||
         (category == BIG_STRAIGHT && faces[0] == 0))) {
            return 30;
    } else if (category == CHOICE) {
        return score;
    }
    return 0;
}
