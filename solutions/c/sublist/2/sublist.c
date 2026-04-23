#include "sublist.h"

int matches_at(int *list_a, size_t len_a, int *list_b, size_t start) {
    for (size_t i = 0; i < len_a; i++) {
        if (list_a[i] != list_b[start + i]) {
            return 0;
        }
    }
    return 1;
}

comparison_result_t check_lists(int *list_to_compare, int *base_list,
                                size_t list_to_compare_element_count,
                                size_t base_list_element_count) {
    if (list_to_compare_element_count == 0 && base_list_element_count == 0) {
        return EQUAL;
    }
    if (list_to_compare_element_count == 0) {
        return SUBLIST;
    }
    if (base_list_element_count == 0) {
        return SUPERLIST;
    }
    if (list_to_compare_element_count == base_list_element_count) {
        if (matches_at(list_to_compare, list_to_compare_element_count, base_list, 0)) {
            return EQUAL;
        }
        return UNEQUAL;
    }
    if (base_list_element_count > list_to_compare_element_count) {
        for (size_t i = 0; i <= base_list_element_count - list_to_compare_element_count; i++) {
            if (matches_at(list_to_compare, list_to_compare_element_count, base_list,  i)) {
                return SUBLIST;
            }
        }
    } else {
        for (size_t i = 0; i <= list_to_compare_element_count - base_list_element_count; i++) {
            if (matches_at(base_list, base_list_element_count, list_to_compare,  i)) {
                return SUPERLIST;
            }
        }
    }
    return UNEQUAL;
}
