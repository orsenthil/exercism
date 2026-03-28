def rotate(text, key):
    lower = "abcdefghijklmnopqrstuvwxyz"
    upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    cipher = ""
    for c in text:
        if c in lower:
            cipher += lower[(lower.index(c) + key) % 26]
        elif c in upper:
            cipher += upper[(upper.index(c) + key) % 26]
        else:
            cipher += c
    return cipher
