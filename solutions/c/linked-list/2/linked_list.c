#include <errno.h>
#include <stdlib.h>
#include "linked_list.h"

struct list_node {
   struct list_node *prev, *next;
   ll_data_t data;
};

struct list {
   struct list_node *first, *last;
   size_t size;
};

struct list* list_create(void)
{
    struct list* list = (struct list*) malloc(sizeof(struct list));
    list->first = list->last = NULL;
    list->size = 0;
    return list;
}

size_t list_count(const struct list* list)
{
    return list->size;
}

void list_push(struct list* list, ll_data_t item_data)
{
    struct list_node* node = (struct list_node*) malloc(sizeof(struct list_node));
    node->prev = list->last;
    node->next = NULL;
    node->data = item_data;

    if (node->prev == NULL)
        list->first = node;
    else
        node->prev->next = node;

    list->last = node;
    list->size += 1;
}

ll_data_t list_pop(struct list* list)
{
    if (list->size <= 0)
    {
        errno = ENODATA;
        return -1;
    }

    struct list_node* node = list->last;
    list->last = node->prev;
    node->prev = NULL;

    if (list->last != NULL)
        list->last->next = NULL;
    else
        list->first = NULL;

    list->size -= 1;

    ll_data_t data = node->data;
    free(node);
    return data;
}

void list_unshift(struct list* list, ll_data_t item_data)
{
    struct list_node* node = (struct list_node*) malloc(sizeof(struct list_node));
    node->prev = NULL;
    node->next = list->first;
    node->data = item_data;

    if (node->next == NULL)
        list->last = node;
    else
        list->first->prev = node;

    list->first = node;
    list->size += 1;
}

ll_data_t list_shift(struct list* list)
{
    if (list->size <= 0)
    {
        errno = ENODATA;
        return -1;
    }

    struct list_node* node = list->first;
    list->first = node->next;
    node->next = NULL;

    if (list->first != NULL)
        list->first->prev = NULL;
    else
        list->last = NULL;

    list->size -= 1;

    ll_data_t data = node->data;
    free(node);
    return data;
}

void list_delete(struct list* list, ll_data_t data)
{
    struct list_node* curr = list->first;

    while (curr != NULL)
    {
        if (curr->data != data)
        {
            curr = curr->next;
            continue;
        }

        if (curr->prev != NULL)
            curr->prev->next = curr->next;
        else
            list->first = curr->next;

        if (curr->next != NULL)
            curr->next->prev = curr->prev;
        else
            list->last = curr->prev;

        list->size -= 1;
        curr->prev = curr->next = NULL;
        free(curr);
        break;
    }
}

void list_destroy(struct list* list)
{
    while (list->first != NULL)
    {
        struct list_node* next = list->first->next;
        free(list->first);
        list->first = next;
    }

    free(list);
}
