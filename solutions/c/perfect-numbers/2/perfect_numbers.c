#include "perfect_numbers.h"

kind classify_number(int number) {
  if (number <= 0) {
    return ERROR;
  }
  int sum_of_factors = 0;

  if (number == 1) {
    return DEFICIENT_NUMBER;
  }

  for (int i = 1; i * i <= number; i++) {
    if ((number % i) == 0) {
      sum_of_factors += i;
      if ((i != number / i) && (number / i != number)) {
        sum_of_factors += (number / i);
      }
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
