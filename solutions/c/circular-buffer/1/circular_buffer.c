#include "circular_buffer.h"
#include <stdlib.h>
#include <errno.h>

circular_buffer_t *new_circular_buffer(size_t capacity) {
    circular_buffer_t *buffer = (circular_buffer_t *)malloc(sizeof(circular_buffer_t));
    buffer->data = (buffer_value_t *)malloc(capacity * sizeof(buffer_value_t));
    buffer->capacity = capacity;
    buffer->count = 0;
    buffer->read_pos = 0;
    buffer->write_pos = 0;
    return buffer;
}
void delete_buffer(circular_buffer_t *buffer) {
    free(buffer->data);
    free(buffer);
}
int16_t write(circular_buffer_t *buffer, buffer_value_t value) {
    if (buffer->count == buffer->capacity) {
        errno = ENOBUFS;
        return EXIT_FAILURE;
    }
    buffer->data[buffer->write_pos] = value;
    buffer->write_pos = (buffer->write_pos + 1) % buffer->capacity;
    buffer->count++;
    return EXIT_SUCCESS;
}
int16_t overwrite(circular_buffer_t *buffer, buffer_value_t value) {
    if (buffer->count == buffer->capacity) {
        buffer->data[buffer->write_pos] = value;
        buffer->write_pos = (buffer->write_pos + 1) % buffer->capacity;
        buffer->read_pos = (buffer->read_pos + 1) % buffer->capacity;
        return EXIT_SUCCESS;
    }
    return write(buffer, value);
}
int16_t read(circular_buffer_t *buffer, buffer_value_t *value) {
    if (buffer->count == 0) { 
        errno = ENODATA;
        return EXIT_FAILURE;
    }
    *value = buffer->data[buffer->read_pos];
    buffer->read_pos = (buffer->read_pos + 1) % buffer->capacity;
    buffer->count--;
    return EXIT_SUCCESS;
}
void clear_buffer(circular_buffer_t *buffer) {
    buffer->count = 0;
    buffer->read_pos = 0;
    buffer->write_pos = 0;
}
