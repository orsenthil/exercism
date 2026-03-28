#ifndef GRADE_SCHOOL_H
#define GRADE_SCHOOL_H

#ifndef MYBOOLEAN_H
#define MYBOOLEAN_H

#define false 0
#define true 1
typedef int bool; // or #define bool int

#endif

#include <stddef.h>
#include <stdint.h>

#define MAX_NAME_LENGTH 20
#define MAX_STUDENTS 20

typedef struct {
   uint8_t grade;
   char name[MAX_NAME_LENGTH];
} student_t;

typedef struct {
   size_t count;
   student_t students[MAX_STUDENTS];
} roster_t;

bool add_student(roster_t *r, char* name, int grade);
roster_t get_grade(roster_t *r, int grade);
void init_roster(roster_t *r);



#endif
