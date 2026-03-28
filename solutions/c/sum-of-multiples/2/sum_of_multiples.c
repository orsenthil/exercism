#include "sum_of_multiples.h"

typedef struct node node_t;
typedef unsigned int tree_data_t;

struct node {
    node_t *right;
    node_t *left;
    tree_data_t data;
};

static void insert_node(node_t **node, tree_data_t data)
{
    if (!*node)
    {
        *node = calloc(1, sizeof(node_t));
        (*node)->data = data;
        return;
    }

    if ((*node)->data == data)
        return;

    if ((*node)->data > data)
        insert_node(&((*node)->left), data);
    else
        insert_node(&((*node)->right), data);
}

static tree_data_t sum_tree(node_t *tree)
{
    size_t size = 100ul, sp = 0;
    tree_data_t result = 0u;

    node_t **stack = calloc(size, sizeof(node_t *));
    while (tree || sp > 0)
    {
        while (tree)
        {
            stack[sp++] = tree;
            if (sp >= size)
            {
                size *= 2;
                stack = realloc(stack, size * sizeof(node_t *));
            }
            tree = tree->left;
        }
        tree = stack[--sp];
        result += tree->data;
        tree = tree->right;
    }
    free(stack);
    return result;
}

static void free_tree(node_t *tree)
{
    if (tree)
    {
        free_tree(tree->left);
        free_tree(tree->right);
        free(tree);
    }
}

unsigned int sum(const unsigned int *factors, const size_t number_of_factors, const unsigned int limit)
{
    node_t *root = calloc(1, sizeof(node_t));
    root->data = 0u;

    for (size_t i = 0; i < number_of_factors; i++)
    {
        if (!factors[i])
            continue;
        for (unsigned int m = 1, p = factors[i]; p < limit; m++, p = m * factors[i])
            insert_node(&root, p);
    }
    unsigned int result = sum_tree(root);
    free_tree(root);
    return result;
}
