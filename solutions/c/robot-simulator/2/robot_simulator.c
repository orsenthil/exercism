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
    while (*commands != '\0')
    {
        if (*commands == 'A')
        {
            if (robot->direction == DIRECTION_NORTH)
            {
                robot->position.y += 1;
            }
            else if (robot->direction == DIRECTION_EAST)
            {
                robot->position.x += 1;
            }
            else if (robot->direction == DIRECTION_SOUTH)
            {
                robot->position.y -= 1;
            }
            else if (robot->direction == DIRECTION_WEST)
            {
                robot->position.x -= 1;
            }
        }
        else if  (*commands == 'R')
        {
            robot->direction = (robot->direction + 1) % DIRECTION_MAX;
        }

        else if (*commands == 'L')
        {
            robot->direction = (robot->direction - 1) % DIRECTION_MAX;
        }

        ++commands;
    }
}
