#include "list_ops.h"
#include <stdlib.h>
#include <string.h>

list_t *new_list(size_t length, list_element_t elements[]) {
    list_t *list = (list_t *)malloc(sizeof(list_t) + length * sizeof(list_element_t));
    list->length = length;
    if (elements != NULL) {
        memcpy(list->elements, elements, length * sizeof(list_element_t));
    }
    return list;
}

void delete_list(list_t *list) {
    free(list);
}

size_t length_list(list_t *list) {
    return list->length;
}


// append entries to a list and return the new list
list_t *append_list(list_t *list1, list_t *list2) {
    list_t *result = new_list(list1->length + list2->length, NULL);
    memcpy(result->elements, list1->elements, list1->length * sizeof(list_element_t));
    memcpy(result->elements + list1->length, list2->elements, list2->length * sizeof(list_element_t));
    return result;
}

// filter list returning only values that satisfy the filter function
list_t *filter_list(list_t *list, bool (*filter)(list_element_t)) {
    list_t *result = new_list(list->length, NULL);
    size_t result_length = 0;
    for (size_t i = 0; i < list->length; i++) {
        if (filter(list->elements[i])) {
            result->elements[result_length++] = list->elements[i];
        }
    }
    result->length = result_length;
    return result;
}

// return a list of elements whose values equal the list value transformed by
// the mapping function
list_t *map_list(list_t *list, list_element_t (*map)(list_element_t)) {
    list_t *result = new_list(list->length, NULL);
    for (size_t i = 0; i < list->length; i++) {
        result->elements[i] = map(list->elements[i]);
    }
    return result;
}

// folds (reduces) the given list from the left with a function
list_element_t foldl_list(list_t *list, list_element_t initial,
                          list_element_t (*foldl)(list_element_t,
                                                  list_element_t)) {
    list_element_t result = initial;
    for (size_t i = 0; i < list->length; i++) {
        result = foldl(result, list->elements[i]);
    }
    return result;
}

// folds (reduces) the given list from the right with a function
list_element_t foldr_list(list_t *list, list_element_t initial,
                          list_element_t (*foldr)(list_element_t,
                                                  list_element_t)) {
    list_element_t result = initial;
    for (size_t i = list->length; i > 0; i--) {
        result = foldr(list->elements[i - 1], result);
    }
    return result;
}

// reverse the elements of the list
list_t *reverse_list(list_t *list) {
    list_t *result = new_list(list->length, NULL);
    for (size_t i = 0; i < list->length; i++) {
        result->elements[i] = list->elements[list->length - i - 1];
    }
    return result;
}
