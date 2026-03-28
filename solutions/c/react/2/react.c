#include "react.h"

#include <stdlib.h>

typedef enum {
    INPUT,
    COMPUTE1,
    COMPUTE2,
} cell_t;


// A list of cells

struct reactor
{
    struct cell *cells;
};

struct compute1_cell {
    struct cell *input;
    compute1 func;
};

struct compute2_cell {
    struct cell *input1;
    struct cell *input2;
    compute2 func;
};

// The list of cells that depend on 'cell'

struct depender {
    struct cell *cell;
    struct depender *next;
};

// A cell can be one of the 3 types
// 1. Input
// 2. Compute1
// 3. Compute3

struct cell {
    cell_t type;
    int value;
    union {
        struct compute1_cell c1;
        struct compute2_cell c2;
    } compute;
    struct callback_t *callbacks;
    struct depender *dependers;
    struct cell *next;
};

struct callback_t {
    callback_id id;
    callback func;
    void *data;
    struct callback_t *next;
};

struct reactor* create_reactor() {
    return calloc(1, sizeof(struct reactor));
}

static void destroy_cell(struct cell *cell) {
    while (cell->callbacks) {
        struct callback_t *cb = cell->callbacks;
        cell->callbacks = cb->next;
        free(cb);
    }
    while (cell->dependers) {
        struct depender *dep = cell->dependers;
        cell->dependers = dep->next;
        free(dep);
    }
    free(cell);
}

void destroy_reactor(struct reactor *reactor) {
    while (reactor->cells) {
        struct cell *cell = reactor->cells;
        reactor->cells = cell->next;
        destroy_cell(cell);
    }
    free(reactor);
}

struct cell* create_input_cell(struct reactor *reactor, int initial_value) {
    struct cell *cell = calloc(1, sizeof(*cell));
    if (!cell)
        return NULL;
    cell->type = INPUT;
    cell->value = initial_value;
    cell->next = reactor->cells;
    reactor->cells = cell;

    return cell;
}

static int calc_cell_value(struct cell *cell) {
    switch (cell->type) {
        case COMPUTE1:
            return (cell->compute.c1.func)(calc_cell_value(cell->compute.c1.input));
        case COMPUTE2:
            return(cell->compute.c2.func)(
                    calc_cell_value(cell->compute.c2.input1),
                    calc_cell_value(cell->compute.c2.input2));
        default:
            return cell->value;
    }
}

struct cell* create_compute1_cell(struct reactor *reactor, struct cell *input, compute1 func)
{
    if (!reactor || !input)
        return NULL;
    struct cell *cell = calloc(1, sizeof(*cell));

    if (!cell)
        return NULL;

    cell->type = COMPUTE1;
    cell->compute.c1.input = input;
    cell->compute.c1.func = func;
    cell->next = reactor->cells;

    struct depender *dep = malloc(sizeof(*dep));

    if (!dep) {
        free(cell);
        return NULL;
    }

    dep->cell = cell;
    dep->next = input->dependers;
    input->dependers = dep;

    cell->value = calc_cell_value(cell);

    reactor->cells = cell;

    return cell;
}

struct cell * create_compute2_cell(struct reactor *reactor, struct cell *input1, struct cell *input2, compute2 func) {
    if (!reactor || !input1 || !input2)
        return NULL;

    struct cell *cell = calloc(1, sizeof(*cell));

    if (!cell)
        return NULL;

    cell->type = COMPUTE2;
    cell->compute.c2.input1 = input1;
    cell->compute.c2.input2 = input2;
    cell->compute.c2.func = func;
    cell->next = reactor->cells;

    struct depender *dep1 = malloc(sizeof(*dep1));

    if (!dep1) {
        free(cell);
        return NULL;
    }

    dep1->cell = cell;
    dep1->next = input1->dependers;

    struct depender *dep2 = malloc(sizeof(*dep2));

    if (!dep2) {
        free(dep1);
        free(cell);
        return NULL;
    }

    dep2->cell = cell;
    dep2->next = input2->dependers;

    input1->dependers = dep1;
    input2->dependers = dep2;

    cell->value = calc_cell_value(cell);
    reactor->cells = cell;

    return cell;
}

int get_cell_value(struct cell *cell) {
    return cell->value;
}

static void update_dependers(struct cell *cell) {

    for(struct depender *depender = cell->dependers; depender != NULL; depender = depender->next) {
        struct cell *dep = depender->cell;
        int oldval = dep->value;
        dep->value = calc_cell_value(dep);
        if (dep->value == oldval)
            continue;
        for(struct callback_t *cb = dep->callbacks; cb != NULL; cb = cb->next)
            (cb->func)(cb->data, dep->value);

        update_dependers(dep);
    }
}

void set_cell_value(struct cell *cell, int new_value) {
    if (!cell || cell->type != INPUT || cell->value == new_value)
        return;

    cell->value = new_value;
    update_dependers(cell);
}

callback_id add_callback(struct cell *cell, void *data, callback func)
{
    if (!cell || cell->type == INPUT)
        return -1;

    struct callback_t *cb = malloc(sizeof(*cb));
    if (!cb)
        return -1;

    cb->id = cell->callbacks ? cell->callbacks->id + 1: 0;
    cb->func = func;
    cb->data = data;
    cb->next = cell->callbacks;
    cell->callbacks = cb;

    return cb->id;
}

void remove_callback(struct cell *cell, callback_id id) {
    if (!cell || cell->type == INPUT)
        return;

    struct callback_t *prev = NULL;

    for (struct callback_t *cb = cell->callbacks; cb != NULL; cb = cb->next) {
        if (cb->id == id) {
            if (!prev)
                cell->callbacks = cb->next;
            else
                prev->next = cb->next;
            free(cb);
            break;
        }

        prev = cb;
    }
}