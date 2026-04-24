#ifndef CIRCULAR_BUFFER_H
#define CIRCULAR_BUFFER_H
#include <stdint.h>
#include <stdlib.h>

typedef int16_t buffer_value_t;                                                                                                                                                                                                           
                                                                                    
typedef struct {                                                                                                              
    buffer_value_t *data;   // heap-allocated array       
    size_t capacity;                                                                                                                                                                                                                      
    size_t count;           // how many items currently stored
    size_t read_pos;                                                                                                                                                                                                                      
    size_t write_pos;                                                                                                                                                                                                                     
} circular_buffer_t;  


circular_buffer_t *new_circular_buffer(size_t capacity);
void delete_buffer(circular_buffer_t *buffer);
int16_t write(circular_buffer_t *buffer, buffer_value_t value);
int16_t overwrite(circular_buffer_t *buffer, buffer_value_t value);
int16_t read(circular_buffer_t *buffer, buffer_value_t *value);
void clear_buffer(circular_buffer_t *buffer);

#endif
