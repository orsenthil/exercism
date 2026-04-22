#include "beer_song.h"
#include <stdio.h>


void recite(uint8_t start_bottles, uint8_t take_down, char **song) {
    int slot = 0;
    for (uint8_t i = 0; i < take_down; i++) {
        uint8_t bottles = start_bottles - i;
        if (bottles > 1) {
            sprintf(song[slot], "%d bottles of beer on the wall, %d bottles of beer.", bottles, bottles);
            sprintf(song[++slot], "Take one down and pass it around, %d %s of beer on the wall.", bottles - 1, (bottles - 1 == 1) ? "bottle" : "bottles");
            sprintf(song[++slot], "");
        } else if (bottles == 1) {
            sprintf(song[slot], "1 bottle of beer on the wall, 1 bottle of beer.");
            sprintf(song[++slot], "Take it down and pass it around, no more bottles of beer on the wall.");
            sprintf(song[++slot], "");
        } else {
            sprintf(song[slot], "No more bottles of beer on the wall, no more bottles of beer.");
            sprintf(song[++slot], "Go to the store and buy some more, 99 bottles of beer on the wall.");
        }
        slot++;
    }
}
