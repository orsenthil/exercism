#include "binary_search_tree.h"

#include <assert.h>
#include <stdlib.h>

static void insert_node(node_t **tree, node_t *node)
{
    if (*tree == NULL) {
        *tree = node;
        return;
    }

    if ((*tree)->data >= node->data)
        insert_node(&(*tree)->left, node);
    else
        insert_node(&(*tree)->right, node);
}

static node_t* create_node(int value)
{
    node_t* node = calloc(1, sizeof(*node));
    if (!node)
        return NULL;
    node->data = value;
    return node;
}

static size_t len_tree(const node_t *tree)
{
    if (!tree) return 0;
    return 1 + len_tree(tree->left) + len_tree(tree->right);
}

node_t * build_tree(int *tree_data, size_t tree_data_len)
{
    assert(tree_data != NULL);
    assert(tree_data_len != 0);

    node_t *root = NULL;

    for(size_t i = 0; i < tree_data_len; i++) {
        node_t* node = create_node(tree_data[i]);
        if (!node)
            goto error;
        insert_node(&root, node);
    }
    return root;

error:
    free_tree(root);
    return NULL;
}

void free_tree(node_t* tree)
{
    if (tree == NULL)
        return;
    free_tree(tree->left);
    free_tree(tree->right);
    free(tree);
}

typedef void(*for_each_func)(node_t*, void*);

static void for_each_sorted(node_t* tree, for_each_func func, void* data)
{
    if (!tree)
        return;
    for_each_sorted(tree->left, func, data);
    (*func)(tree, data);
    for_each_sorted(tree->right, func, data);
}

struct _store_data {
    int *array;
    int index;
};

static void store(node_t *node, void *data)
{
    struct _store_data *sd = data;
    sd->array[sd->index++] = node->data;
}

int* sorted_data(node_t* tree)
{
    if (!tree) return NULL;

    int *data = malloc(sizeof(*data) * len_tree(tree));
    if (!data) return NULL;

    struct _store_data sd = {
            .array = data,
            .index = 0,
    };
    for_each_sorted(tree, &store, &sd);
    return data;
}
