#include "queen_attack.h"

attack_status_t can_attack(position_t queen_1, position_t queen_2) {
    position_t queen1 = queen_1;
    position_t queen2 = queen_2;

    if (queen1.row >= 8 || queen2.row >= 8 || queen1.column >= 8 || queen2.column >= 8 || (queen1.row == queen2.row && queen1.column == queen2.column))
        return INVALID_POSITION;

    int result = (queen1.row == queen2.row || queen1.column == queen2.column || queen1.column - queen2.column == queen1.row - queen2.row || queen1.column - queen2.column == -(queen1.row - queen2.row));
    return result;
}
