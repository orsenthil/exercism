#include "react.h"
#include <stdlib.h>
#include <string.h>

#define MAX_CELLS 128
#define MAX_CALLBACKS 128

// define struct reactor

typedef struct reactor {
   struct cell *cells[MAX_CELLS];
   int num_cells;
} reactor_t;

// define struct callback

typedef struct callback {
    int id;
    int active;
    void *data;
    void (*callback)(void *data, int value);
} callback_t;


// define struct cell

typedef struct cell {
   enum {
      CELL_INPUT,
      CELL_COMPUTE1,
      CELL_COMPUTE2,
   } type;
   int value;
   struct cell *deps[2];
   union {
    compute1 fn1;
    compute2 fn2;
   } union_fn;
   struct callback callbacks[MAX_CALLBACKS];
   int num_callbacks;
   int next_id;
   struct reactor *reactor;
} cell_t;


struct reactor *create_reactor(void) {
    reactor_t *reactor = malloc(sizeof(reactor_t));
    reactor->num_cells = 0;
    return reactor;
}

void destroy_reactor(struct reactor *reactor) {
    for (int i = 0; i < reactor->num_cells; i++) {
        free(reactor->cells[i]);
    }
    free(reactor);
}

struct cell *create_input_cell(struct reactor *reactor, int initial_value) {
    cell_t *cell = malloc(sizeof(cell_t));
    memset(cell, 0, sizeof(*cell));
    cell->type = CELL_INPUT;
    cell->value = initial_value;
    reactor->cells[reactor->num_cells] = cell;
    reactor->num_cells++;
    cell->reactor = reactor;
    return cell;
}

struct cell *create_compute1_cell(struct reactor *reactor, struct cell *dep, compute1 fn) {
    cell_t *cell = malloc(sizeof(cell_t));
    memset(cell, 0, sizeof(*cell));
    cell->value = fn(dep->value);
    cell->type = CELL_COMPUTE1;
    cell->deps[0] = dep;
    cell->union_fn.fn1 = fn;
    reactor->cells[reactor->num_cells] = cell;
    reactor->num_cells++;
    cell->reactor = reactor;
    return cell;
}

struct cell *create_compute2_cell(struct reactor *reactor, struct cell *dep1, struct cell *dep2, compute2 fn) {
    cell_t *cell = malloc(sizeof(cell_t));
    memset(cell, 0, sizeof(*cell));
    cell->value = fn(dep1->value, dep2->value);
    cell->type = CELL_COMPUTE2;
    cell->deps[0] = dep1;
    cell->deps[1] = dep2;
    cell->union_fn.fn2 = fn;
    reactor->cells[reactor->num_cells] = cell;
    reactor->num_cells++;
    cell->reactor = reactor;
    return cell;
}

int get_cell_value(struct cell *cell) {
    return cell->value;
}


void set_cell_value(struct cell *cell, int new_value) {
    cell->value = new_value;
    struct reactor *reactor = cell->reactor;
    // for each compute cel in the reactor-cells (in order)
        // safe cell->vaue as old_value somewhere
    
    int old_values[MAX_CELLS];
    for (int i = 0; i < reactor->num_cells; i++) {
        old_values[i] = reactor->cells[i]->value;
    }

    // 2. Recompute all compute cells (your loop -> already correct)
    for (int i = 0; i < reactor->num_cells; i++) {
        cell_t *compute_cell = reactor->cells[i];
        if (compute_cell->type == CELL_COMPUTE1) {
            compute_cell->value = compute_cell->union_fn.fn1(compute_cell->deps[0]->value);
        } else if (compute_cell->type == CELL_COMPUTE2) {
            compute_cell->value = compute_cell->union_fn.fn2(compute_cell->deps[0]->value, compute_cell->deps[1]->value);
        }
    }

    // 3. Call callbacks for each compute cell that has changed (your loop -> already correct)
    for (int i = 0; i < reactor->num_cells; i++) {
        cell_t *compute_cell = reactor->cells[i];
        if (compute_cell->type != CELL_INPUT && compute_cell->value != old_values[i]) {
            for (int j = 0; j < compute_cell->num_callbacks; j++) {
                if (compute_cell->callbacks[j].active) {
                    compute_cell->callbacks[j].callback(compute_cell->callbacks[j].data, compute_cell->value);
                }
            }
        }
    }

}


callback_id add_callback(struct cell *cell, void *data, callback fn) {
    struct callback *cb = &cell->callbacks[cell->num_callbacks++];
    cb->id = cell->next_id++;
    cb->active = 1;
    cb->data = data;
    cb->callback = fn;
    return cb->id;
}

void remove_callback(struct cell *cell, callback_id id) {
    for (int i = 0; i < cell->num_callbacks; i++) {
        if (cell->callbacks[i].id == id) {
            cell->callbacks[i].active = 0;
        }
    }
}
