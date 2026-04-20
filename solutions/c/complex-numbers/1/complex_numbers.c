#include "complex_numbers.h"

complex_t c_add(complex_t a, complex_t b)
{
   return (complex_t){ .real = a.real + b.real, .imag = a.imag + b.imag };
}

complex_t c_sub(complex_t a, complex_t b)
{
   return (complex_t){ .real = a.real - b.real, .imag = a.imag - b.imag };
}

complex_t c_mul(complex_t a, complex_t b)
{
   return (complex_t){ .real = a.real * b.real - a.imag * b.imag, .imag = a.real * b.imag + a.imag * b.real };
}

complex_t c_div(complex_t a, complex_t b)
{
   double denominator = b.real * b.real + b.imag * b.imag;
   return (complex_t){ .real = (a.real * b.real + a.imag * b.imag) / denominator, .imag = (a.imag * b.real - a.real * b.imag) / denominator };
}

double c_abs(complex_t x)
{
   return sqrt(x.real * x.real + x.imag * x.imag);
}

complex_t c_conjugate(complex_t x)
{
   return (complex_t){ .real = x.real, .imag = -x.imag };
}

double c_real(complex_t x)
{
   return x.real;
}

double c_imag(complex_t x)
{
   return x.imag;
}

complex_t c_exp(complex_t x)
{
   double exp_real = exp(x.real);
   return (complex_t){ .real = exp_real * cos(x.imag), .imag = exp_real * sin(x.imag) };
}
