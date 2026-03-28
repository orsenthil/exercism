def sum_of_powers(number, power):
    sum = 0
    for digit in str(number):
        sum += pow(int(digit), power)
    return sum


def is_armstrong_number(number):
    number_of_digits = len(str(number))
    return number == sum_of_powers(number, number_of_digits)
