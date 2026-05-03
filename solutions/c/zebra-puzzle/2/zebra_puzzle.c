#include "zebra_puzzle.h"
#include <stdlib.h>

typedef enum { NORWEGIAN, ENGLISHMAN, SPANIARD, UKRAINIAN, JAPANESE } nationality_t;
typedef enum { RED, GREEN, IVORY, YELLOW, BLUE } color_t;
typedef enum { DOG, SNAIL, HORSE, ZEBRA, FOX } pet_t;
typedef enum { TEA, COFFEE, MILK, ORANGE_JUICE, WATER } drink_t;
typedef enum { DANCING, PAINTING, READING, FOOTBALL, CHESS } hobby_t;


static int next_to(int a, int b) {
    return abs(a - b) == 1;
}

static int find(int *arr, int n, int value) {
    for (int i = 0; i < n; i++) {
        if (arr[i] == value) {
            return i;
        }
    }
    return -1;
}

static int next_permutation(int *arr, int n) {
    // step 1: find largest i where arr[i] < arr[i+1]
    int i = -1;

    for (int a = n - 2; a >= 0; a--) {
        if (arr[a] < arr[a + 1]) {
            i = a;
            break;
        }
    }

    // if no such i, we're at the last permutation — reset and return 0
    if (i < 0) {
        for (int a = 0; a < n/2; a++) {
            int tmp = arr[a];
            arr[a] = arr[n-1-a];
            arr[n-1-a] = tmp;
        }
        return 0;
    }

    // step 2: find largest j > i where arr[j] > arr[i]
    int j = -1;
    for (int a = n - 1; a > i; a--) {
        if (arr[a] > arr[i]) {
            j = a;
            break;
        }
    }

    // step 3: swap arr[i] and arr[j]
    int tmp = arr[i];
    arr[i] = arr[j];
    arr[j] = tmp;

    // step 4: reverse everything from i+1 to end
    for (int a = 0; a < (n - i - 1)/2; a++) {
        int tmp = arr[i + 1 + a];
        arr[i + 1 + a] = arr[n - 1 - a];
        arr[n - 1 - a] = tmp;
    }

    return 1;
}

solution_t solve_puzzle(void) {
    nationality_t nationals[5] = { NORWEGIAN, ENGLISHMAN, SPANIARD, UKRAINIAN, JAPANESE };
    color_t colors[5] = { RED, GREEN, IVORY, YELLOW, BLUE };
    pet_t pets[5] = { DOG, SNAIL, HORSE, ZEBRA, FOX };
    drink_t drinks[5] = { TEA, COFFEE, MILK, ORANGE_JUICE, WATER };
    hobby_t hobbies[5] = { DANCING, PAINTING, READING, FOOTBALL, CHESS };

    const char *names[] = {"Norwegian", "Englishman", "Spaniard", "Ukrainian", "Japanese"};

    do {
        // check nationality constraints here (clue 10)
        if (nationals[0] != NORWEGIAN) continue;
    
        do {
            // check nat+color constraints here (clues 2, 6, 15)
            // The correct check is: find the house i where nationals[i] == ENGLISHMAN, then verify colors[i] == RED.
            int eng = find((int *)nationals, 5, ENGLISHMAN);
            if (eng == -1 || colors[eng] != RED) continue;

            // Clue 6 - 6. The green house is immediately to the right of the ivory house.
            int ivory = find((int *)colors, 5, IVORY);
            int green = find((int *)colors, 5, GREEN);
            if (ivory + 1 != green) continue;

            // 15. The Norwegian lives next to the blue house.
            int blue = find((int *)colors, 5, BLUE);
            if (!next_to(0, blue)) continue;

            // Drink loop
            do {
                // 5. The Ukrainian drinks tea.
                int tea = find((int *)drinks, 5, TEA);
                if (tea == -1 || nationals[tea] != UKRAINIAN) continue;

                // 4. The person in the green house drinks coffee.
                int coffee = find((int *)drinks, 5, COFFEE);
                if (coffee == -1 || colors[coffee] != GREEN) continue;

                // 9. The person in the middle house drinks milk.
                if (drinks[2] != MILK) continue;

                // Pet loop
                do {
                    // 3. The Spaniard owns the dog.
                    int dog = find((int *)pets, 5, DOG);
                    if (dog == -1 || nationals[dog] != SPANIARD) continue;
                    
                    // Hobby Lobby loop
                    do {
                        // Clue 7: snail owner → dancing
                        // find SNAIL in pets, check hobbies[snail_house] == DANCING
                        int snail = find((int *)pets, 5, SNAIL);
                        if (snail == -1 || hobbies[snail] != DANCING) continue;

                        // Clue 8: yellow house → painting
                        // Clue 8: find YELLOW in colors, check hobbies[yellow_house] == PAINTING
                        int yellow = find((int *)colors, 5, YELLOW);
                        if (yellow == -1 || hobbies[yellow] != PAINTING) continue;

                        // Clue 11: reading next to fox
                        // find READING in hobbies, check next_to(reading_house, fox_house)
                        int reading = find((int *)hobbies, 5, READING);
                        int fox = find((int *)pets, 5, FOX);
                        if (reading == -1 || !next_to(reading, fox)) continue;

                        // Clue 12: painter next to horse
                        // find PAINTING in hobbies, check next_to(painting_house, horse_house)
                        int painting = find((int *)hobbies, 5, PAINTING);
                        int horse = find((int *)pets, 5, HORSE);
                        if (painting == -1 || !next_to(painting, horse)) continue;

                        // Clue 13: football → orange juice
                        // find FOOTBALL in hobbies, check drinks[football_house] == ORANGE_JUICE
                        int football = find((int *)hobbies, 5, FOOTBALL);
                        if (football == -1 || drinks[football] != ORANGE_JUICE) continue;

                        // Clue 14: Japanese → chess
                        // find JAPANESE in nationals, check hobbies[japanese_house] == CHESS
                        int japanese = find((int *)nationals, 5, JAPANESE);
                        if (japanese == -1 || hobbies[japanese] != CHESS) continue;

                        // return solution
                        solution_t solution;
                        // find which house has WATER.
                        int house_index = find((int *)drinks, 5, WATER);
                        solution.drinks_water = names[nationals[house_index]];

                        // find which house has ZEBRA.
                        house_index = find((int *)pets, 5, ZEBRA);
                        solution.owns_zebra = names[nationals[house_index]];

                        return solution;

                    } while (next_permutation((int *)hobbies, 5));

                } while (next_permutation((int *)pets, 5));

            } while (next_permutation((int *)drinks, 5));
       
    
        } while (next_permutation((int *)colors, 5));
    } while (next_permutation((int *)nationals, 5));

    return (solution_t){ .drinks_water = NULL, .owns_zebra = NULL };

}

