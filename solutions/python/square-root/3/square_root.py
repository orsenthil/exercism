def square_root(number):
    x  = number / 2.0
    for i in range(10):
        x = (x + number / x) / 2.0
    return int(x)