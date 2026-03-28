import math

def score(x, y):
    """
    The outer circle has a radius of 10 units (this is equivalent to the total radius for the entire target), the middle circle a radius of 5 units, and the inner circle a radius of 1. Of course, they are all centered at the same point — that is, the circles are concentric defined by the coordinates (0, 0).

    :param x:
    :param y:
    :return:
    """
    # distance = sqrt((x2 - x1)^2 + (y2 - y1)^2)

    d = math.sqrt((x - 0) ** 2 + (y - 0) ** 2)
    if d > 10:
        return 0
    elif d > 5:
        return 1
    elif d > 1:
        return 5
    else:
        return 10
