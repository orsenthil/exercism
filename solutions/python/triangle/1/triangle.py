def equilateral(sides):
    if 0 in sides:
        return False

    a, b, c = sides

    return a == b == c


def isosceles(sides):

    if 0 in sides:
        return False

    a, b, c = sides
    return a == b or b == c or a == c


def scalene(sides):
    if 0 in sides:
        return False
    a, b, c = sides
    return a != b and b != c and a != c
