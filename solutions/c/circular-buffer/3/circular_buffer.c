#include "circular_buffer.h"

#include <errno.h>
#include <stdlib.h>

typedef enum
{
    BUF_INDEX_READ,
    BUF_INDEX_INSERT,
    BUF_INDEX_COUNT,
} BUF_INDEX_t;

typedef enum
{
    BUF_WRITE_MODE_ERROR_IF_FULL,
    BUF_WRITE_MODE_OVERWRITE_IF_FULL,
    BUF_WRITE_MODE_COUNT,
} BUF_WRITE_MODE_t;

struct circular_buffer_s
{
    size_t capacity;
    size_t indices[BUF_INDEX_COUNT];
    bool is_full;
    buffer_value_t *values;
};

static size_t count_elements_internal(const circular_buffer_t * const buffer)
{
    if (buffer == NULL)
    {
        return 0u;
    }

    if (buffer->is_full)
    {
        return buffer->capacity;
    }

    return (buffer->indices[BUF_INDEX_INSERT] >= buffer->indices[BUF_INDEX_READ])
        ? (buffer->indices[BUF_INDEX_INSERT] - buffer->indices[BUF_INDEX_READ])
        : (buffer->capacity - buffer->indices[BUF_INDEX_READ] + buffer->indices[BUF_INDEX_INSERT]);
}

static void advance_index(circular_buffer_t * const buffer, BUF_INDEX_t index_type)
{
    if (buffer == NULL)
    {
        return;
    }

    const bool wrap_around = ((buffer->indices[index_type] + 1u) == buffer->capacity);

    buffer->indices[index_type] = wrap_around ? 0u : (buffer->indices[index_type] + 1u);
    return;
}

static uint16_t write_internal(circular_buffer_t * const buffer, const buffer_value_t value, const BUF_WRITE_MODE_t mode)
{
    if ((buffer == NULL) || (buffer->values == NULL))
    {
        errno = EINVAL;
        return EXIT_FAILURE;
    }

    if (buffer->is_full)
    {
        if (mode == BUF_WRITE_MODE_ERROR_IF_FULL)
        {
            errno = ENOBUFS;
            return EXIT_FAILURE;
        }

        advance_index(buffer, BUF_INDEX_READ);
    }

    buffer->values[buffer->indices[BUF_INDEX_INSERT]] = value;
    advance_index(buffer, BUF_INDEX_INSERT);

    if (buffer->indices[BUF_INDEX_INSERT] == buffer->indices[BUF_INDEX_READ])
    {
        buffer->is_full = true;
    }
    return EXIT_SUCCESS;
}

circular_buffer_t * new_circular_buffer(const size_t capacity)
{
    if (capacity < 1U)
    {
        return NULL;
    }

    circular_buffer_t * const buffer = (circular_buffer_t *) malloc(sizeof(circular_buffer_t));
    buffer_value_t * const values = (buffer_value_t *) calloc(capacity, sizeof(buffer_value_t));
    if ((buffer == NULL) || (values == NULL))
    {
        return NULL;
    }
    *buffer = (circular_buffer_t) {
        .capacity = capacity,
        .values = values,
    };
    return buffer;
}

int16_t delete_buffer(circular_buffer_t * const buffer)
{
    if (buffer == NULL)
    {
        errno = EINVAL;
        return EXIT_FAILURE;
    }
    free(buffer->values);
    free(buffer);
    return EXIT_SUCCESS;
}

int16_t write(circular_buffer_t * const buffer, const buffer_value_t value)
{
    return write_internal(buffer, value, BUF_WRITE_MODE_ERROR_IF_FULL);
}

int16_t overwrite(circular_buffer_t * const buffer, const buffer_value_t value)
{
    return write_internal(buffer, value, BUF_WRITE_MODE_OVERWRITE_IF_FULL);
}

int16_t read(circular_buffer_t * const buffer, buffer_value_t * const out_value)
{
    if ((buffer == NULL) || (out_value == NULL))
    {
        errno = EINVAL;
        return EXIT_FAILURE;
    }
    const size_t cnt = count_elements_internal(buffer);

    if (cnt == 0u)
    {
        errno = ENODATA;
        return EXIT_FAILURE;
    }

    *out_value = buffer->values[buffer->indices[BUF_INDEX_READ]];
    advance_index(buffer, BUF_INDEX_READ);
    buffer->is_full = false;
    return EXIT_SUCCESS;
}

int16_t clear_buffer(circular_buffer_t *const buffer) {
    if ((buffer == NULL) || (buffer->values == NULL)) {
        errno = EINVAL;
        return EXIT_FAILURE;
    }

    buffer->indices[BUF_INDEX_INSERT] = 0u;
    buffer->indices[BUF_INDEX_READ] = 0u;

    buffer->is_full = false;

    return EXIT_SUCCESS;
}
