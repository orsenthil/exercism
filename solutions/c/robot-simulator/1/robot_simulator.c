#include "robot_simulator.h"
#include <string.h>

robot_status_t robot_create(robot_direction_t direction, int x, int y) {
  return (robot_status_t){direction, (robot_position_t){x, y}};
}

void robot_move(robot_status_t *robot, const char *commands) {
  robot_direction_t current_direction;
  int length = strlen(commands);

  for (int i = 0; i < length; i++) {
    current_direction = robot->direction;
    switch (commands[i]) {
    case 'L':
      robot->direction = (current_direction + 3) % 4;
      break;
    case 'R':
      robot->direction = (current_direction + 1) % 4;
      break;
    case 'A':
      switch (current_direction) {
      case DIRECTION_NORTH:
        robot->position.y += 1;
        break;
      case DIRECTION_EAST:
        robot->position.x += 1;
        break;
      case DIRECTION_SOUTH:
        robot->position.y -= 1;
        break;
      case DIRECTION_WEST:
        robot->position.x -= 1;
        break;
      default:
        break;
      }
    }
  }
}
