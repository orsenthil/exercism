#include "variable_length_quantity.h"

int encode(const uint32_t *integers, size_t integers_len, uint8_t *output)
{
   // write to `output`, return final output's length
   // `output` buffer should be enough to hold the full result

   int total = 0;

   for (size_t i = 0; i < integers_len; i++) {
      uint32_t current_integer = integers[i];
      int temp_len = 0;
      uint8_t temp[5];
      do {
         temp[temp_len++] = current_integer & 0x7F;
         current_integer >>= 7;
      } while (current_integer > 0);

      total += temp_len;

      for (int j = temp_len - 1; j >= 0; j--) {
         *output = (j > 0) ? (temp[j] | 0x80) : temp[j];
         output++;
      }
   }

   return total;
}

int decode(const uint8_t *bytes, size_t buffer_len, uint32_t *output)
{
   // write to `output`, return final output's length
   // return -1 if error
   // `output` buffer should be enough to hold the full result

   int total = 0;
   uint32_t current = 0;


   for (size_t i = 0; i < buffer_len; i++) {
      uint8_t byte = bytes[i];
      current = (current << 7) | (byte & 0x7F);
      if (!(byte & 0x80)) {
         output[total++] = current;
         current = 0;
      }
   }

   if (bytes[buffer_len - 1] & 0x80) {
      return -1;
   }

   return total;
}
