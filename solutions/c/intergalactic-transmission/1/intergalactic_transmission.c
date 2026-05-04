#include "intergalactic_transmission.h"

#include <stddef.h>


int transmit_sequence(uint8_t *buffer, const uint8_t *message, int message_length) {
    uint32_t accumulator = 0;
    int bits_in_accumulator = 0;

    // Write the loop that pushes each input byte's bits into accumulator and tracks bits_in_accumulator.
    //  When bits_in_accumulator >= 7, pull out the top 7 bits and emit a transmission byte (just hardcode parity = 0 for now — we'll fix that next).
    if (message == NULL || message_length == 0) {
        return 0;
    }

    int index = 0;

    for (int i = 0; i < message_length; i++) {
        accumulator <<= 8;
        accumulator |= message[i];
        bits_in_accumulator += 8;
        while (bits_in_accumulator >= 7) {
            uint8_t transmission_byte = (accumulator >> (bits_in_accumulator - 7)) & 0x7F;
            uint8_t parity = __builtin_popcount(transmission_byte) % 2;
            transmission_byte = (transmission_byte << 1) | parity;
            buffer[index++] = transmission_byte;
            bits_in_accumulator -= 7;
        }
    }

    if (bits_in_accumulator > 0) {
        accumulator = accumulator & (1u << bits_in_accumulator) - 1;
        accumulator <<= 7 - bits_in_accumulator;
        uint8_t leftover = accumulator & 0x7F;
        uint8_t parity = __builtin_popcount(leftover) % 2;
        buffer[index++] = (leftover << 1) | parity;
    }

    return index;
}

int decode_message(uint8_t *buffer, const uint8_t *message, int message_length) {
    uint32_t accumulator = 0;
    int bits_in_accumulator = 0;
    int index = 0;

    if (message == NULL || message_length == 0) {
        return 0;
    }

    // for each input byte:
    // check parity -> return WRONG_PARITY if parity is wrong

    for (int i = 0; i < message_length; i++) {
        uint8_t byte = message[i];
        uint8_t parity = __builtin_popcount(byte) % 2;
        if (parity != 0) {
            return WRONG_PARITY;
        }

        // Step 2 — now add the accumulator logic inside the same loop, after the parity check. For each byte:

        byte = (byte >> 1) & 0x7F;

        accumulator <<= 7;
        accumulator |= byte;
        bits_in_accumulator += 7;
        while (bits_in_accumulator >= 8) {
            uint8_t data_byte = (accumulator >> (bits_in_accumulator - 8)) & 0xFF;
            buffer[index++] = data_byte;
            bits_in_accumulator -= 8;
        }
    }

    return index;
}


