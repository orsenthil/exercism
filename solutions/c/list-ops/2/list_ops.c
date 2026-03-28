#include "list_ops.h"
#include <stdlib.h>
#include <memory.h>

// constructs a new list
list_t *new_list(size_t length, list_element_t elements[]) {
    list_t *list = malloc(sizeof(list_t) + length * sizeof (list_element_t));
    list->length = length;
    memcpy(list->elements, elements, length * sizeof(list_element_t));
    return list;
}

list_t *append_list(list_t *list1, list_t *list2) {
    size_t list1Length = list1->length;
    size_t list2Length = list2->length;
    size_t combinedLength = list1Length + list2Length;

    list_element_t *elements = malloc(sizeof(*elements) * combinedLength);

    memcpy(elements, list1->elements, list1Length * sizeof(*elements));
    memcpy(elements + list1Length, list2->elements, list2Length * sizeof(*elements));

    return new_list(combinedLength, elements);
}

list_t *filter_list(list_t *list, bool (*filter)(list_element_t)) {
    list_element_t *elements = malloc(sizeof(*elements) * list->length);
    int newListIndex = 0;

    for (size_t i = 0; i < list->length; i++) {
        if (filter(list->elements[i])) {
            elements[newListIndex] = list->elements[i];
            newListIndex++;
        }
    }

    int newLength = newListIndex;
    list_element_t  *newElements = realloc(elements, newLength * sizeof(*newElements));

    return new_list(newLength, newElements);
}

size_t length_list(list_t *list)  {
    return list->length;
}

list_t *map_list(list_t *list, list_element_t (*map)(list_element_t)) {
    list_element_t *transformed = malloc(sizeof(*transformed) * list->length);
    for(size_t i = 0; i < list->length; i++) {
        transformed[i] = map(list->elements[i]);
    }

    return new_list(list->length, transformed);
}

list_element_t foldl_list(list_t *list, list_element_t initial, list_element_t (*foldl)(list_element_t, list_element_t)) {

    list_element_t result = initial;

    for(size_t i = 0; i < list->length; i++) {
        result = foldl(result, list->elements[i]);
    }

    return result;
}

list_element_t foldr_list(list_t *list, list_element_t initial, list_element_t (*foldr)(list_element_t, list_element_t)) {
    list_element_t result = initial;
    if (list->length == 0) {
        return result;
    }

    for (int i = list->length - 1; i>=0; i--) {
        result = foldr(list->elements[i], result);
    }

    return result;
}

list_t *reverse_list(list_t *list) {
    list_element_t *reversed = malloc(sizeof(*reversed) * list->length);
    for(size_t i = 0; i < list->length; i++) {
        reversed[i] = list->elements[list->length - 1 - i];
    }

    return new_list(list->length, reversed);
}

void delete_list(list_t *list) {
    free(list);
}
