#ifndef RAINDROPS_H
#define RAINDROPS_H

#define DROP_NAME_LENGTH 5

typedef struct RainDrop {
    int divisor;
    char dropName[DROP_NAME_LENGTH];
}RainDrop;

char *convert(char result[], int drops);

#endif
