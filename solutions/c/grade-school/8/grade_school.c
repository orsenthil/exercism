#include "grade_school.h"
#include <string.h>

static bool exist(roster_t *roster, char *name) {
    for (size_t i = 0; i < roster->count; i++) {
        if (strcmp(name, roster->students[i].name) == 0) {
            return true;
        }
    }
    return false;
}

static size_t grade_idx(roster_t *roster, uint8_t grade) {
    for (size_t i = 0; i < roster->count; i++) {
        if (roster->students[i].grade >= grade) {
            return i;
        }
    }
    return roster->count;
}

static size_t insert_idx(roster_t *roster, uint8_t grade, char name[MAX_NAME_LENGTH]) {
    for (size_t i = grade_idx(roster, grade); i < roster->count; i++) {
        student_t current = roster->students[i];
        if (current.grade > grade || strcmp(current.name, name) >= 0) {
            return i;
        }
    }
    return roster->count;
}

void init_roster(roster_t *roster) {
    roster->count = 0;
}

bool add_student(roster_t *roster, char name[MAX_NAME_LENGTH], uint8_t grade) {
    if (exist(roster, name)) {
        return false;
    }

    student_t s = {.grade = grade};
    strcpy(s.name, name);

    size_t idx = insert_idx(roster, grade, name);

    memmove(&roster->students[idx + 1], &roster->students[idx], (roster->count - idx) * sizeof(student_t));

    roster->students[idx] = s;
    roster->count++;
    return true;
}

roster_t get_grade(roster_t *roster, uint8_t grade) {
    size_t start = grade_idx(roster, grade);
    size_t stop = grade_idx(roster, grade + 1);

    roster_t r;
    init_roster(&r);

    for (size_t i = start; i < stop; i++) {
        student_t current = roster->students[i];
        add_student(&r, current.name, current.grade);
    }
    return r;
}

