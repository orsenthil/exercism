#include "high_scores.h"

#include <stdint.h>
#include <stdlib.h>
#include <string.h>

int32_t latest(const int32_t *scores, size_t scores_len) {
  return scores[scores_len - 1];
}

int32_t personal_best(const int32_t *scores, size_t scores_len) {
  int32_t highest = scores[0];
  for (int i = 0; i < (int)scores_len; i++) {
    if (scores[i] > highest) {
      highest = scores[i];
    }
  }
  return highest;
}

int compare_desc(const void *a, const void *b) {
  return (*(int32_t *)b - *(int32_t *)a); // descending
}

size_t personal_top_three(const int32_t *scores, size_t scores_len,
                          int32_t *output) {

  int32_t copy[scores_len];
  memcpy(copy, scores, scores_len * sizeof(int32_t));

  qsort(copy, scores_len, sizeof(int32_t), compare_desc);

  output[0] = copy[0];

  size_t count = scores_len < 3 ? scores_len : 3;

  for (size_t i = 1; i < count; i++) {
    output[i] = copy[i];
  }

  return count;
}
