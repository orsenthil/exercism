#include "robot_simulator.h"
#include <string.h>

robot_status_t robot_create(robot_direction_t direction, int x, int y)
{
    robot_position_t position;
    position.x = x;
    position.y = y;
    robot_status_t status;
    status.direction = direction;
    status.position = position;
    return status;
}

void robot_move(robot_status_t *robot, const char *commands)
{
    int commandlen= strlen(commands);
    for (int i = 0; i < commandlen; i++) {
        if (commands[i] == 'R') {
            robot->position.x += 1;
        }
    }
}
