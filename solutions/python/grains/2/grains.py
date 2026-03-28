def square(number):
    """
    The number of grains on each square doubles.

    The first slot has 1 grain, and each subsequent slot has twice as many grains as the previous slot.
    The slot should be a valid square number.

    :param number: slot number
    :return: number of grains on the slot
    """
    if number < 1 or number > 64:
        raise ValueError("square must be between 1 and 64")

    return 2 ** (number - 1)


def total():
    """
    The total number of grains on a chessboard with 64 squares is calculated by doubling the number of grains on each subsequent square, starting with 1 grain on the first square (square 1).

    Now let's see how this calculation unfolds:

    - On square 1: We have 1 grain.
    - On square 2: We have 2 grains (doubled from the previous square).
    - On square 3: We have 4 grains (doubled from the previous square).
    - ...
    - On square n (any square from 1 to 64): We have 2^(n-1) grains. Each square has twice the number of grains as the previous square.

    Hence, the total number of grains on the chessboard can be calculated by summing up the grains on each square:

    Total = 2^0 + 2^1 + 2^2 + ... + 2^63

    This mathematical series can be expressed in a simplified form using a geometric progression formula:

    Total = (2^64) - 1

    The -1 comes from excluding the first square, which already accounts for 1 grain. Therefore, the total number of grains on a 64 square chessboard is 2^64 -1.

    :return: The total number of grains on a 64 square chessboard.
    """
    return (2 ** 64) - 1