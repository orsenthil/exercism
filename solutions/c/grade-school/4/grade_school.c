#include "grade_school.h"
#include "string.h"

void init_roster(roster_t *r) {
    r->count = 0;
}

bool add_student(roster_t *r, char* name, int grade) {
    int to_move_index = r->count;
    bool rv = true;
    student_t new_student;
    new_student.grade = grade;
    strcpy(new_student.name, name);

    if (r->count == 0) {
        r->students[r->count] = new_student;
        ++r->count;
        goto end;
    }
    else {
       for (int i = r->count - 1; i >= 0; i--) {
           char temp_name[MAX_NAME_LENGTH];
           strcpy(temp_name, r->students[i].name);
           int value = strcmp(temp_name, name);

           if (value == 0) {
               rv = false;
               goto end;
           } else {
               if (r->students[i].grade == grade) {
                   if (value > 0) to_move_index = i;
                   else
                       goto move_students;
               }
               else if (r->students[i].grade < grade) {
                   goto move_students;
               }
               else {
                   to_move_index = i;
               }
           }
       }
    }
    move_students:
    for (int j = r->count; j > to_move_index -1 && j >0; --j) {
        r->students[j+1] = r->students[j];
    }
    r->students[to_move_index] = new_student;
    ++r->count;
    rv = true;
    goto end;

    end:
    return rv;
}

roster_t get_grade(roster_t *r, int grade) {
    static roster_t rv;
    init_roster(&rv);

    for (size_t i = 0; i < r->count; ++i) {
        if (r->students[i].grade == grade) {
            add_student(&rv, r->students[i].name, r->students[i].grade);
        }
    }
    return rv;
}