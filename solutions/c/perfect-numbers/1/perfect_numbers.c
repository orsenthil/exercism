#include "perfect_numbers.h"

kind classify_number(int number) {
  if (number <= 0) {
    return ERROR;
  }
  int sum_of_factors = 0;

  for (int i = 1; i <= number - 1; i++) {
    if ((number % i) == 0) {
      sum_of_factors += i;
    }
  }

  if (sum_of_factors == number) {
    return PERFECT_NUMBER;
  }

  if (sum_of_factors < number) {
    return DEFICIENT_NUMBER;
  }

  if (sum_of_factors > number) {
    return ABUNDANT_NUMBER;
  }

  return ERROR;
}
