#include "two_bucket.h"
#include <string.h>


typedef struct {
    bucket_liters_t current_size_bucket_1;
    bucket_liters_t current_size_bucket_2;
    int move_count;
} state_t;

bucket_result_t measure(bucket_liters_t bucket_1_size,
    bucket_liters_t bucket_2_size,
    bucket_liters_t goal_volume, bucket_id_t start_bucket) {

    state_t queue[256];
    bool visited[256][256];
    memset(visited, 0, sizeof(visited));

    int head = 0, tail = 0;
    // initial state
    // first action is filling the starting bucket. That counts as move 1. Your initial enqueue should reflect that first fill already applied.
    state_t initial_state = {
        .current_size_bucket_1 = start_bucket == BUCKET_ID_1 ? bucket_1_size : 0,
        .current_size_bucket_2 = start_bucket == BUCKET_ID_2 ? bucket_2_size : 0,
        .move_count = 1
    };

    // breadth first search
    queue[tail++] = initial_state;
    visited[initial_state.current_size_bucket_1][initial_state.current_size_bucket_2] = true;
    while (head < tail) {
        state_t current_state = queue[head++];
        if (current_state.current_size_bucket_1 == goal_volume || current_state.current_size_bucket_2 == goal_volume) {
            return (bucket_result_t){
                .possible = true,
                .move_count = current_state.move_count,
                .goal_bucket = current_state.current_size_bucket_1 == goal_volume ? BUCKET_ID_1 : BUCKET_ID_2,
                .other_bucket_liters = current_state.current_size_bucket_1 == goal_volume ? current_state.current_size_bucket_2 : current_state.current_size_bucket_1
            };
        }

        // Pour Logic

        bucket_liters_t new_b2 = (current_state.current_size_bucket_1 + current_state.current_size_bucket_2 < bucket_2_size) ? current_state.current_size_bucket_1 + current_state.current_size_bucket_2 : bucket_2_size;
        bucket_liters_t new_b1 = current_state.current_size_bucket_1 - (new_b2 - current_state.current_size_bucket_2);

        // forbidden: start bucket empty, other bucket full
        bool forbidden = (start_bucket == BUCKET_ID_1 && new_b1 == 0 && new_b2 == bucket_2_size) || (start_bucket == BUCKET_ID_2 && new_b2 == 0 && new_b1 == bucket_1_size);

        if (!forbidden && !visited[new_b1][new_b2]) {

            visited[new_b1][new_b2] = true;

            // pour from bucket 1 to bucket 2
            queue[tail++] = (state_t){
                .current_size_bucket_1 = new_b1,
                .current_size_bucket_2 = new_b2,
                .move_count = current_state.move_count + 1
            };

        }

        bucket_liters_t b1_after = (current_state.current_size_bucket_1 + current_state.current_size_bucket_2 < bucket_1_size) ?  current_state.current_size_bucket_1 + current_state.current_size_bucket_2: bucket_1_size;
        bucket_liters_t b2_after = current_state.current_size_bucket_2 - (b1_after - current_state.current_size_bucket_1);

        // forbidden: start bucket empty, other bucket full
        forbidden = (start_bucket == BUCKET_ID_1 && b1_after == 0 && b2_after == bucket_2_size) || (start_bucket == BUCKET_ID_2 && b2_after == 0 && b1_after == bucket_1_size);


        if (!forbidden && !visited[b1_after][b2_after]) {
            visited[b1_after][b2_after] = true;

            // pour from bucket 2 to bucket 1
            queue[tail++] = (state_t){
                .current_size_bucket_1 = b1_after,
                .current_size_bucket_2 = b2_after,
                .move_count = current_state.move_count + 1
            };

        }

        forbidden = start_bucket == BUCKET_ID_1 && current_state.current_size_bucket_2 == bucket_2_size;

        if (!forbidden && !visited[0][current_state.current_size_bucket_2]) {
            visited[0][current_state.current_size_bucket_2] = true;

            // empty bucket 1
            queue[tail++] = (state_t){
                .current_size_bucket_1 = 0,
                .current_size_bucket_2 = current_state.current_size_bucket_2,
                .move_count = current_state.move_count + 1
            };

        }

        forbidden = start_bucket == BUCKET_ID_2 && current_state.current_size_bucket_1 == bucket_1_size;

        if (!forbidden && !visited[current_state.current_size_bucket_1][0]) {
            visited[current_state.current_size_bucket_1][0] = true;
            // empty bucket 2
            queue[tail++] = (state_t){
                .current_size_bucket_1 = current_state.current_size_bucket_1,
                .current_size_bucket_2 = 0,
                .move_count = current_state.move_count + 1
            };

        }

        forbidden = start_bucket == BUCKET_ID_2 && current_state.current_size_bucket_2 == 0;

        if (!forbidden && !visited[bucket_1_size][current_state.current_size_bucket_2]) {
            visited[bucket_1_size][current_state.current_size_bucket_2] = true;

            // fill bucket 1
            queue[tail++] = (state_t){
                .current_size_bucket_1 = bucket_1_size,
                .current_size_bucket_2 = current_state.current_size_bucket_2,
                .move_count = current_state.move_count + 1
            };

        }

        forbidden = start_bucket == BUCKET_ID_1 && current_state.current_size_bucket_1 == 0;

        if (!forbidden && !visited[current_state.current_size_bucket_1][bucket_2_size]) {
            visited[current_state.current_size_bucket_1][bucket_2_size] = true;
            // fill bucket 2
            queue[tail++] = (state_t){
                .current_size_bucket_1 = current_state.current_size_bucket_1,
                .current_size_bucket_2 = bucket_2_size,
                .move_count = current_state.move_count + 1
            };

        }

     }

    return (bucket_result_t){
        .possible = false,
        .move_count = 0,
        .goal_bucket = BUCKET_ID_1,
        .other_bucket_liters = 0
    };
 
}
