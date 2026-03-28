def is_valid(isbn):
    isbn = isbn.replace("-", "")
    if len(isbn) != 10:
        return False
    start = 10
    total = 0
    for i, c in enumerate(isbn):
        if not c.isdigit() and c != "X":
            return False
        if c == "X" and i != 9:
            return False
        total += start * (10 if c == "X" else int(c))
        start -= 1

    return total % 11 == 0