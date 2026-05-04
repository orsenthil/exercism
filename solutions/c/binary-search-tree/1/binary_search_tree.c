#include "binary_search_tree.h"
#include <stdlib.h>

static node_t *new_node(int data) {
    node_t *node = (node_t *)malloc(sizeof(node_t));
    node->data = data;
    node->left = NULL;
    node->right = NULL;
    return node;
}

static size_t count_nodes(node_t *node) {
    if (node == NULL) {
        return 0;
    }
    return 1 + count_nodes(node->left) + count_nodes(node->right);
}

static void insert_node(node_t *node, int data) {
    if (data <= node->data) {
        if (node->left == NULL) {
            node->left = new_node(data);
        } else {
            insert_node(node->left, data);
        }
    }
    else {
        if (node->right == NULL) {
            node->right = new_node(data);
        } else {
            insert_node(node->right, data);
        }
    }
}

static void inorder_traversal(node_t *node, int *sorteddata, size_t *size) {
    if (node == NULL) {
        return;
    }
    inorder_traversal(node->left, sorteddata, size);
    sorteddata[*size] = node->data;
    (*size)++;
    inorder_traversal(node->right, sorteddata, size);
}


node_t *build_tree(int *tree_data, size_t tree_data_len) {
    node_t *root = new_node(tree_data[0]);
    for (size_t i = 1; i < tree_data_len; i++) {
        insert_node(root, tree_data[i]);
    }
    return root;
}



void free_tree(node_t *tree) {
    if (tree == NULL) {
        return;
    }
    free_tree(tree->left);
    free_tree(tree->right);
    free(tree);
}

int *sorted_data(node_t *tree) {
    if (tree == NULL) {
        return NULL;
    }
    size_t size = 0;
    size_t index = 0;

    size = count_nodes(tree);
    int *sorteddata = (int *)malloc(sizeof(int) * size);
    inorder_traversal(tree, sorteddata, &index);
    return sorteddata;
}
