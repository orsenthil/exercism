#include "linked_list.h"
#include <stdlib.h>

struct list_node {
   struct list_node *prev, *next;
   ll_data_t data;
};

struct list {
   struct list_node *first, *last;
};


// constructs a new (empty) list
struct list *list_create(void)
{
    struct list *list = malloc(sizeof(struct list));
    list->first = NULL;
    list->last = NULL;
    return list;
}

// counts the items on a list
size_t list_count(const struct list *list){
   size_t count = 0;
   struct list_node *node = list->first;
   while (node != NULL) {
      count++;
      node = node->next;
   }
   return count;
}

// inserts item at back of a list
void list_push(struct list *list, ll_data_t item_data) {
   struct list_node *node = malloc(sizeof(struct list_node));
   node->data = item_data;
   node->next = NULL;
   node->prev = list->last;
   if (list->last != NULL) {
      list->last->next = node;
   }
   list->last = node;
   if (list->first == NULL) {
      list->first = node;
   }
}

// removes item from back of a list
ll_data_t list_pop(struct list *list) {
   if (list->last == NULL) {
      return 0;
   }
   ll_data_t data = list->last->data;
   struct list_node *node = list->last;
   list->last = list->last->prev;
   if (list->last == NULL) {
      list->first = NULL;
   }
   if (list->last != NULL) {
      list->last->next = NULL;
   }
   free(node);
   return data;
}

// inserts item at front of a list
void list_unshift(struct list *list, ll_data_t item_data) {
   struct list_node *node = malloc(sizeof(struct list_node));
   node->data = item_data;
   node->next = list->first;
   node->prev = NULL;
   if (list->last == NULL) {
      list->last = node;
   }
   if (list->first != NULL) {
      list->first->prev = node;
   }
   list->first = node;
}

// removes item from front of a list
ll_data_t list_shift(struct list *list) {
   if (list->first == NULL) {
      return 0;
   }
   ll_data_t data = list->first->data;
   struct list_node *node = list->first;
   list->first = list->first->next;
   if (list->first == NULL) {
      list->last = NULL;
   }
   if (list->first != NULL) {
      list->first->prev = NULL;
   }
   free(node);
   return data;
}

// deletes a node that holds the matching data
void list_delete(struct list *list, ll_data_t data) {
   struct list_node *node = list->first;
   while (node != NULL) {
      if (node->data == data) {
         if (node->prev != NULL) {
            node->prev->next = node->next;
         }
         if (node->next != NULL) {
            node->next->prev = node->prev;
         }
         if (node == list->first) {
            list->first = node->next;
         }
         if (node == list->last) {
            list->last = node->prev;
         }
         free(node);
         return;
      } else {
         node = node->next;
      }
   }
   return;
}

// destroys an entire list
// list will be a dangling pointer after calling this method on it
void list_destroy(struct list *list) {
   struct list_node *node = list->first;
   while (node != NULL) {
      struct list_node *next = node->next;
      free(node);
      node = next;
   }
   free(list);
}
