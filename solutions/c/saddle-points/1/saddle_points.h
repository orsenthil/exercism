#ifndef SADDLE_POINTS_H
#define SADDLE_POINTS_H

#include <stdint.h>
#include <stddef.h>

typedef struct {                                                                                                                                                                                                                          
    size_t row;                                                             
    size_t column;                                                                                                                                                                                                                        
} saddle_point_t;                                                           
                                                                                        
typedef struct {                                                                                                                                                                                                                          
    size_t count;
    saddle_point_t *points;  // heap-allocated array                                                                                                                                                                                      
} saddle_points_t;    

saddle_points_t *saddle_points(size_t rows, size_t cols, uint8_t matrix[rows][cols]); 
void free_saddle_points(saddle_points_t *saddle_points);

#endif
