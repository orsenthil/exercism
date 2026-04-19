#include "dnd_character.h"
#include <math.h>
#include <stdlib.h>
#include <time.h>

int modifier(int score) { return floor((score - 10) / 2.0); }

int ability(void) {
  int rolls[4];
  int minimum = 7;
  int total = 0;
  for (int i = 0; i < 4; i++) {
    rolls[i] = (rand() % 6 + 1);
    if (rolls[i] < minimum) {
      minimum = rolls[i];
    }
    total += rolls[i];
  }
  return total - minimum;
}

dnd_character_t make_dnd_character(void) {
  srand(time(NULL));
  dnd_character_t character;

  character.charisma = ability();
  character.constitution = ability();
  character.dexterity = ability();
  character.intelligence = ability();
  character.strength = ability();
  character.wisdom = ability();

  character.hitpoints = 10 + modifier(character.constitution);

  return character;
}
