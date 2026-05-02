#include "meetup.h"
#include <string.h>

static int dayofweek(int y, int m, int d) {
    static int t[] = {0, 3, 2, 5, 0, 3, 5, 1, 4, 6, 2, 4};
    if (m < 3) y--;
    return (y + y/4 - y/100 + y/400 + t[m-1] + d) % 7;
}

static int parse_day(const char *day_of_week) {
    // return 0 for "Sunday", 1 for "Monday", ... 6 for "Saturday"
    if (strcmp(day_of_week, "Sunday") == 0) return 0;
    if (strcmp(day_of_week, "Monday") == 0) return 1;
    if (strcmp(day_of_week, "Tuesday") == 0) return 2;
    if (strcmp(day_of_week, "Wednesday") == 0) return 3;
    if (strcmp(day_of_week, "Thursday") == 0) return 4;
    if (strcmp(day_of_week, "Friday") == 0) return 5;
    if (strcmp(day_of_week, "Saturday") == 0) return 6;
    return -1;
}

static int days_in_month(unsigned int year, unsigned int month) {
    int days[] = {31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31};
    if (month == 2) {
        // leap year: divisible by 4, except centuries, except 400-year centuries
        if ((year % 4 == 0 && year % 100 != 0) || year % 400 == 0)
            return 29;
    }
    return days[month - 1];
}

int meetup_day_of_month(unsigned int year, unsigned int month, const char *week, const char *day_of_week) {
    int starting_day_of_the_month = 1;
    if (strcmp(week, "first") == 0) {
        starting_day_of_the_month = 1;
    } else if (strcmp(week, "second") == 0) {
        starting_day_of_the_month = 8;
    } else if (strcmp(week, "third") == 0) {
        starting_day_of_the_month = 15;
    } else if (strcmp(week, "fourth") == 0) {
        starting_day_of_the_month = 22;
    } else if (strcmp(week, "teenth") == 0) {
        starting_day_of_the_month = 13;
    } else if (strcmp(week, "last") == 0) {
        starting_day_of_the_month = days_in_month(year, month) - 6;
    }
    int day = starting_day_of_the_month;
    while (dayofweek(year, month, day) != parse_day(day_of_week)) {
        day++;
    }
    return day;
}
