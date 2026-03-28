"""
This exercise stub and the test suite contain several enumerated constants.

Enumerated constants can be done with a NAME assigned to an arbitrary,
but unique value. An integer is traditionally used because it’s memory
efficient.
It is a common practice to export both constants and functions that work with
those constants (ex. the constants in the os, subprocess and re modules).

You can learn more here: https://en.wikipedia.org/wiki/Enumerated_type
"""

from enum import Enum

# Possible sublist categories.
# Change the values as you see fit.
listEnum = Enum('listEnum', 'SUBLIST SUPERLIST EQUAL UNEQUAL')

SUBLIST = listEnum.SUBLIST
SUPERLIST = listEnum.SUPERLIST
EQUAL = listEnum.EQUAL
UNEQUAL = listEnum.UNEQUAL


def sublist(list_one, list_two):
    if list_one == list_two:
        return EQUAL

    str_list_one = []
    for element in list_one:
        str_list_one.append(f'#{element}#')

    str_list_two = []
    for element in list_two:
        str_list_two.append(f'#{element}#')

    str_list_one = '-'.join(str_list_one)
    str_list_two = '-'.join(str_list_two)

    if str_list_one in str_list_two:
        return SUBLIST

    if str_list_two in str_list_one:
        return SUPERLIST

    return UNEQUAL